package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Dictinary *map[interface{}]interface{}

func LoadLocales(path string)  error {
	content,err := ioutil.ReadFile(path)
	if err != nil{
		return err
	}
	//var m map[interface{}]interface{}
	//m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(content), &Dictinary)

	if err != nil{
		return err
	}
	//Dictinary = &m
	return nil
}

func T(key string) string{
	//TODO

	return ""
}