package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

type jsonobject struct {
    Object ObjectType
}
 
type ObjectType struct {
    Buffer_size int
    Games []Game
}

type Game struct {
	Title string
	System string
	SDATime string
}

func main() {
    //Read in gamelist
    file, e := ioutil.ReadFile("gamelist.json")
    if e != nil {
    	fmt.Printf("There was an error reading the file! %v\n", e)
    	os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    var jsontype jsonobject
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Results: %v\n", jsontype)
}