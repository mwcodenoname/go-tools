package config

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFromYAMLFile(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println("workdir:" + dir)
	// fmt.Println(GetConfig("dns_server_addr_port"))
	// ReadFromYAMLFile("config.yaml")
	fmt.Println(GetConfig("test"))
}
