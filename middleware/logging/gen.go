package logging

//go:generate optiongen --option_with_struct_name=false
func OptionsOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"level":            "debug",
		"encodeing":        "json",
		"outputPaths":      []string{""},
		"errorOutputPaths": []string{"stderr"},
		"encoderConfig": map[string]string{
			"messageKey":   "message",
			"levelKey":     "level",
			"levelEncoder": "lowercase",
		},
		"development": true,
		"name":        "",
		"logerrLevel": 1,
	}
}
