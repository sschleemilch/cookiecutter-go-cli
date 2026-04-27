package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"{{ cookiecutter.module_name }}/internal/logger"
	"{{ cookiecutter.module_name }}/version"
)

var cfgFile string

var introLines = []string{
	version.GetVersion().Details(),
}

var rootCmd = &cobra.Command{
	Use:     "{{ cookiecutter.bin_name }}",
	Version: version.GetVersion().Details(),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Init(viper.GetString("log.level"), viper.GetBool("log.caller"), viper.GetBool("log.json"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(strings.Join(introLines, "\n"))
		return nil
	},
}
