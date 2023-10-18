package main

import (
	"log"

	"github.com/linkisensei/civitdownloader/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()
}
