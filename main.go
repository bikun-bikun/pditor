package main

import (
	"fmt"
	"os"

	"pditor/commands"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pditor"
	app.Usage = "PDF etitor"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "rotation",
			Usage:   "PDF File rotation",
			Aliases: []string{"r"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "pages, p",
					Usage: "default all page rotate, Odd page option=o Even page option=e",
					Value: "all",
				},
				cli.Int64Flag{
					Name:  "left-angle, la",
					Usage: "This flag is left rotation angle setting values{0, 90, 180, 270}",
					Value: 0,
				},
				cli.Int64Flag{
					Name:  "right-angle, ra",
					Usage: "This flag is right rotation angle setting values{0, 90, 180, 270}",
					Value: 0,
				},
				cli.StringFlag{
					Name:  "path, p",
					Usage: "This flag is file path setting. Set target pdf file path",
					Value: "",
				},
			},
			Action: rotation,
		},
		commands.DemoCreate,
		commands.DemoList,
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("pditor! I say!")
		return nil
	}

	app.Run(os.Args)
}

func rotation(c *cli.Context) error {
	fmt.Println("Rotate PDF!!!")
	return nil
}
