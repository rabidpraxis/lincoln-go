package main

import (
  "os"

  "github.com/urfave/cli"
  "github.com/projectdx/lincoln/commands"
)

func main() {
  app := cli.NewApp()
  app.Commands = []cli.Command{
    {
      Name:    "list",
      Aliases: []string{"l"},
      Usage:   "list from yaml file",
      Action:  commands.ListAction,
    },
  }

  app.Run(os.Args)
}
