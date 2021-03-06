package param

var parameters = make(map[string]map[string]interface{})

func SetNetworkParam(network string, name string, value interface{}) {
	if parameters[network] == nil {
		parameters[network] = make(map[string]interface{})
	}

	parameters[network][name] = value
}

func GetNetworkParam(network string, name string, defaultValue interface{}) interface{} {
	if parameters[network] == nil || parameters[network][name] == nil {
		return defaultValue
	}

	return parameters[network][name]
}
