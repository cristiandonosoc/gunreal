// Binary unreal_lister is a little test application to test out the unreal project querying code.
package main

import (
	"fmt"

	"github.com/cristiandonosoc/gunreal/cmd/gunreal/project"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gunreal",
		Short: "Tool for dealing with Unreal projects",
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("DEV")
		},
	}
)

func init() {
	// TODO(cdc): Evaluate using Viper for configs.
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(project.ProjectSectionCmd)
}

func main() {
	rootCmd.Execute()
}
