package commands

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/urfave/cli"
	"fmt"
)

const (
	commandName string = "demoCreate"
	usage       string = "demoCreate Command create a new \"hello.pdf\" File. "
)

var DemoList = cli.Command{
	Name: "list",
	Usage: "list is Linux Command ls",
	Aliases: []string{"ls"},
	Action: ls,
}

var DemoCreate = cli.Command{
	Name:    commandName,
	Usage:   usage,
	Aliases: []string{"d"},
	Action:  demo,
}

func ls(c *cli.Context) error{
	fmt.Printf("c.Args() : %+v\n", c.Args())
	return nil
}

// Demo ...
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
