package cmd

import (
	"github.com/toolbox/props"
	"github.com/urfave/cli/v2"
	"strings"
)

// NewPropsCmd propertiesCommand
func NewPropsCmd() *cli.Command {
	return &cli.Command{
		Name:    "properties",
		Aliases: []string{"props", "p"},
		Usage:   "toolbox props",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "content",
				Aliases: []string{"c"},
				Usage:   "toolbox props -c \"xxx\" ",
			},
			&cli.IntFlag{
				Name:    "mode",
				Aliases: []string{"m"},
				Value:   0,
				Usage:   "0: auto;1: toUnicode;2: toStr",
			},
		},
		Action: func(context *cli.Context) error {
			value := context.Value("content")
			mode := context.Value("mode")
			originStr := value.(string)
			switch mode {
			case 0:
				if strings.Contains(originStr, "\\u") {
					return props.Convert2Str(originStr)
				} else {
					return props.Convert2Unicode(originStr)
				}
			case 1:
				return props.Convert2Unicode(originStr)
			case 2:
				return props.Convert2Str(originStr)
			default:
				return nil
			}

		},
	}
}
