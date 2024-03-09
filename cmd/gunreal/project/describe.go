package project

import (
	"context"
	"fmt"
	"time"

	"github.com/cristiandonosoc/gunreal/pkg/unreal"
	"github.com/spf13/cobra"
)

var (
	describeCmd = &cobra.Command{
		Use:   "describe",
		Short: "describe an Unreal project managed by Gunreal",
		RunE:  executeDescribe,
	}
)

func init() {
	ProjectSectionCmd.AddCommand(describeCmd)

}

func executeDescribe(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()

	project, err := unreal.NewProject(gGunrealConfig)
	if err != nil {
		return fmt.Errorf("reading project: %w", err)
	}

	if err := project.IndexModules(ctx); err != nil {
		return fmt.Errorf("indexing unreal project: %w", err)
	}

	duration := time.Since(start)

	fmt.Printf("Indexing took %v to execute.\n\n", duration)

	description, err := project.Describe()
	if err != nil {
		return fmt.Errorf("describing project: %w", err)
	}

	fmt.Println(description)

	return nil

}
