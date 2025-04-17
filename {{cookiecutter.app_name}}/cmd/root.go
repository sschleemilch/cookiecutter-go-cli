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
	Version: version.GetVersion().String(),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Init(viper.GetString("log.level"), viper.GetBool("log.caller"), viper.GetString("log.file"), viper.GetBool("log.json"))
	},
{% if cookiecutter.subcommands != "y" %}
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.GetVersion().Details())
		log.Info().Msg("Running")
	},
{% endif %}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func bindFlag(key string, flag string) {
	if err := viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(flag)); err != nil {
		log.Warn().Err(err).Msgf("[Viper] Could not bind flag: %s", flag)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Config
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "{{ cookiecutter.config_type }} config file")

	// Logging
	rootCmd.PersistentFlags().String("log-level", "info", "Set the log level (debug, info, warn, error, fatal, panic)")
	bindFlag("log.level", "log-level")

	rootCmd.PersistentFlags().String("log-file", "", "Write logs in json format to this file")
	bindFlag("log.file", "log-file")

	rootCmd.PersistentFlags().Bool("log-caller", false, "Include the caller file and line number")
	bindFlag("log.caller", "log-caller")

	rootCmd.PersistentFlags().Bool("log-json", false, "Log as json messages")
	bindFlag("log.json", "log-json")
}
