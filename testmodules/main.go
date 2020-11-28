package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("")
	if err != nil {
		fmt.Println("could not establish connection")
	}
	fmt.Println("Session: ", discord)
}
