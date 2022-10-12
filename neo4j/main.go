package main

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var (
	NEO4J_URI      = ""
	NEO4J_USERNAME = ""
	NEO4J_PASSWORD = ""
	NEO4J_DB       = ""
)

func GetDriver() neo4j.Driver {
	driver, err := neo4j.NewDriver(
		NEO4J_URI,
		neo4j.BasicAuth(NEO4J_USERNAME, NEO4J_PASSWORD, ""),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return driver
}

func main() {
	driver := GetDriver()
	defer driver.Close()

	data, err := helloWorld(driver)
	fmt.Println(err)
	fmt.Println(data)
}

func helloWorld(driver neo4j.Driver) (string, error) {
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite, DatabaseName: NEO4J_DB})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run("", nil)
		// result, err := transaction.Run(
		// 	"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
		// 	map[string]interface{}{"message": "hello, world"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return "", nil
	// return greeting.(string), nil
}
