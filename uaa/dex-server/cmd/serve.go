package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"mmrath.com/gobase/common/log"
	"mmrath.com/gobase/uaa/pkg/config"

	"github.com/ghodss/yaml"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/dexidp/dex/api"
	"github.com/dexidp/dex/server"
	"github.com/dexidp/dex/storage"
)

func commandServe() *cobra.Command {
	return &cobra.Command{
		Use:     "serve [ config file ]",
		Short:   "Connect to the storage and begin serving requests.",
		Long:    ``,
		Example: "dex serve config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			if err := serve(cmd, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}
}

func serve(cmd *cobra.Command, args []string) error {
	switch len(args) {
	default:
		return errors.New("surplus arguments")
	case 0:
		// TODO(ericchiang): Consider having a default config file location.
		return errors.New("no arguments provided")
	case 1:
	}

	configFile := args[0]
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %v", configFile, err)
	}

	var c config.Config
	if err := yaml.Unmarshal(configData, &c); err != nil {
		return fmt.Errorf("error parse config file %s: %v", configFile, err)
	}


	if c.Logger.Level != "" {
		log.Infof("config using log level: %s", c.Logger.Level)
	}
	if err := c.Validate(); err != nil {
		return err
	}

	log.Infof("config issuer: %s", c.Issuer)

	prometheusRegistry := prometheus.NewRegistry()
	err = prometheusRegistry.Register(prometheus.NewGoCollector())
	if err != nil {
		return fmt.Errorf("failed to register Go runtime metrics: %v", err)
	}

	err = prometheusRegistry.Register(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	if err != nil {
		return fmt.Errorf("failed to register process metrics: %v", err)
	}

	grpcMetrics := grpcprometheus.NewServerMetrics()
	err = prometheusRegistry.Register(grpcMetrics)
	if err != nil {
		return fmt.Errorf("failed to register gRPC server metrics: %v", err)
	}

	var grpcOptions []grpc.ServerOption

	if c.GRPC.TLSCert != "" {
		// Parse certificates from certificate file and key file for server.
		cert, err := tls.LoadX509KeyPair(c.GRPC.TLSCert, c.GRPC.TLSKey)
		if err != nil {
			return fmt.Errorf("invalid config: error parsing gRPC certificate file: %v", err)
		}

		tlsConfig := tls.Config{
			Certificates:             []tls.Certificate{cert},
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
		}

		if c.GRPC.TLSClientCA != "" {
			// Parse certificates from client CA file to a new CertPool.
			cPool := x509.NewCertPool()
			clientCert, err := ioutil.ReadFile(c.GRPC.TLSClientCA)
			if err != nil {
				return fmt.Errorf("invalid config: reading from client CA file: %v", err)
			}
			if !cPool.AppendCertsFromPEM(clientCert) {
				return errors.New("invalid config: failed to parse client CA")
			}

			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
			tlsConfig.ClientCAs = cPool

			// Only add metrics if client auth is enabled
			grpcOptions = append(grpcOptions,
				grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
				grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
			)
		}

		grpcOptions = append(grpcOptions, grpc.Creds(credentials.NewTLS(&tlsConfig)))
	}

	logger := log.NewLogger()

	s, err := c.Storage.Config.Open(logger)
	if err != nil {
		return fmt.Errorf("failed to initialize storage: %v", err)
	}
	log.Infof("config storage: %s", c.Storage.Type)

	if len(c.StaticClients) > 0 {
		for _, client := range c.StaticClients {
			log.Infof("config static client: %s", client.Name)
		}
		s = storage.WithStaticClients(s, c.StaticClients)
	}
	if len(c.StaticPasswords) > 0 {
		passwords := make([]storage.Password, len(c.StaticPasswords))
		for i, p := range c.StaticPasswords {
			passwords[i] = storage.Password(p)
		}
		s = storage.WithStaticPasswords(s, passwords, logger)
	}

	storageConnectors := make([]storage.Connector, len(c.StaticConnectors))
	for i, c := range c.StaticConnectors {
		if c.ID == "" || c.Name == "" || c.Type == "" {
			return fmt.Errorf("invalid config: ID, Type and Name fields are required for a connector")
		}
		if c.Config == nil {
			return fmt.Errorf("invalid config: no config field for connector %q", c.ID)
		}
		log.Infof("config connector: %s", c.ID)

		// convert to a storage connector object
		conn, err := config.ToStorageConnector(c)
		if err != nil {
			return fmt.Errorf("failed to initialize storage connectors: %v", err)
		}
		storageConnectors[i] = conn

	}

	if c.EnablePasswordDB {
		storageConnectors = append(storageConnectors, storage.Connector{
			ID:   server.LocalConnector,
			Name: "Email",
			Type: server.LocalConnector,
		})
		log.Infof("config connector: local passwords enabled")
	}

	s = storage.WithStaticConnectors(s, storageConnectors)

	if len(c.OAuth2.ResponseTypes) > 0 {
		log.Infof("config response types accepted: %s", c.OAuth2.ResponseTypes)
	}
	if c.OAuth2.SkipApprovalScreen {
		log.Infof("config skipping approval screen")
	}
	if len(c.Web.AllowedOrigins) > 0 {
		log.Infof("config allowed origins: %s", c.Web.AllowedOrigins)
	}

	// explicitly convert to UTC.
	now := func() time.Time { return time.Now().UTC() }

	serverConfig := server.Config{
		SupportedResponseTypes: c.OAuth2.ResponseTypes,
		SkipApprovalScreen:     c.OAuth2.SkipApprovalScreen,
		AllowedOrigins:         c.Web.AllowedOrigins,
		Issuer:                 c.Issuer,
		Storage:                s,
		Web:                    c.Frontend,
		Logger:                 logger,
		Now:                    now,
		PrometheusRegistry:     prometheusRegistry,
	}
	if c.Expiry.SigningKeys != "" {
		signingKeys, err := time.ParseDuration(c.Expiry.SigningKeys)
		if err != nil {
			return fmt.Errorf("invalid config value %q for signing keys expiry: %v", c.Expiry.SigningKeys, err)
		}
		logger.Infof("config signing keys expire after: %v", signingKeys)
		serverConfig.RotateKeysAfter = signingKeys
	}
	if c.Expiry.IDTokens != "" {
		idTokens, err := time.ParseDuration(c.Expiry.IDTokens)
		if err != nil {
			return fmt.Errorf("invalid config value %q for id token expiry: %v", c.Expiry.IDTokens, err)
		}
		logger.Infof("config id tokens valid for: %v", idTokens)
		serverConfig.IDTokensValidFor = idTokens
	}
	if c.Expiry.AuthRequests != "" {
		authRequests, err := time.ParseDuration(c.Expiry.AuthRequests)
		if err != nil {
			return fmt.Errorf("invalid config value %q for auth request expiry: %v", c.Expiry.AuthRequests, err)
		}
		logger.Infof("config auth requests valid for: %v", authRequests)
		serverConfig.AuthRequestsValidFor = authRequests
	}

	serv, err := server.NewServer(context.Background(), serverConfig)
	if err != nil {
		return fmt.Errorf("failed to initialize server: %v", err)
	}

	telemetryServ := http.NewServeMux()
	telemetryServ.Handle("/metrics", promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{}))

	errc := make(chan error, 3)
	if c.Telemetry.HTTP != "" {
		logger.Infof("listening (http/telemetry) on %s", c.Telemetry.HTTP)
		go func() {
			err := http.ListenAndServe(c.Telemetry.HTTP, telemetryServ)
			errc <- fmt.Errorf("listening on %s failed: %v", c.Telemetry.HTTP, err)
		}()
	}
	if c.Web.HTTP != "" {
		logger.Infof("listening (http) on %s", c.Web.HTTP)
		go func() {
			err := http.ListenAndServe(c.Web.HTTP, serv)
			errc <- fmt.Errorf("listening on %s failed: %v", c.Web.HTTP, err)
		}()
	}
	if c.Web.HTTPS != "" {
		httpsSrv := &http.Server{
			Addr:    c.Web.HTTPS,
			Handler: serv,
			TLSConfig: &tls.Config{
				PreferServerCipherSuites: true,
				MinVersion:               tls.VersionTLS12,
			},
		}

		logger.Infof("listening (https) on %s", c.Web.HTTPS)
		go func() {
			err = httpsSrv.ListenAndServeTLS(c.Web.TLSCert, c.Web.TLSKey)
			errc <- fmt.Errorf("listening on %s failed: %v", c.Web.HTTPS, err)
		}()
	}
	if c.GRPC.Addr != "" {
		logger.Infof("listening (grpc) on %s", c.GRPC.Addr)
		go func() {
			errc <- func() error {
				list, err := net.Listen("tcp", c.GRPC.Addr)
				if err != nil {
					return fmt.Errorf("listening on %s failed: %v", c.GRPC.Addr, err)
				}
				s := grpc.NewServer(grpcOptions...)
				api.RegisterDexServer(s, server.NewAPI(serverConfig.Storage, logger))
				grpcMetrics.InitializeMetrics(s)
				err = s.Serve(list)
				return fmt.Errorf("listening on %s failed: %v", c.GRPC.Addr, err)
			}()
		}()
	}

	return <-errc
}

var (
	logLevels  = []string{"debug", "info", "error"}
	logFormats = []string{"json", "text"}
)
