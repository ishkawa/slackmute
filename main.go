package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	for _, token := range strings.Split(os.Getenv("TOKEN"), ",") {
		ids, err := GetChannels(token)
		if err != nil {
			log.Fatal(err)
		}

		err = MuteChannels(token, ids)
		if err != nil {
			log.Fatal(err)
		}
	}
}
