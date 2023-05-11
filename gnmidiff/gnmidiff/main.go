package main

import (
	"fmt"
	"os"

	"github.com/lyckety/ygot/gnmidiff/cmd"
)

func main() {
	rootCmd := cmd.RootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
