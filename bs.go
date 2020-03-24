package main

import (
	"log"
	"os"
	"bs/cmd"

)

func main() {
	app := cmd.NewCmdApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}