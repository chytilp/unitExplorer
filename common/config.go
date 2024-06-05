package common

import (
	"io"
	"os"

	toml "github.com/pelletier/go-toml"
)

var config *Config

type Source struct {
	Name     string
	SourceId int
}

type Config struct {
	Sources      []Source
	ApiUrl       string
	DatabaseFile string
}

func (c *Config) FindSourceId(name string) *int {
	for _, source := range c.Sources {
		if source.Name == name {
			return &source.SourceId
		}
	}
	return nil
}

func GetConfig() *Config {
	if config == nil {
		config = read()
	}
	return config
}

func read() *Config {
	file, err := os.Open("config.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config Config

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
