package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"math/rand"
	"time"
)

type Game struct {
	Title string `json:"Title"`
	Console string `json:"System"`
	Time string `json:"SDA_Time"`
}

func getRandomGame(fromlist []Game) Game{
	rand.Seed(time.Now().UnixNano())
	return fromlist[rand.Intn(len(fromlist))]
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
    fmt.Println(getRandomGame(games))
}