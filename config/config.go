package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	JsonConfigVar *StructConfig
)

func ReadConfig() error {
	if abs, exists := findFile("./config/config.json"); exists {
		log.Println("Config detected - Reading file")

		c, err := ioutil.ReadFile(abs)
		if err != nil {
			return err
		}

		log.Println(string(c))
		err = json.Unmarshal(c, &JsonConfigVar)
		if err != nil {
			return err
		}

		JsonConfigVar.RemoteConfig, err = getConfig()
		if err != nil {
			return err
		}
	}
	return nil
}

func findFile(path string) (string, bool) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", false
	}
	log.Println(abs)

	file, err := os.Open(abs)
	if err != nil {
		return "", false
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return abs, true
}

func UpdateConfig() {
	var err error
	JsonConfigVar.RemoteConfig, err = getConfig()
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func getConfig() (*RemoteConfigStruct, error) {
	rawJson, err := requestConfig()
	if err != nil {
		return &RemoteConfigStruct{}, err
	}
	log.Printf("struct json recieved:\n%v", string(rawJson))
	return decodeGotConfig(rawJson)
}

func decodeGotConfig(rawJson json.RawMessage) (*RemoteConfigStruct, error) {
	var c RemoteConfigStruct
	err := json.Unmarshal(rawJson, &c)
	return &c, err
}

func requestConfig() (json.RawMessage, error) {
	resp, err := http.Get(fmt.Sprintf("http://%v/bindicatorconfig", *JsonConfigVar.JsonConfig.ConfigIpAddress))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
