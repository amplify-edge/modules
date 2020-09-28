package main

import (
	"github.com/getcouragenow/packages/sys-account/rpc/v2"
	//"github.com/getcouragenow/modules/sys/sys-account/rpc/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth cli",
}

func main() {
	rootCmd.AddCommand(rpc.AuthServiceClientCommand())
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("command failed: %v", err)
	}
}
