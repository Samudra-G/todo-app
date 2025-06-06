package utils

import "html/template"

// Dict returns a map that can be passed into templates
func Dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, nil
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, nil
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func Add(a, b int) int {
	return a + b
}

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"dict": Dict,
		"add":  Add,
	}
}
