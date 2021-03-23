package workflow

import (
	"context"
	"fmt"
	"log"

	"github.com/raydeann/tink/client"
	"github.com/raydeann/tink/protos/workflow"
	"github.com/spf13/cobra"
)

var (
	fTemplate = "template"
	fHardware = "hardware"
	template  string
	hardware  string
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "create a workflow",
		Example: "tink workflow create [flags]",
		PreRunE: func(c *cobra.Command, args []string) error {
			tmp, _ := c.Flags().GetString(fTemplate)
			err := validateID(tmp)
			return err
		},
		Run: func(c *cobra.Command, args []string) {
			createWorkflow(args)
		},
	}
	flags := cmd.PersistentFlags()
	flags.StringVarP(&template, "template", "t", "", "workflow template")
	flags.StringVarP(&hardware, "hardware", "r", "", "workflow targeted hardwares")

	_ = cmd.MarkPersistentFlagRequired(fHardware)
	_ = cmd.MarkPersistentFlagRequired(fTemplate)
	return cmd
}

func createWorkflow(args []string) {
	req := workflow.CreateRequest{Template: template, Hardware: hardware}
	res, err := client.WorkflowClient.CreateWorkflow(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created Workflow: ", res.Id)
}
