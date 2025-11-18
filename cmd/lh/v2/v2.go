package v2

import (
	apiv2 "github.com/guardian360/go-lighthouse/api/v2"
	"github.com/guardian360/go-lighthouse/client"
	"github.com/guardian360/go-lighthouse/cmd/lh/config"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "v2",
		Short: "Commands for Lighthouse API v2",
		Long:  "Manage resources using the Lighthouse API version 2",
	}
)

func init() {
	RootCmd.AddCommand(probesCmd)
}

func newAPI() *apiv2.API {
	return apiv2.New(client.New(config.Config))
}
