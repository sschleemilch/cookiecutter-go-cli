package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"{{ cookiecutter.module_name }}/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Detailed version information",
	Run: func(cmd *cobra.Command, args []string) {
		v := version.GetVersion()
		fmt.Println("Version:\t", v.Number)
		fmt.Println("Build Date:\t", v.BuildDate)
		fmt.Println("Git ref:\t", v.GitRef)
		fmt.Println("sha256:\t\t", v.Sha)
		fmt.Println("OS:\t\t", v.Os)
		fmt.Println("Arch:\t\t", v.Arch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
