package main

import (
	"fmt"
	"go-votes/internal/votes"
	"io"
	"log"
	"net/http"
)

const (
	PORT = "9090"
)

func registerRoutes() {
	http.HandleFunc("/vote", votes.PostVote)
}

func startServer() {
	http.ListenAndServe(":"+PORT, nil)
}

func main() {
	fmt.Println("Attempting to running at port: ", PORT, "!!")

	registerRoutes()
	startServer()
}

func requestSomeService(url string) (parsed string, ok bool) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		return err.Error(), false
	}

	fmt.Println(res.StatusCode)

	if res.StatusCode != http.StatusOK {
		message := fmt.Sprint("Error requesting to "+url+" with ", (res.StatusCode))
		fmt.Println(message)
		return message, false
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
		return err.Error(), false
	}

	parsed = string(body)

	fmt.Println(parsed)

	return parsed, true
}
