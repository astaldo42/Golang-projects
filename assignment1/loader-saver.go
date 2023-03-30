package main

import "io/ioutil"
import "encoding/json"

func SaveDataToJSON(filename string, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

func LoadDataFromJSON(filename string, data interface{}) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, data)
}
