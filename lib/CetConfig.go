package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type schoolConfig struct {
	SchoolName    string
	MangerType    string
	MangerURL     string
	CalendarFirst string
}

type Config struct {
	School schoolConfig
}

func ReadConfig() Config {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
	}
	var conf Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println(err)
	}
	return conf
}
