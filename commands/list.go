package commands

import (
  "fmt"
  "log"
  "io/ioutil"
  "github.com/urfave/cli"
  "gopkg.in/yaml.v2"
)

type Resource struct {
  Name string
}

type App struct {
  Name string `yaml:"name"`
  Url  string `yaml:"url"`
}

type Config struct {
  DefaultBranch string `yaml:"default-branch"`
  Resources []Resource `yaml:"resources"`
  Apps      []App      `yaml:"apps"`
}

func ListAction(c *cli.Context) error {
  t := Config{}

  data, err := ioutil.ReadFile(c.Args().Get(0))
  if err != nil {
    log.Fatalf("error: %v", err)
  }

  err = yaml.Unmarshal([]byte(data), &t)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  fmt.Printf("--- t:\n%v\n\n", t)

  fmt.Print(t.Resources[0].Name)

  return nil
}
