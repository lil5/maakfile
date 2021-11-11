package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lil5/maakfile/global"
	"github.com/lil5/maakfile/global/classes"
	"github.com/lil5/maakfile/initialize"
	"github.com/lil5/maakfile/list"
	"github.com/lil5/maakfile/run"
	"github.com/urfave/cli/v2"
)

const version = "1.0.0"

func main() {
	app := &cli.App{
		Name:    "maak",
		Usage:   "Script runner with parallel & watch options",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "generate Maakfile.toml",
				Action: func(c *cli.Context) error {
					err := execInit()
					returnError(&err)
					return nil
				},
			},
			{
				Name:  "run",
				Usage: "run script from Maakfile.toml",
				Action: func(c *cli.Context) error {
					name := c.Args().First()
					err := execRun(name)
					returnError(&err)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls", "ll"},
				Usage:   "list all available scripts",
				Action: func(c *cli.Context) error {
					err := execList()
					returnError(&err)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func returnError(err *error) {
	if *err != nil {
		fmt.Fprintf(os.Stderr, "%s%v%s\n", global.ColorRed, *err, global.ColorReset)
		os.Exit(1)
	}
}

func execRun(name string) error {
	c, err := classes.GetConfig()
	if err != nil {
		return err
	}

	if name == "" {
		list.PrintList(c)

		name = readFromStdin()
	}

	command, err := c.FindCommandFromName(name)
	if err != nil {
		return err
	}

	return run.Run(c, command)
}

func execInit() error {
	return initialize.GenerateConfig()
}

func execList() error {
	c, err := classes.GetConfig()
	if err != nil {
		return err
	}
	list.PrintList(c)
	return nil
}

func readFromStdin() string {
	fmt.Printf("\nType in one of the script names listed above:\n%s~>%s ", global.ColorBlue, global.ColorReset)
	r := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// remove the delimeter from the string
	return strings.TrimSuffix(input, "\n")
}
