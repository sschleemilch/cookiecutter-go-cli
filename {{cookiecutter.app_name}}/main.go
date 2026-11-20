package main

import (
	"encoding/json"
	"fmt"
	"os"

	"{{ cookiecutter.module_name }}/app"
	"{{ cookiecutter.module_name }}/config"
	"{{ cookiecutter.module_name }}/version"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: {{ cookiecutter.app_name }} <config.json>")
		os.Exit(1)
	}

	cfg, err := readConfig(os.Args[1])
	if err != nil {
		fmt.Printf("could not load config: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println(version.GetVersion().Details())

	app.Run(cfg)
}

func readConfig(path string) (config.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return config.Config{}, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	cfg := config.Defaults()
	if err := json.Unmarshal(b, &cfg); err != nil {
		return config.Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return cfg, nil
}
