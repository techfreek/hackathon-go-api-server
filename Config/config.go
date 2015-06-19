package config

import (
	"os"
	"fmt"
	"encoding/json"
)

type Eventbrite struct {
	Event_id string `json:"event_id"`
	Oathtoken string `json:"oathtoken"`
}

var Config struct {
	Event Eventbrite `json:"eventbrite"`
	Img_dir string `json:"imagesDir"`
	Max_Spots int `json:"spots"`
}

var ConfigFileName = "../config.json";

//If no config file has been created, create one
func create_Config() {
	conf, _ := os.Create(ConfigFileName)
	defer conf.Close();

	fmt.Println("Eventbrite Configuration:");
	fmt.Printf("Event ID:  "); fmt.Scan(Config.Event.Event_id);
	fmt.Printf("Oathtoken: "); fmt.Scan(Config.Event.Oathtoken);
	fmt.Print("Misc Configuration: ");
	fmt.Printf("Img Dir:   "); fmt.Scan(Config.Img_dir);
	fmt.Printf("Max Spots:   "); fmt.Scan(Config.Max_Spots);
	str, _ := json.Marshal(Config);
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
		if os.IsNotExist(err) {
			fmt.Println(err);
		} else {
			create_Config()
			return
		}
	}

	//start decoding
	decoder := json.NewDecoder(conf);
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println(err);
	}
}