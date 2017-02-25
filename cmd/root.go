// Copyright © 2017 James Eckersall <james.eckersall@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
	sslCert string
	sslKey  string
	sslCa   string
	debug   bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "goh",
	Short: "Query and execute commands against hosts matched from puppetdb",
	Long:  "Query and execute commands against hosts matched from puppetdb",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCm.d.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// global flags
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$HOME/.goh.yaml", "config file.")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging.")
	RootCmd.PersistentFlags().StringVar(&sslCert, "ssl_cert", "", "Client SSL certificate file to connect to puppetdb.")
	RootCmd.PersistentFlags().StringVar(&sslKey, "ssl_key", "", "Client SSL key file to connect to puppetdb.")
	RootCmd.PersistentFlags().StringVar(&sslCa, "ssl_ca", "", "SSL CA file to connect to puppetdb.")

	viper.SetDefault("debug", false)
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("ssl_cert", RootCmd.PersistentFlags().Lookup("ssl_cert"))
	viper.BindPFlag("ssl_key", RootCmd.PersistentFlags().Lookup("ssl_key"))
	viper.BindPFlag("ssl_ca", RootCmd.PersistentFlags().Lookup("ssl_ca"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".goh")  // name of config file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
