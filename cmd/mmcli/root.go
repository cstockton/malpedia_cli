package main

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	apikeyName = "apikey"
	configName = "config"
)

var (
	cfgFile    string
	apiKey     string
	log        = logrus.New()
	jsonFormat bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "malpedia_cli",
	Short: "a commandline application used to interact with the malpedia API",
	Long: `Malpedia_cli is a interface to malpedias REST API, making interaction with the service more
streamlines and intuitive. It allows users to scan samples against Malpedia's yara rules, lookup malware
families, lookup actors, download yara rules, download samples and pivot from IOCs. The service itself
requires an API key which can be acquired once the user had an account. Malpedia_cli can take the API
key via a commandline argument, a config file or a config file in the home directory.

Usage examples of the application are as follows:
	- malpedia_cli actor <actor name>
	- malpedia_cli actors
	- malpedia_cli downloadFamily <family name>
	- malpedia_cli downloadSample <hash1>
	- malpedia_cli downloadYara <family>
	- malpedia_cli downloadYara <TLP Level>
	- malpedia_cli families
	- malpedia_cli familySamples <family>
	- malpedia_cli family <family>
	- malpedia_cli scanBinary <file1>
	- malpedia_cli scanYara <file1>
	- malpedia_cli scanYaraAgainstFamily <file1>
	- malpedia_cli version
`,
}

func init() {
	// I have never used cobra, but this pattern feels a little off. Seems like
	// the config should be parsed once before cmd.Execute() in main.go.
	cobra.OnInitialize(func() {
		if err := initConfig(); err != nil {
			log.Fatal(err)
		}
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.malpedia_cli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&jsonFormat, "json", "j", false, "will return raw json data to the user")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() error {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return err
		}

		// Search config in home directory with name ".malpedia_cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".malpedia_cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil && apiKey == "" && !jsonFormat {
		return errors.Wrapf(err, "config error (file: %v)", viper.ConfigFileUsed())
	}

	vConfig := viper.GetViper()

	if vConfig.IsSet(apikeyName) {
		apiKey = vConfig.GetString(apikeyName)
	} else if apiKey == "" {
		return errors.New("no apikey available")
	}
}
