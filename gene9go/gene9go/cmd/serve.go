package cmd

import (
	"github.com/lileio/lile"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the RPC server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatal(lile.Serve())
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
