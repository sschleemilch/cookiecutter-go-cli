package cmd

import (
	"fmt"
	"os"
	"strings"
	"{{ cookiecutter.module_name }}/version"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "{{ cookiecutter.bin_name }}",
	Short: "A brief description of your application",
	Long: `A longer description`,
	Version: version.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initializeLogger(viper.GetString("log.level"), viper.GetBool("log.caller"), viper.GetString("log.file"))
	},
{% if cookiecutter.subcommands != "y" %}
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Running root info")
		log.Debug().Msg("Running root debug")
	},
{% endif %}
}

func initializeLogger(logLevel string, logCaller bool, logFile string) {
		level, err := zerolog.ParseLevel(strings.ToLower(logLevel))
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid log level")
		}
		zerolog.SetGlobalLevel(level)

		consoleWriter := &zerolog.ConsoleWriter{Out: os.Stderr}

		var logFileFd *os.File
		if logFile != "" {
			logFileFd, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatal().Err(err).Msgf("Could not open logfile: %s", logFile)
			}
		}

		if logFileFd != nil {
			log.Logger = zerolog.New(zerolog.MultiLevelWriter(consoleWriter, logFileFd))
		} else {
			log.Logger = zerolog.New(zerolog.MultiLevelWriter(consoleWriter))
		}

		if logCaller {
			log.Logger = log.With().Caller().Logger()
		}

		log.Logger = log.With().Timestamp().Logger()
	}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "{{ cookiecutter.config_type }} config file (default: $HOME/{{ cookiecutter.default_config_prefix }}.{{ cookiecutter.config_type }})")
	rootCmd.PersistentFlags().String("log-level", "info", "Set the log level (debug, info, warn, error, fatal, panic) (default: info)")
	rootCmd.PersistentFlags().String("log-file", "", "Write logs in json format to this file (default: '')")
	rootCmd.PersistentFlags().Bool("log-caller", false, "Include the caller file and line number (default: false)")
	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("log.file", rootCmd.PersistentFlags().Lookup("log-file"))
	viper.BindPFlag("log.caller", rootCmd.PersistentFlags().Lookup("log-caller"))
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
