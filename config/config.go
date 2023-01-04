package config

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

type Work struct {
	Name string
	Desc string
	Excute string
	Duration int
	Args string
}

type Config struct {
	Server struct {
		Mode string
		Port string
	}

	DB map[string]map[string]interface{}

	Work []Work

	Log struct {
		Level string
		Fpath string
		Msize int
		Mage int
		Mbackup int
	}
}

func NewConfig(fpath string) *Config {
	c := new(Config)

	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			fmt.Println(c)
		}
	}
	return c
}