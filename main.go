package main

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
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
		{
			Name:    "demoCreate",
			Aliases: []string{"d"},
			Action:  demo,
		},
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

func demo(c *cli.Context) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.MoveTo(20, 20)
	pdf.LineTo(170, 20)
	pdf.ClosePath()
	pdf.SetLineWidth(1)
	pdf.DrawPath("D")
	pdf.SetFont("Arial", "B", 16)
	pdf.Text(40, 50, "Hello, world")
	pdf.OutputFileAndClose("hello.pdf")
	return nil
}
