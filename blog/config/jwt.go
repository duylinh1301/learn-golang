package config

var Jwt = loadJwtConfig()

func loadJwtConfig() map[string]interface{} {
	return map[string]interface{}{
		"secret_key": "golangapp",
		"ttl":        40320, //
	}
}
