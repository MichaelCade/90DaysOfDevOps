package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	nats "github.com/nats-io/nats.go"
	"math/rand"
	"time"
)

func generateAndStoreString() (string, error) {
	// Connect to the database
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Generate a random string
	// Define a string of characters to use
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generate a random string of length 10
	randomString := make([]byte, 64)
	for i := range randomString {
		randomString[i] = characters[rand.Intn(len(characters))]
	}

	// Insert the random number into the database
	_, err = db.Exec("INSERT INTO generator_async(random_string) VALUES(?)", string(randomString))
	if err != nil {
		return "", err
	}

	fmt.Printf("Random string %s has been inserted into the database\n", string(randomString))
	return string(randomString), nil
}

func main() {
	err := createGeneratordb()
	if err != nil {
		panic(err.Error())
	}

	nc, _ := nats.Connect("nats://my-nats:4222")
	defer nc.Close()

	nc.Subscribe("generator", func(msg *nats.Msg) {
		s, err := generateAndStoreString()
		if err != nil {
			print(err)
		}
		nc.Publish("generator_reply", []byte(s))
		nc.Publish("confirmation", []byte(s))
	})

	nc.Subscribe("confirmation_reply", func(msg *nats.Msg) {
		stringReceived(string(msg.Data))
	})
	// Subscribe to the queue
	// when a message comes in call generateAndStoreString() then put the string on the
	// reply queue. also add a message onto the confirmation queue

	// subscribe to the confirmation reply queue
	// when a message comes in call

	for {
		time.Sleep(1 * time.Second)
	}

}

func createGeneratordb() error {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	if err != nil {
		return err
	}
	defer db.Close()

	// try to create a table for us
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS generator_async(random_string VARCHAR(100), seen BOOLEAN, requested BOOLEAN)")

	return err
}

func stringReceived(input string) {

	// Connect to the database
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	if err != nil {
		print(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE generator_async SET requested = true WHERE random_string = ?", input)
	if err != nil {
		print(err)
	}
}
