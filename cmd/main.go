package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Goalt/hostmonitor/cmd/subcomands"
	"github.com/Goalt/hostmonitor/cmd/variables"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "hostmonitor",
		Usage:    "./hostmonitor",
		Commands: subcomands.Get(),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    variables.DebugLevel,
				Value:   1,
				EnvVars: []string{variables.DebugLevel},
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println(os.Environ())

			for now := range time.Tick(time.Second * 4) {
				fmt.Println(now)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
