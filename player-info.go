package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const PlayersURL = "https://raw.githubusercontent.com/BurntSushi/nflgame/master/nflgame/players.json"

type PlayerInfo struct {
	Birthdate  string
	College    string
	FirstName  string `json:"first_name"`
	FullName   string `json:"full_name"`
	GsisId     string `json:"gsis_id"`
	GsisName   string `json:"gsis_name"`
	Height     int
	LastName   string `json:"last_name"`
	ProfileId  int    `json:"profile_id"`
	ProfileURL string `json:"profile_url"`
	Weight     int
	YearsPro   int `json:"years_pro"`
}

func jsonToStruct(data json.RawMessage) *PlayerInfo {
	var players PlayerInfo
	readerType := bytes.NewReader(data)

	if err := json.NewDecoder(readerType).Decode(&players); err != nil {
		log.Fatal(err)
	}
	return &players
}

func main() {
	var playersData map[string]json.RawMessage
	response, error := http.Get(PlayersURL)
	if error != nil {
		log.Fatal(error)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	unmarshalErr := json.Unmarshal(body, &playersData)
	response.Body.Close()
	if err != nil {
		log.Fatal(unmarshalErr)
	}

	var playersMap = make(map[string]*PlayerInfo)

	for key, val := range playersData {
		playersMap[key] = jsonToStruct(val)
	}

	for _, player := range playersMap {
		fmt.Printf("# %10s => %10s %10s\n", player.FullName, player.College, player.ProfileURL)
	}
}
