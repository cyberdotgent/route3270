package main

import "github.com/BurntSushi/toml"

type userInfo struct {
	Password string
	Servers []string
}

type serverInfo struct {
	Hostname string
	Port int32
	Description string
	Mnemoric string
}

type tomlConfig struct {
	Port int32
	Servers map[string]serverInfo
	Users map[string]userInfo
}

func parseConfig(filename string) tomlConfig {
	var config tomlConfig

	if _, err := toml.DecodeFile(filename, &config); err != nil {
		panic(err)
	}

	return config
}