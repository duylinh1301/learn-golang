package config

var Jwt = NewConfigJwt()

type configJwt struct {
	SecretKey string
	Ttl       int64
}

func NewConfigJwt() *configJwt {
	return &configJwt{
		SecretKey: "golangapp",
		Ttl:       1440,
	}
}
