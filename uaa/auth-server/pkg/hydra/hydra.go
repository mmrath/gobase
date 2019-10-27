package hydra

/** general hydra stuff **/
const ScopeOffline string = "offline"
const ScopeOpenId string = "openid"
var AutoScopes = [...]string{ScopeOffline, ScopeOpenId}

type Service interface {
}

type service struct {
	hydraAdminUrl string
}


