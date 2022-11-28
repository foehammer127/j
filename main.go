package main

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/keg"
	"github.com/rwxrob/help"
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
	},
}


func main() {


	Cmd.Run()
}

