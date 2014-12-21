package fgms

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//= A Host Row
type JSON_HostConf struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Comment string `json:"comment"`
}

//= Main Server Configuration
type JSON_ServerConf struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
	IsHub   bool   `json:"is_hub"`

	LogFile string `json:"log_file"`

	TelnetPort int `json:"telnet_port"`

	PlayerExpiresSecs int `json:"player_expires"`
	OutOfReachNm      int `json:"out_of_reach"`

	Tracked        bool   `json:"tracked"`
	TrackingServer string `json:"tracking_server"`
}

//= Whole Payload from File
type Config struct {
	Server     JSON_ServerConf `json:"server"`
	Relays     []JSON_HostConf `json:"relays"`
	Crossfeeds []JSON_HostConf `json:"crossfeeds"`
	Blacklists []string        `json:"blacklists"`
}

// Read a config file and set internal variables accordingly.
func LoadConfig(config_path string) (Config, error) {

	var conf Config

	//configFilePath := "/home/gogo/src/github.com/FreeFLighSim/go-fgms/fgms_example.json"

	// Read file
	filebyte, err := ioutil.ReadFile(config_path)
	if err != nil {
		log.Fatal("Could not read JSON config file: `" + config_path + "` ")
		return conf, err
	}
	// Parse JSON

	err = json.Unmarshal(filebyte, &conf)
	if err != nil {
		log.Fatalln("JSON Decode Error from: ", config_path, err)
		return conf, err
	}

	return conf, nil
}
