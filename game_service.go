package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

type Game struct {
	Title string `json:"Title"`
	Console string `json:"System"`
	Time string `json:"SDA_Time"`
}

func main() {
	//Read in gamelist
    file, e := ioutil.ReadFile("gamelist.json")
    if e != nil {
    	fmt.Printf("There was an error reading the file! %v\n", e)
    	os.Exit(1)
    }

    var games []Game
    json.Unmarshal(file, &games)
    fmt.Println(games[4].Console)
}