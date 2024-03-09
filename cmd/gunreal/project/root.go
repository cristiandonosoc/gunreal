// Package project has the cli commands for the project section.
package project

import (
	"fmt"

	gunreal_config "github.com/cristiandonosoc/gunreal/pkg/config"

	"github.com/spf13/cobra"
)

var (
	gFlags = struct {
		configPath string
	}{}

	gGunrealConfig *gunreal_config.GunrealConfig

	ProjectSectionCmd = &cobra.Command{
		Use:   "project",
		Short: "Commands for dealing with projects",
		Long:  "These are all the commands that will deal with the unreal project.",

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := gunreal_config.LoadConfig(gFlags.configPath)
			if err != nil {
				return fmt.Errorf("loading config file: %w", err)
			}

			gGunrealConfig = config
			return nil
		},
	}
)

func init() {
	ProjectSectionCmd.PersistentFlags().StringVar(&gFlags.configPath, "config-path", "gunreal.yml",
		"Path the config file")
}
