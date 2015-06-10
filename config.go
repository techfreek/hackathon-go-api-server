package main

import (
	"os"
	"fmt"
	"encoding/json"
	"strings"
)

type Eventbrite struct {
	Event_id string `json:"event_id"`
	Oathtoken string `json:"oathtoken"`
}

var Config struct {
	Event Eventbrite `json:"eventbrite"`
	Img_dir string `json:"imagesDir"`
}

var ConfigFileName = "config.json";

//If no config file has been created, create one
func create_Config() {
	conf, _ := os.Create(ConfigFileName)
	defer conf.Close();

	fmt.Println("Eventbrite Configuration:");
	fmt.Printf("Event ID:  "); fmt.Scan(Config.Event.Event_id);
	fmt.Printf("Oathtoken: "); fmt.Scan(Config.Event.Oathtoken);
	fmt.Print("Misc Configuration: ");
	fmt.Printf("Img Dir:   "); fmt.Scan(Config.Img_dir);
	
	str, _ := json.Marshal(config);
	_, err := conf.Write(str);
	
	if err != nil {
		fmt.Println(err);
	}
	
}

//loads config file if it exists, otherwise it creates a new config
func Init_Config() {
	conf, err := os.Open(ConfigFileName)
	defer conf.Close();

	if err != nil {
		if os.IsNotExit(err) {
			fmt.Println(err);
		} else {
			return create_config()
		}
	}

	//start decoding
	decoder := json.NewDecoder(conf);
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println(err);
	}
}