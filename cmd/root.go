package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fileLocation string
var configPath string

var rootCmd = &cobra.Command{
	Use:   "note",
	Short: "note - CLI for the 5 Ls",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
}

// I can't imagine a situation in which this fails - non login shells?
var home, _ = homedir.Dir()

//Execute is the entrypoint of cmd calls
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		reportError(err.Error(), true)
	}
}

func init() {
	viper.SetDefault("noteslocation", "")
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("toml")
	//https://0x46.net/thoughts/2019/02/01/dotfile-madness/
	configPath = os.Getenv("XDG_CONFIG_HOME")
	if configPath == "" {
		configPath = home
	}
	configPath = filepath.Join(configPath, ".config", "note")
	viper.AddConfigPath(configPath)
	viper.SetConfigName(".note")

	err := tryReadConfig()
	if err != nil {
		os.Exit(1)
	}

	viper.AutomaticEnv() // read in environment variables that match
}

func tryReadConfig() (err error) {
	if err := viper.ReadInConfig(); err == nil {
		fileLocation = viper.GetString("noteslocation")
	} else {
		if _, err := os.Stat(configPath + string(os.PathSeparator) + ".punch.toml"); err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(configPath, os.ModePerm)
				if err != nil {
					reportError("Couldn't generate default config file", false)
					return err
				}
				err = viper.WriteConfigAs(configPath + string(os.PathSeparator) + ".punch.toml")
				if err != nil {
					reportError("Couldn't generate default config file", false)
					return err
				}
			}
		} else {
			reportError("You have an issue in your current config", false)
			return errors.New("configuration error")
		}

		fmt.Println("Generated default config.")
		_ = tryReadConfig()
	}
	return nil
}
