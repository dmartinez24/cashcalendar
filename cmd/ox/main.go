package main

import (
	"cashcalendar/migrations"
	"context"
	"log"
	"os"

	"github.com/wawandco/ox/cli"
	"github.com/wawandco/ox/plugins/tools/soda"
)

// main function for the tooling cli, will be invoked by Oxpecker
// when found in the source code. In here you can add/remove plugins that
// your app will use as part of its development lifecycle.
func main() {
	cli.Use(soda.NewCommand(migrations.FS()))
	err := cli.Run(context.Background(), os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
