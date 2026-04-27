package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func bindFlag(key string, flag string) {
	if err := viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(flag)); err != nil {
		log.Warn().Err(err).Msgf("[Viper] Could not bind flag: %s", flag)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetEnvPrefix("{{ cookiecutter.env_prefix }}")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Config
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "{{ cookiecutter.config_type }} config file")

	// Logging
	rootCmd.PersistentFlags().String("log-level", "info", "Set the log level (debug, info, warn, error, fatal, panic)")
	bindFlag("log.level", "log-level")

	rootCmd.PersistentFlags().Bool("log-caller", false, "Include the caller file and line number")
	bindFlag("log.caller", "log-caller")

	rootCmd.PersistentFlags().Bool("log-json", false, "Log as json messages")
	bindFlag("log.json", "log-json")
}
