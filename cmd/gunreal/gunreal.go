// Binary unreal_lister is a little test application to test out the unreal project querying code.
package main

import (
	"fmt"
	"os"

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

func internalMain() error {
	return rootCmd.Execute()

}

func main() {
	if err := internalMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
