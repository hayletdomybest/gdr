package main

import (
	"fmt"
	"os"

	"github.com/hayletdomybest/gdr/cmd"
)

func main() {
	root := cmd.NewRootCmd()
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
