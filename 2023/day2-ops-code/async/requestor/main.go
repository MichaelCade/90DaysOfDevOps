package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	nats "github.com/nats-io/nats.go"
	"time"
)

func storeString(input string) error {
	// Connect to the database
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	defer db.Close()
	// Insert the random number into the database
	_, err = db.Exec("INSERT INTO requestor_async(random_string) VALUES(?)", input)
	if err != nil {
		return err
	}

	fmt.Printf("Random string %s has been inserted into the database\n", input)
	return nil
}

func getStringFromDB(input string) error {
	// see if the string exists in the db, if so return nil
	// if not, return an error
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	defer db.Close()
	result, err := db.Query("SELECT * FROM requestor_async WHERE random_string = ?", input)
	if err != nil {
		return err
	}

	for result.Next() {
		var randomString string
		err = result.Scan(&randomString)
		if err != nil {
			return err
		}
		if randomString == input {
			return nil
		}
	}

	return errors.New("string not found")
}

func main() {

	err := createRequestordb()
	if err != nil {
		panic(err.Error())
	}
	// setup a goroutine loop calling the generator every minute, saving the result in the DB

	nc, _ := nats.Connect("nats://my-nats:4222")
	defer nc.Close()

	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				nc.Publish("generator", []byte(""))
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	nc.Subscribe("generator_reply", func(msg *nats.Msg) {
		err := storeString(string(msg.Data))
		if err != nil {
			print(err)
		}
	})

	nc.Subscribe("confirmation", func(msg *nats.Msg) {
		err := getStringFromDB(string(msg.Data))
		if err != nil {
			print(err)
		}
		nc.Publish("confirmation_reply", []byte(string(msg.Data)))
	})
	// create a goroutine here to listen for messages on the queue to check, see if we have them

	for {
		time.Sleep(10 * time.Second)
	}

}

func createRequestordb() error {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	if err != nil {
		return err
	}
	defer db.Close()

	// try to create a table for us
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS requestor_async(random_string VARCHAR(100))")

	return err
}
