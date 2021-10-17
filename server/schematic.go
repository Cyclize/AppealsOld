package main

import (
	t "github.com/TyphoonMC/TyphoonCore"
)

type SpawnConfig struct {
	Schematic *string     `json:"schematic"`
	Location  *t.Location `json:"location"`
}

type Config struct {
	Spawn *SpawnConfig `json:"spawn"`
}

var (
	config Config
	spawn  *t.Map
)

func loadConfig(core *t.Core) {
	core.GetConfig(&config)

	if config.Spawn != nil && config.Spawn.Schematic != nil {
		m, err := t.LoadSchematic(*config.Spawn.Schematic)
		if err != nil {
			panic(err)
		}
		spawn = m
	}
}
