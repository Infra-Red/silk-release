package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"code.cloudfoundry.org/go-db-helpers/db"
)

type Config struct {
	UnderlayIP     string    `json:"underlay_ip"`
	SubnetRange    string    `json:"subnet_range"`
	SubnetMask     int       `json:"subnet_mask"`
	Database       db.Config `json:"database"`
	LocalStateFile string    `json:"local_state_file"`
}

func LoadConfig(filePath string) (Config, error) {
	var cfg Config
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf("reading file %s: %s", filePath, err)
	}

	err = json.Unmarshal(contents, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("unmarshaling contents: %s", err)
	}
	return cfg, nil
}
