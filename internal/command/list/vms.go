package list

import (
	"fmt"
	"github.com/cirruslabs/orchard/pkg/client"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

func newListVMsCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "vms",
		Short: "List VMs",
		RunE:  runListVMs,
	}

	return command
}

func runListVMs(cmd *cobra.Command, args []string) error {
	client, err := client.New()
	if err != nil {
		return err
	}

	vms, err := client.VMs().List(cmd.Context())
	if err != nil {
		return err
	}

	if quiet {
		for _, vm := range vms {
			fmt.Println(vm.Name)
		}

		return nil
	}

	table := uitable.New()

	table.AddRow("Name", "Image", "Status", "Restart policy")

	for _, vm := range vms {
		restartPolicyInfo := fmt.Sprintf("%s (%d restarts)", vm.RestartPolicy, vm.RestartCount)

		table.AddRow(vm.Name, vm.Image, vm.Status, restartPolicyInfo)
	}

	fmt.Println(table)

	return nil
}
