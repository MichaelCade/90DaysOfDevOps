package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"strings"
	"time"
)

func storeString(input string) error {
	// Connect to the database
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	defer db.Close()
	// Insert the random number into the database
	_, err = db.Exec("INSERT INTO requestor(random_string) VALUES(?)", input)
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
	result, err := db.Query("SELECT * FROM requestor WHERE random_string = ?", input)
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

func getStringFromGenerator() {
	// make a request to the generator
	// save sthe string to db
	r, err := http.Get("http://generator-service:8080/new")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("body from generator: %s", string(body)))

	storeString(string(body))

}

func main() {
	// setup a goroutine loop calling the generator every minute, saving the result in the DB

	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				getStringFromGenerator()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Create a new HTTP server
	server := &http.Server{
		Addr: ":8080",
	}

	err := createRequestordb()
	if err != nil {
		panic(err.Error())
	}

	// Handle requests to /generate
	http.HandleFunc("/check/", func(w http.ResponseWriter, r *http.Request) {
		// get the value after check from the url
		id := strings.TrimPrefix(r.URL.Path, "/check/")

		// check if it exists in the db
		err := getStringFromDB(id)
		if err != nil {
			http.Error(w, "string not found", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "string found", http.StatusOK)
	})

	// Start the server
	fmt.Println("Listening on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func createRequestordb() error {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	if err != nil {
		return err
	}
	defer db.Close()

	// try to create a table for us
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS requestor(random_string VARCHAR(100))")

	return err
}
