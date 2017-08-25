package cmd

import (
	"github.com/lileio/lile"
	"github.com/mattn/lile-example/gene9go/subscribers"
	"github.com/spf13/cobra"
)

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to and process queue messages",
	Run: func(cmd *cobra.Command, args []string) {
		lile.Subscribe(&subscribers.Gene9goServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(subscribeCmd)
}
