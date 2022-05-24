package conf

//go:generate optiongen --option_with_struct_name=false
func OptionsOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"Port":      8080,
		"KeepAlive": true,
		"Debug":     false,
	}
}
