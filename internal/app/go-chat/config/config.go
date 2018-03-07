package config

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	defaultIP   = "127.0.0.1"
	defaultPort = 5000
	defaultLog  = "logs/go-chat.log"
)

// Server configuration structure.
type Server struct {
	// IP speficies address on which server will listen
	IP string `yaml:"ip"`
	// Port specifies port on which server will listen
	Port uint16 `yaml:"port"`
	// Log points to exact file which will be used to store all logs
	Log string `yaml:"log"`
}

// Configuration structure which handle complete settings.
type Configuration struct {
	// Server holds configuration specific for telnet server
	Server Server `yaml:"server"`
}

// NewDefaultConfig initializes configuration with default values.
func NewDefaultConfig() Configuration {
	return Configuration{
		Server{
			IP:   defaultIP,
			Port: defaultPort,
			Log:  defaultLog,
		},
	}
}

// LoadConfig creates new configuration based on the content of specified file.
// When file doesn't exist or it's unparsable it will use default configuration.
func LoadConfig(filename string) Configuration {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Infof("Couldn't open %s configuration file. Using default settings.", filename)
		return NewDefaultConfig()
	}

	var c Configuration
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		log.Infof("Couldn't read %s configuration file. Err: %s. Using default settings.", filename, err)
		return NewDefaultConfig()
	}

	return c
}
