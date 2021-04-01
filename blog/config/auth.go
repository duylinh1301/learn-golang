package config

var Auth = NewConfigAuth()

type configAuth struct {
	AuthType string
}

func NewConfigAuth() *configAuth {
	return &configAuth{
		AuthType: "jwt", // jwt, session
	}
}
