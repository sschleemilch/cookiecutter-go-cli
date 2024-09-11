package cmd

import (
	"{{ cookiecutter.module_name }}/version"
	"github.com/spf13/cobra"
	"fmt"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Detailed version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("OS / Arch:", version.OsArch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
