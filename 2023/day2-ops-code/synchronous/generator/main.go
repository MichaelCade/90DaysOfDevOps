package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/http"
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
	_, err = db.Exec("INSERT INTO generator(random_string) VALUES(?)", string(randomString))
	if err != nil {
		return "", err
	}

	fmt.Printf("Random string %s has been inserted into the database\n", string(randomString))
	return string(randomString), nil
}

func main() {
	// Create a new HTTP server
	server := &http.Server{
		Addr: ":8080",
	}

	err := createGeneratordb()
	if err != nil {
		panic(err.Error())
	}

	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				checkStringReceived()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Handle requests to /generate
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		// Generate a random number
		randomString, err := generateAndStoreString()
		if err != nil {
			http.Error(w, "unable to generate and save random string", http.StatusInternalServerError)
			return
		}

		print(fmt.Sprintf("random string: %s", randomString))
		w.Write([]byte(randomString))
	})

	// Start the server
	fmt.Println("Listening on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func createGeneratordb() error {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	if err != nil {
		return err
	}
	defer db.Close()

	// try to create a table for us
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS generator(random_string VARCHAR(100), seen BOOLEAN)")

	return err
}

func checkStringReceived() {
	// get a list of strings from database that dont have the "seen" bool set top true
	// loop over them and make a call to the requestor's 'check' endpoint and if we get a 200 then set seen to true

	// Connect to the database
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql")
	if err != nil {
		print(err)
	}
	defer db.Close()

	// Insert the random number into the database
	results, err := db.Query("SELECT random_string FROM generator WHERE seen IS NOT true")
	if err != nil {
		print(err)
	}

	// loop over results
	for results.Next() {
		var randomString string
		err = results.Scan(&randomString)
		if err != nil {
			print(err)
		}

		// make a call to the requestor's 'check' endpoint
		// if we get a 200 then set seen to true
		r, err := http.Get("http://requestor-service:8080/check/" + randomString)
		if err != nil {
			print(err)
		}
		if r.StatusCode == 200 {
			_, err = db.Exec("UPDATE generator SET seen = true WHERE random_string = ?", randomString)
			if err != nil {
				print(err)
			}
		} else {
			fmt.Println(fmt.Sprintf("Random string has not been received by the requestor: %s", randomString))
		}
	}
}
