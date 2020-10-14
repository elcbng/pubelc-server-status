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
	URL    string
}

var conf ConfigStruct

func init() {
	hello()
	fmt.Printf("\033[1;34m%s", "Initializing\n\n")
	fmt.Println("Reading Config.json :")
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Listen port :", conf.Port)
	fmt.Println("Listen URL :", conf.URL)
	fmt.Println("Debug flag :", conf.Debug)
	fmt.Println("Preset pass :", conf.Passwd)
	fmt.Println("\nInit finished !")
	fmt.Println()
}

func hello() {
	fmt.Printf("\033[1;31m%s", "===Wel Cum 2 Gostat!===\n")
}

func GetDebug() (r int) {
	return conf.Debug
}

func GetPort() string {
	r := strconv.Itoa(conf.Port)
	return r
}

func GetURL() string {
	return conf.URL
}
