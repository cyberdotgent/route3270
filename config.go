package main

import "github.com/BurntSushi/toml"

type userInfo struct {
	Password string
	DestinationHost string
	DestinationPort int32
}

type tomlConfig struct {
	Port int32
	Users map[string]userInfo
}

func parseConfig(filename string) tomlConfig {
	var config tomlConfig

	if _, err := toml.DecodeFile(filename, &config); err != nil {
		panic(err)
	}

	return config
}