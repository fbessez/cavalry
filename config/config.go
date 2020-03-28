// Package config will load the ENV variables for constants we want to set dynamically.
package config

import (
	"github.com/kelseyhightower/envconfig"
)

// CONSTANTS is the instance of the Constants struct
var CONSTANTS Constants

type (
	// Constants is a struct that holds ... constants
	Constants struct {
		APIToken 	string `yaml:"api_token"   envconfig:"api_token"  required:"true"`
		UserID 		string `yaml:"user_id" 	   envconfig:"user_id"    required:"true"`
		StartHour int  `yaml:"start_hour"  envconfig:"start_hour" required: "true"`
		EndHour 	int  `yaml:"end_hour"    envconfig:"end_hour"   required: "true"`
	}
)

func init() {
	err := envconfig.Process("cavalry", &CONSTANTS)
	if err != nil {
		panic(err)
	}
}
