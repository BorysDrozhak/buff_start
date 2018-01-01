package main

import (
	"log"

	"BorysDrozhak/buffalo_test/coke/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
