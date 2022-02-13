package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Goalt/hostmonitor/cmd/subcomands"
	"github.com/Goalt/hostmonitor/cmd/variables"
	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/provider"
	"github.com/sethvargo/go-signalcontext"
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
			cfg := config.Config{
				Logger: config.Logger{
					SetReportCaller: true,
					Level:           ctx.Int(variables.DebugLevel),
				},
				GRPCServer: config.GRPCServer{
					Port: 20000,
				},
				ProxyServer: config.ProxyServer{
					Port: 8080,
				},
				Updater: config.Updater{
					Interval: 3 * time.Second,
				},
			}
			fmt.Printf("%+v\n", cfg)

			signalCtx, cancel := signalcontext.OnInterrupt()
			defer cancel()
			
			app, cleanup, err := provider.InitializeApp(cfg, signalCtx)
			if cleanup != nil {
				defer cleanup()
			}
			if err != nil {
				fmt.Println(err)
				return nil
			}

			err = app.Run()
			if err != nil {
				fmt.Println(err)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
