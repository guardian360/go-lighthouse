package v2

import (
	"github.com/spf13/cobra"
)

var (
	listProbesCmd = &cobra.Command{
		Use:   "list",
		Short: "List all probes",
		Run:   listProbes,
	}

	probesCmd = &cobra.Command{
		Use:   "probes",
		Short: "Manage probes",
		Long:  "Commands to create, update, delete, and list probes",
	}
)

func init() {
	probesCmd.AddCommand(listProbesCmd)
}

func listProbes(cmd *cobra.Command, args []string) {
	resp, err := newAPI().Probes().Get()
	if err != nil {
		cmd.PrintErrln("Error fetching probes:", err)
		return
	}
	for _, probe := range resp.Data {
		cmd.Printf("%+v\n", probe)
	}
}
