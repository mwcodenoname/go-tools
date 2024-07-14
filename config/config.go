package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var AppConfig map[string]string

func init() {
	AppConfig = make(map[string]string)
	ReadFromYAMLFile("config.yaml")
}

func GetConfig(key string) string {
	return AppConfig[key]
}

func ReadFromYAMLFile(filename string) {
	workdir := os.Getenv("XCONFIG_WORKDIR")
	fmt.Println("XCONFIG_WORKDIR:" + workdir)
	if workdir == "" {
		workdir, _ = os.Getwd()
	} else {
		if strings.HasSuffix(workdir, string(os.PathSeparator)) {
			workdir = strings.TrimSuffix(workdir, string(os.PathSeparator))
		}
	}
	configPath := filepath.Join(workdir, "config.yaml")

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
