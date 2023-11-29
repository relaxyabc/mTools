package main

import (
	"fmt"
	"github.com/toolbox/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.App{
		Name:                 "toolbox",
		Usage:                "A personal toolbox",
		Version:              "1.0.0",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			cmd.NewPropsCmd(),
			cmd.NewHashCmd(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println("parse param error,process exit. err: ", err)
		return
	}
}
