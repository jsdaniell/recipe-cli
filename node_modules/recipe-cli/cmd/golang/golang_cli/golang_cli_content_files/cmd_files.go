package golang_cli_content_files

import (
	"log"
	"os"
)

func CreateCmdPackage(username, projectName string) {
	err := os.Mkdir(projectName + "/cmd", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeApiServerFile(username, projectName)
}

func writeApiServerFile(username, projectName string){
	var content = `package cmd

import (
	"fmt"
	"github.com/` + username + `/` + projectName + `/cmd/other_command"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "` + projectName + `",
	Short: "This is a simple cli tool to generate different boilerplate for different type of projects.",
	Long: ` + "`" + `The original idea comes from the necessity of create different projects with different newer frameworks that's
	coming from the newer technologies.` + "`" + `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(` + "`" + `Welcome to your CLI application!` + "`" + `)
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

	RootCmd.AddCommand(golang.OtherSubCommand)
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
`

	file, err := os.Create(projectName + "/cmd/root.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}