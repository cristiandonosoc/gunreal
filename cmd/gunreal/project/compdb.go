package project

import (
	"fmt"

	"github.com/cristiandonosoc/gunreal/pkg/unreal"
	"github.com/spf13/cobra"
)

var (
	compdbCmd = &cobra.Command{
		Use: "compdb",
		Short: "Generates an usable Compilation Dabatase",
		RunE: executeCompdb,
		SilenceUsage: true,
	}
)

func init() {
	ProjectSectionCmd.AddCommand(compdbCmd)
}

func executeCompdb(cmd *cobra.Command, args []string) error {
	project, err := unreal.NewProject(gGunrealConfig)
	if err != nil {
		return fmt.Errorf("reading project: %w", err)
	}

	if err := project.GenerateCompDB(); err != nil {
		return fmt.Errorf("generating compdb: %w", err)
	}

	return nil
}
