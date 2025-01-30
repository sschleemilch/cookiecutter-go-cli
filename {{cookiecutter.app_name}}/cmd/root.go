package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"{{ cookiecutter.module_name }}/logger"
	"{{ cookiecutter.module_name }}/version"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "{{ cookiecutter.bin_name }}",
	Short:   "A brief description of your application",
	Long:    `A longer description`,
	Version: fmt.Sprintf("%s", version.GetVersion()),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Init(viper.GetString("log.level"), viper.GetBool("log.caller"), viper.GetString("log.file"), viper.GetBool("log.json"))
	},
{% if cookiecutter.subcommands != "y" %}
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Running root info")
		log.Debug().Msg("Running root debug")
	},
{% endif %}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Config
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "{{ cookiecutter.config_type }} config file")

	// Logging
	rootCmd.PersistentFlags().String("log-level", "info", "Set the log level (debug, info, warn, error, fatal, panic)")
	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	rootCmd.PersistentFlags().String("log-file", "", "Write logs in json format to this file")
	viper.BindPFlag("log.file", rootCmd.PersistentFlags().Lookup("log-file"))

	rootCmd.PersistentFlags().Bool("log-caller", false, "Include the caller file and line number")
	viper.BindPFlag("log.caller", rootCmd.PersistentFlags().Lookup("log-caller"))

	rootCmd.Flags().Bool("log-json", false, "Log as json messages")
	viper.BindPFlag("log.json", rootCmd.Flags().Lookup("log-json"))
}
