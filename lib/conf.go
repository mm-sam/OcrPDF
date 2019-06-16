package lib

import (
	"github.com/apex/log"
	"github.com/json-iterator/go"
	"os"
)

type Config struct {
	Driver    string `json:"driver"`
	Vendor    map[string]interface{} `json:"vendor"`
	save_path string
}

func NewConfig(filename string) (err error, c *Config) {
	c = &Config{}
	c.save_path = filename
	err = c.load(filename)
	return
}

func (c *Config) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.WithError(err)
		return err
	}
	defer file.Close()
	decoder := jsoniter.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		log.WithError(err)
	}
	return err
}

func (c *Config) Save() error {
	file, err := os.Create(c.save_path)
	if err != nil {
		log.WithError(err)
		return err
	}
	defer file.Close()
	data, err := jsoniter.MarshalIndent(c, "", "    ")
	if err != nil {
		log.WithError(err)
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		log.WithError(err)
	}
	return err
}
