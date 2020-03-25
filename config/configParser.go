package config

import (
	"encoding/json"
	"io/ioutil"
)

type configData struct {
	Address, ListenAddress, DbHost, DbName string
}

func getConfigData() *configData {
	file, _ := ioutil.ReadFile("config/config.json")
	configData := configData{}
	_ = json.Unmarshal([]byte(file), &configData)
	return &configData
}

//GetListenAddress ...
func GetListenAddress() string {
	return getConfigData().ListenAddress
}

//GetAddress ...
func GetAddress() string {
	return getConfigData().Address
}

//GetDbHost ...
func GetDbHost() string {
	return getConfigData().DbHost
}

//GetDbName ...
func GetDbName() string {
	return getConfigData().DbName
}
