package commands

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/urfave/cli"
)

const (
	commandName string = "demoCreate"
	usage       string = "demoCreate Command create a new \"hello.pdf\" File. "
)

// DemoCreate return cli.Command
func DemoCreate() cli.Command {
	demoCreate := cli.Command{

		Name:    commandName,
		Usage:   usage,
		Aliases: []string{"d"},
		Action:  demo,
	}

	return demoCreate
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
