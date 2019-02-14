package main

import (
	"os"

	"github.com/Fize/logcatbeat/cmd"

	_ "github.com/Fize/logcatbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
