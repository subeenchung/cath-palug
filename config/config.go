package config

import (
	"log"
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

type MainConfig struct {
	App AppConfig `toml:"app"`
	DB DBConfig `toml:"db"`
	Smtp SmtpConfig `toml:"smtp"`		
}

type AppConfig struct {
	ListeningPort int `toml:"listening_port"`
	SigningKey string `toml:"signing_key"`
	Debug bool	`toml:"debug"`
}

type DBConfig struct {
	Type string `toml:"type"`
	IP string `toml:"ip"`
	Port int `toml:"port"`
	Dbname string `toml:"db_name"`
	User string `toml:"user"`
	Password string `toml:"password"`
}

type SmtpConfig struct {
	Server1	string `toml:"server1"`
	Server2 string `toml:"server2"`
}

func LoadConfig(path string)MainConfig {
	var mc MainConfig
	fb, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if _, err := toml.Decode(string(fb), &mc); err != nil {
		log.Fatal(err)
	}
	return mc
}
