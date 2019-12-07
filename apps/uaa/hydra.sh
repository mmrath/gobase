export DSN=postgres://hydra:password@host.docker.internal:5432/uaa?sslmode=disable
docker run -it --rm \
  --network uaa \
  oryd/hydra:latest \
  migrate sql --yes $DSN

export SECRETS_SYSTEM=2E91sAs7fukU2Mg5wNHM2qOR9y.jGzcN

docker run -d \
  --name ory-hydra-example--hydra \
  --network uaa \
  -p 9000:4444 \
  -p 9001:4445 \
  -e SECRETS_SYSTEM=$SECRETS_SYSTEM \
  -e DSN=$DSN \
  -e URLS_SELF_ISSUER=http://127.0.0.1:9000/ \
  -e URLS_CONSENT=http://127.0.0.1:9020/oauth2/consent \
  -e URLS_LOGIN=http://127.0.0.1:9020/oauth2/login \
  oryd/hydra:latest serve all --dangerous-force-http

# create oidc client
docker run --rm -it \
  --network uaa \
  oryd/hydra:latest \
  clients create \
    --endpoint http://ory-hydra-example--hydra:4445 \
    --id example-consumer \
    --secret example-secret \
    -g authorization_code,refresh_token \
    -r token,code,id_token \
    --scope openid,offline \
    --callbacks http://127.0.0.1:9010/example/auth/callback

