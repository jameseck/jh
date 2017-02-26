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
	"github.com/jameseck/goh/puppetdb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	order        string
	includefacts bool
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query for nodes from puppetdb",
	Long:  "Query for nodes from puppetdb",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("query called")
		fmt.Printf("order: %s\n", viper.Get("query.order"))
		fmt.Printf("puppetdb-server: %s\n", viper.Get("puppetdb-server"))
		fmt.Printf("ssl-cert: %s\n", viper.Get("ssl-cert"))
		fmt.Printf("ssl-key: %s\n", viper.Get("ssl-key"))
		fmt.Printf("ssl-ca: %s\n", viper.Get("ssl-ca"))
		p := puppetdb.New(viper.GetString("ssl-cert"), viper.GetString("ssl-key"), viper.GetString("ssl-ca"), viper.GetString("puppetdb-server"))
		p.Get("bob")

		fmt.Printf("%v", factfilter)
		//	fmt.Println(ret)
	},
}

func init() {
	RootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")
	queryCmd.PersistentFlags().StringVarP(&order, "order", "o", "fqdn", "Sort order for node list (fqdn,fact).")
	queryCmd.PersistentFlags().BoolVarP(&includefacts, "include-facts", "i", true, "Include facts in output that were used in criteria.")

	viper.BindPFlag("query.order", queryCmd.PersistentFlags().Lookup("order"))
	viper.BindPFlag("includefacts", queryCmd.PersistentFlags().Lookup("includefacts"))

}
