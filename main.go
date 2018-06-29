package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/ILUGD/splatter/theme"
)

//Date  Data Structure for containing the Meetup Date
type Date struct {
	D int `json:"d"`
	M int `json:"m"`
	Y int `json:"y"`
}

//Time  Data Structure for containing Timings
type Time struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

//Document  Parent Data Structure for the Poster Details
type Document struct {
	Title         string   `json:"title"`
	EventDate     Date     `json:"date"`
	Venue         string   `json:"venue"`
	Timings       Time     `json:"time"`
	Background    string   `json:"background"`
	SponsorLogos  []string `json:"logos"`
	GroupWebsites []string `json:"websites"`
}

var configFlag = flag.String("config", "NULL", "Path for the JSON config")
var themeFlag = flag.String("theme", "NULL", "path for the Theme YAML file")

func main() {
	flag.Parse()

	if *configFlag == "NULL" || *themeFlag == "NULL" {
		fmt.Printf("Usage: %s -config <config.json> -theme <theme.yaml>")
	}

	configFile, err := os.Open(*configFlag)
	defer configFile.Close()
	must(err)

	theme.ParseFile(*themeFlag)

	var imageDetails Document

	bytes, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(bytes, &imageDetails)

	configDir, _ := path.Split(*configFlag)
	imageDetails.Background = path.Join(configDir, imageDetails.Background)
	must(err)

	GeneratePoster(imageDetails)
}

//must  Function for handling errors
func must(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
