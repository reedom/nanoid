package main

import (
	"fmt"
	"os"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nanoid"
	app.Usage = "Generate a new random ID(by nanoid library)"
	app.UsageText = `nanoid [global options]

EXAMPLE:
   nanoid
     -> vtOidCnXVkbZNL2POqgtq
   nanoid --prefix image/ --suffix .png
     -> image/vtOidCnXVkbZNL2POqgtq.png`

	var prefix string
	var suffix string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "prefix, p",
			Usage:       "Add `value` before the ID",
			Destination: &prefix,
		},
		cli.StringFlag{
			Name:        "suffix, s",
			Usage:       "Add `value` after the ID",
			Destination: &suffix,
		},
	}
	app.Action = func(c *cli.Context) {
		id, _ := gonanoid.New()
		fmt.Printf("%v%v%v\n", prefix, id, suffix)
	}

	_ = app.Run(os.Args)
}
