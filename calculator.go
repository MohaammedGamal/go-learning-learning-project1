package main

import (
	"fmt"
	"log"
	"log/slog"
	"math/rand/v2"
	"os"
)

func main() {
	outputText("gemy")
	mr, name := returnMrName("Mr", "Mohamed")
	fmt.Println("Hello", mr, name)
	output := getInput("test data")
	fmt.Println("Output from a user input :", output)
}

func outputText(text string) {
	fmt.Println("Output : ", text)
	slog.Info("Output text", "user", os.Getenv("USERNAME"))
	slog.Info("Extra", "extra info", "this is extra info")
	fmt.Println("Generated random number")
	log.Println("Generated random number", returnRandomNumber())
}

func returnRandomNumber() int {
	return rand.IntN(100)
}

func returnMrName(mr, name string) (string, string) {
	return mr, name
}

func getInput(info string) string {
	fmt.Print("Please enter ", info, ": ")
	var input string
	fmt.Scanln(&input)
	return input
}
