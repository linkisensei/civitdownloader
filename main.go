package main

import (
	"github.com/linkisensei/civitdownloader/app/config"
	"github.com/linkisensei/civitdownloader/cmd"
)

func main() {
	config.Config.Init()
	cmd.Execute()
}
