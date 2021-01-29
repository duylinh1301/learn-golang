package config

var Jwt = loadJwtConfig()

func loadJwtConfig() map[string]interface{} {
	return map[string]interface{}{
		"ttl": 40320, //
	}
}
