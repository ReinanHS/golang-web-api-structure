package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func NewGreetCommand() *MigrateCommand {
	gc := &MigrateCommand{
		fs: flag.NewFlagSet("migrate", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "name", "up", "Migrate the DB to the most recent version available")

	return gc
}

type MigrateCommand struct {
	fs *flag.FlagSet

	name string
}

func (g *MigrateCommand) Name() string {
	return g.fs.Name()
}

func (g *MigrateCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *MigrateCommand) Run() error {

	if g.fs.Args()[0] == "up" {
		print("Up migration")
	}

	return nil
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a sub-command")
	}

	cmds := []Runner{
		NewGreetCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			err := cmd.Init(os.Args[2:])
			if err != nil {
				return err
			}
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
