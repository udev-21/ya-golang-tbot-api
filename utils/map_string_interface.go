package utils

import "encoding/json"

func ConvertToMapStringInterface(obj interface{}) (map[string]interface{}, error) {
	raw, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(raw, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
