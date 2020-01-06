package auth

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

func SsoGetHandler(fs http.FileSystem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := false
		if r.URL.Query().Get("auth_error") != "" {
			err = true
		}

		data := struct {
			RedirectURL string
			Error       bool
		}{RedirectURL: r.URL.Query().Get("redirectUrl"), Error: err}
		renderTemplate(w, fs, "auth/login.html", &data)
	}
}

func renderTemplate(w http.ResponseWriter, fs http.FileSystem, tmpl string, p interface{}) {
	file, err := fs.Open(tmpl)
	if err != nil {
		log.Error().Err(err).Msg("file not found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error().Err(err).Msg("failed to read file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	s := string(b)

	templates, err := template.New(tmpl).Parse(s)
	if err != nil {
		log.Error().Err(err).Msg("failed to load template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		log.Error().Err(err).Msg("failed to execute template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SsoPostHandler(sso *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		pUri := r.PostFormValue("redirectUrl")
		u, g, err := sso.Auth(r.PostFormValue("username"), r.PostFormValue("password"))
		if err != nil {
			if sso.Is401(err) {
				log.Err(err).Msg("failed to login")
				http.Redirect(w, r, fmt.Sprintf("/sso?s_url=%s&auth_error=true", pUri), 301)
				return
			}
			log.Err(err).Msg("internal error login")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Not able to service this request. Please try again later.")
			return
		}

		vh := sso.CookieValidityMinutes
		exp := time.Now().Add(time.Minute * time.Duration(vh)).UTC()
		tok, _ := sso.BuildToken(u, g, exp)
		c := sso.BuildCookie(tok, exp)
		http.SetCookie(w, &c)
		http.Redirect(w, r, pUri, http.StatusFound)
		return
	}

}

func TokenHandler(sso *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("username")
		password := r.PostFormValue("password")
		u, g, err := sso.Auth(email, password)
		logEvent := log.Err(err).Str("email", email)
		if err != nil {
			if sso.Is401(err) {
				logEvent.Msg("authentication failed")
				fmt.Fprintf(w, "Unauthorized.")
				return
			}
			logEvent.Msg("authentication: internal error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Not able to service the request. Please try again later.")
			return
		}

		tok, _ := sso.BuildToken(u, g, time.Now().Add(time.Minute*time.Duration(sso.CookieValidityMinutes)).UTC())
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tok)
		return
	}
}

func LogoutHandler(sso *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expT := time.Now().Add(time.Hour * time.Duration(-1))
		lc := sso.Logout(expT)
		http.SetCookie(w, &lc)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "You have been logged out.")
		return
	}
}

type User struct {
	UID   string
	Email string
}
