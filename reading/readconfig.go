package main

import (
	"encoding/json"
	"os"
	// "strings"
)

type ConfigData struct{
	Username string
	AdditionalProducts []Product
}

var Config ConfigData

func LoadConfig() (err error){
	// data, err := os.ReadFile("config.json")
	// if (err == nil){
	// 	decoder := json.NewDecoder(strings.NewReader(string(data)))
	// 	err = decoder.Decode(&Config)
	// }
	file, err := os.Open("config.json")
	if (err == nil){
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config)
	}
	return
}

func init(){
	err := LoadConfig()
	if (err != nil){
		Printfln("Error loading config: %v", err.Error())
	}
	Printfln("Username: %v", Config.Username)
	Products = append(Products, Config.AdditionalProducts...)
}