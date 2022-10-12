package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	test()
	dir, _ := os.Getwd()
	fmt.Println(dir)
}

func test() {
	modFileDir := "data/github.com/!azure/azure-amqp-common-go@v1.1.4"

	dir, _ := os.Getwd()
	fmt.Println(dir)

	os.Chdir(modFileDir)
	cmd := exec.Command("sh", "-c", "ls")
	data, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to exec go mod graph of ", err)
	}
	fmt.Println(string(data))
}
