package store

import (
	"github.com/urfave/cli/v2"
	store_server "bs/services/store/server"
	store_client "bs/clients/store"
	"log"
)

func NewStoreCmd()  *cli.Command{

	return &cli.Command{
		Name:                   "store",
		Usage:                  "store for project",
		Subcommands: []*cli.Command{
			NewStoreServerSubcommand(),
			NewStoreUploadSubcommand(),
			NewStorDownloadSubcommand(),
		},
	}

}

func NewStoreServerSubcommand() *cli.Command {
	return &cli.Command{
		Name:                   "server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "port",
				Usage: "port",
				Required:true,
			},
			&cli.StringFlag{
				Name: "path",
				Usage: "path",
				Required:true,
			},
		},

		Action: func(context *cli.Context) error {
			port := context.String("port")
			path := context.String("path")
			err := store_server.NewUploadFileApp(port , path).Run()
			return err
		},
	}
}

func NewStoreUploadSubcommand() *cli.Command {
	return &cli.Command{
		Name : "upload",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Required:true,
				Name: "address",
				Usage: "language for the greeting",
			},

			&cli.StringFlag{
				Required:true,
				Name: "path",
				Usage: "path of file to upload ",
			},
			&cli.StringFlag{
				Required:true,
				Name: "name",
				Usage: "fileName for new project",
			},
			&cli.StringFlag{
				Required:true,
				Name: "tag",
				Usage: "tag for project in store",
			},
		},
		Action: func(c *cli.Context) error {
			store_client.NewUploadFileClient(c.String("address")).UploadFile(c.String("path") , c.String("name"))
			return nil
		},
	}
}

func NewStorDownloadSubcommand() *cli.Command  {
	return &cli.Command{
		Name : "download",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Required:true,
				Name :"address",
			},
			&cli.StringFlag{
				Required:true,
				Name:"source",

			},
			&cli.StringFlag{
				Required:true,
				Name:"target",
			},
			&cli.StringFlag{
				Required:true,
				Name:"tag",
			},
		},
		Action: func(c *cli.Context) error {
			err := store_client.NewUploadFileClient(c.String("address")).DownloadFile(c.String("source") , c.String("target") , c.String("tag"))
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		},
	}
}

