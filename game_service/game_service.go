package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"math/rand"
	"time"
	"net/http"
)

var games []Game

type Game struct {
	Title string `json:"Title"`
	Console string `json:"System"`
	Time string `json:"SDA_Time"`
}

func getRandomGame(fromlist []Game) Game{
	rand.Seed(time.Now().UnixNano())
	return fromlist[rand.Intn(len(fromlist))]
}

func init() {
	//Read in gamelist
    file, e := ioutil.ReadFile("gamelist.json")
    if e != nil {
    	fmt.Printf("There was an error reading the file! %v\n", e)
    	os.Exit(1)
    }

    json.Unmarshal(file, &games)

	http.HandleFunc("/", serveGame)
}

func serveGame(w http.ResponseWriter, r *http.Request) {
	var game = getRandomGame(games)
	var gameString = "You should try "+game.Title+", for "+game.Console+
	". The fastest time recorded on SDA for this game is "+game.Time+".";
	fmt.Fprintf(w, gameString)
}

/*func main() {
	
    fmt.Println(getRandomGame(games))
}*/