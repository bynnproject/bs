package cmd

import (
	"github.com/urfave/cli/v2"
	"bs/cmd/store"
)

func NewCmdApp()  *cli.App{

	return &cli.App{
		Commands: []*cli.Command{
				store.NewStoreCmd(),
		},
	}

}
