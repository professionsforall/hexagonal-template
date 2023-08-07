package cmd

import (
	"github.com/professionsforall/hexagonal-template/internal/adapters/http"
	"github.com/spf13/cobra"
)

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "serves application at given port",
	Run: func(cmd *cobra.Command, args []string) {
		http.Init(cmd.Context())
	},
}

func init() {
	rootCommand.AddCommand(serveCommand)
}
