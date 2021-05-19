package cmd

import (
	"fmt"
	"github.com/jsdaniell/recipe-cli/cmd/golang"
	"github.com/jsdaniell/recipe-cli/cmd/react_native"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "recipe-cli",
	Short: "This is a simple cli tool to generate different boilerplate for different type of projects.",
	Long: `The original idea comes from the necessity of create different projects with different newer frameworks that's
coming from the newer technologies.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(`Welcome to the recipe-cli, to use you have to choose the type of your project to generate it:
Actually these are enabled (recipe-cli --project-language):

golang`)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.recipe-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	RootCmd.AddCommand(golang.GoCmd)
	RootCmd.AddCommand(react_native.ReactNativeCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".recipe-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".recipe-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
