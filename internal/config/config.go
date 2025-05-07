package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return Write(*cfg)
}

func Write(cfg Config) error {
	fileName, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println(err)
		return Config{}, err
	}

	result := Config{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return Config{}, err
	}

	return result, nil
}

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFilePath := dir + configFileName

	return configFilePath, nil
}
