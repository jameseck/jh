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
	"github.com/jameseck/jh/puppetdb"
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
		fmt.Println("query called")
		//fmt.Printf("order: %s\n", viper.Get("query.order"))
		//fmt.Printf("puppetdb-server: %s\n", viper.Get("puppetdb-server"))
		//fmt.Printf("ssl-cert: %s\n", viper.Get("ssl-cert"))
		//fmt.Printf("ssl-key: %s\n", viper.Get("ssl-key"))
		//fmt.Printf("ssl-ca: %s\n", viper.Get("ssl-ca"))

		fmt.Printf("all: %#v\n", viper.AllSettings())

		fmt.Printf("cfg\n")
		fmt.Printf("%#v\n", Conf)
		err := viper.Unmarshal(&Conf)
		if err != nil {
			fmt.Printf("unable to decode into struct, %v\n", err)
		}
		fmt.Printf("cfg\n")
		fmt.Printf("%#v\n", Conf)

		f := puppetdb.FactFilters{
			Filters: []puppetdb.FactFilter{
				puppetdb.FactFilter{
					Name:     "osfamily",
					Operator: "=",
					Value:    "RedHat",
				},
				puppetdb.FactFilter{
					Name:     "kernel",
					Operator: "=",
					Value:    "Linux",
				},
			},
		}

		for i := range Conf.Servers {
			s := Conf.Servers[i]

			conn = puppetdb.New(s.Cert, s.Key, s.Ca, s.Fqdn)
			fmt.Printf(conn.Get(args[0], f, "and"))
		}

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

	viper.BindPFlags(queryCmd.PersistentFlags())

}
