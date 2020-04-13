package main

import (
	"github.com/hedan806/go-ambari-client/internal/cmd/generate/commands"

	//_ "github.com/hedan806/go-ambari-client/internal/cmd/generate/commands/genexamples"
	_ "github.com/hedan806/go-ambari-client/internal/cmd/generate/commands/gensource"
	_ "github.com/hedan806/go-ambari-client/internal/cmd/generate/commands/genstruct"
	//_ "github.com/hedan806/go-ambari-client/internal/cmd/generate/commands/gentests"
)

func main() {
	commands.Execute()
}
