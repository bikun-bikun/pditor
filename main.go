package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pditor"
	app.Usage = "make an explosive entrance"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		fmt.Println("pditor! I say!")
		return nil
	}

	app.Run(os.Args)
}
