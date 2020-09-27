package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/adnanbrq/validation"
)

// Config model
type Config struct {
	DbUser string
	DbPass string
	DbName string
	DbHost string
	DbPort int
}

// configValidator model used to validate the config file
type configValidator struct {
	DbUser interface{} `json:"db_user" valid:"required|string"`
	DbPass interface{} `json:"db_pass" valid:"required|string"`
	DbName interface{} `json:"db_name" valid:"required|string"`
	DbHost interface{} `json:"db_host" valid:"required|string"`
	DbPort interface{} `json:"db_port" valid:"required|number"`
}

func (c *configValidator) ToConfig() Config {
	return Config{
		DbUser: c.DbUser.(string),
		DbPass: c.DbPass.(string),
		DbName: c.DbName.(string),
		DbHost: c.DbHost.(string),
		DbPort: int(c.DbPort.(float64)),
	}
}

// ReadConfig will either read config.json
func ReadConfig() Config {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	cfg := configValidator{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	if errs := validation.Validate(cfg); len(errs) > 0 {
		for key, msgs := range errs {
			for _, value := range msgs {
				fmt.Printf("%s %s\n", key, value)
			}
		}
		os.Exit(1)
	}

	return cfg.ToConfig()
}
