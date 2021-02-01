package helpers

// InterfaceToMap convert data of interface{} to map[string]interface{}
func InterfaceToMap(data interface{}) map[string]interface{} {
	return data.(map[string]interface{})
}

// GetValueOrDefault return default value if value nil
func GetValueOrDefault(value interface{}, defaultValue interface{}) interface{} {
	if value == nil || value == "" {
		return defaultValue
	}
	return value
}
