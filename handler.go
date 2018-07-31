package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	toml "github.com/pelletier/go-toml"
)

var env string
var configFile string
var config *toml.Tree

func initConfig() {
	var err error
	config, err = toml.LoadFile(configFile)
	if err != nil {
		panic(err)
	}
}

func getTeam(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	file := resource + "/team.json"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	str := string(b)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(str), &data)
	if err != nil {
		panic(err)
	}
	renderJSON(w, data)
}

func updateTeam(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	b, _ := ioutil.ReadAll(req.Body)
	err := ioutil.WriteFile(resource+"/team.json", b, 0644)
	if err != nil {
		panic(err)
	}
}

func getGame(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	file := resource + "/game.json"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	str := string(b)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(str), &data)
	if err != nil {
		panic(err)
	}
	renderJSON(w, data)
}

func updateGame(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	b, _ := ioutil.ReadAll(req.Body)
	err := ioutil.WriteFile(resource+"/game.json", b, 0644)
	if err != nil {
		panic(err)
	}
}

func getApostleTeam(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	file := resource + "/apostleTeam.json"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	str := string(b)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(str), &data)
	if err != nil {
		panic(err)
	}
	renderJSON(w, data)
}

func updateApostleTeam(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	b, _ := ioutil.ReadAll(req.Body)
	err := ioutil.WriteFile(resource+"/apostleTeam.json", b, 0644)
	if err != nil {
		panic(err)
	}
}
