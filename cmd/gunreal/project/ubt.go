package project

import (
	"fmt"

	"github.com/cristiandonosoc/gunreal/pkg/unreal"
	"github.com/spf13/cobra"
)

var (
	ubtCmd = &cobra.Command{
		Use:   "ubt",
		Short: "Runs UBT",
		RunE:  executeUBT,
		SilenceUsage: true,
	}
)

func init() {
	ProjectSectionCmd.AddCommand(ubtCmd)
}

func executeUBT(cmd *cobra.Command, args []string) error {
	project, err := unreal.NewProject(gGunrealConfig)
	if err != nil {
		return fmt.Errorf("reading project: %w", err)
	}

	if err := project.UBT(args); err != nil {
		return fmt.Errorf("running UBT: %w", err)
	}

	return nil
}
