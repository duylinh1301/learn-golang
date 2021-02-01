package config

var Auth = loadAuthConfig()

func loadAuthConfig() map[string]interface{} {
	return map[string]interface{}{
		"auth_type": map[string]interface{}{
			"jwt":     1,
			"session": 2,
		},
	}
}
