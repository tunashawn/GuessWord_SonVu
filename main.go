package main

import (
	"GuessWord_SonVu/api"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("Guessing word...")
	mode := os.Args[1]
	size := 5 //default as specified in API doc
	seed := 0
	if len(os.Args) >= 3 {
		s := os.Args[2]
		var err error
		size, err = strconv.Atoi(s)
		if size < 2 {
			log.Fatal("size must be greater than 2")
		}
		if err != nil {
			log.Fatal(fmt.Errorf("error converting size to int: %s", err))
		}
	}

	if mode == api.RANDOM {
		if len(os.Args) == 4 {
			s := os.Args[3]
			var err error
			seed, err = strconv.Atoi(s)

			if err != nil {
				log.Fatal(fmt.Errorf("error converting seed to int: %s", err))
			}
		} else {
			log.Fatal("you need to provide a random seed")
		}
	}

	res, err := api.GuessWord(mode, size, seed)
	if err != nil {
		log.Fatal(fmt.Errorf("error guessing daily: %s", err))
	}
	fmt.Printf("mode = %s, size = %d, correct word: %s", mode, size, res)
}
