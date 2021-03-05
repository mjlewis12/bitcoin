/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type Rate struct {
	Code string  `json:"code"`
	Rate float64 `json:"rate_float"`
}

type Rates struct {
	USD Rate `json:"USD"`
}

type Timestamp struct {
	Updated time.Time `json:"updatedISO"`
}

type Response struct {
	Time      Timestamp `json:"time"`
	Chartname string    `json:"chartName"`
	Rates     Rates     `json:"bpi"`
}

//t, _ := time.Parse(time.RFC3339,res.Time.Updated)
//fmt.Println(t)
//.Rates.USD.Rate
// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Displays bitcoin's current price.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, _ := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
		defer resp.Body.Close()

		b, _ := ioutil.ReadAll(resp.Body)
		var res Response
		json.Unmarshal(b, &res)

		fmt.Println(res.Time.Updated.Format(time.RFC1123), res.Rates.USD.Rate)
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
