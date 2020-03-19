package main

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func prepareTemplateValues(envName string) map[string]interface{} {

	var templateValues map[string]interface{}
	decodedValues := []byte(envName)

	err := json.Unmarshal(decodedValues, &templateValues)
	if err != nil {
		log.Fatalln("error:", err)
	}

	return templateValues
}

func getConf(fileName string) map[string]interface{} {

	var conf map[string]interface{}
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return conf
}

func mergeValues(a, b map[string]interface{}) map[string]interface{} {

	for k, v := range b {
		a[k] = v
	}

	return a
}
