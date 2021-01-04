package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func greet(c *cli.Context) error {
	fmt.Printf("Hello %q", c.Args().Get(0))
	return nil
}

func main() {
	app := &cli.App{
		Action: greet,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
