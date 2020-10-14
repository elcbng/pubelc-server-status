package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type ConfigStruct struct {
	Debug  int
	Port   int
	Passwd string
}

var conf ConfigStruct

func init() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Listen port :", conf.Port)
	fmt.Println("Debug flag :", conf.Debug)
	fmt.Println("Preset pass :", conf.Passwd, "\n")
}

func GetDebug() (r int) {
	return conf.Debug
}

func GetPort() string {
	r := strconv.Itoa(conf.Port)
	return r
}
