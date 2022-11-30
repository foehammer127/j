package main

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/keg"
	"github.com/rwxrob/help"
	"github.com/rwxrob/grep"
	"github.com/rwxrob/good"
)

var Cmd = &Z.Cmd {
	Name: "j",
	Description: `
		Command Monolith for foehammer.
		`,

	Commands: []*Z.Cmd {
		
		help.Cmd,
		
		// Imported:
		keg.Cmd,
		grep.Cmd,
		good.Cmd,
	},
}


func main() {


	Cmd.Run()
}

