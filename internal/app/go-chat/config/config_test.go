package config

import (
	"path/filepath"
	"testing"
	//"github.com/diltram/go-chat/test/fixtures/file"
)

//TestLoadConfig creates configuration based on the test file.
//All parameters should be correctly loaded and set.
func TestLoadConfig(t *testing.T) {
	expected := Configuration{
		Server: Server{
			IP:   "0.0.0.0",
			Port: 5555,
			Log:  "logs/go-chat.log",
		},
	}
	conf := LoadConfig(filepath.Join("testdata", "config.yml"))

	if conf != expected {
		t.Errorf("Generated config is different. Expected: %+v, got: %+v", expected, conf)
	}
}

//TestLoadConfigNoFile tests default config as file doesn't exist.
func TestLoadConfigNoFile(t *testing.T) {
	expected := NewDefaultConfig()
	conf := LoadConfig("non_existing_path")

	if conf != expected {
		t.Errorf("Generated config is different. Expected: %+v, got: %+v", expected, conf)
	}
}

//TestLoadConfigNoFile tests default config as file doesn't exist.
func TestLoadConfigCorruptedData(t *testing.T) {
	expected := NewDefaultConfig()
	conf := LoadConfig(filepath.Join("testdata", "corrupted_config.yml"))

	if conf != expected {
		t.Errorf("Generated config is different. Expected: %+v, got: %+v", expected, conf)
	}
}
