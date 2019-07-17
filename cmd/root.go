// Copyright © 2019 Timothy Jarman
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
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var fromDate string
var toDate string
var format string

// Dates struct
type dates struct {
	list []string
}

func calculateDates(fromDate string, toDate string) dates {
	var dates dates

	f, err := time.Parse("2006-01-02", fromDate)
	if err != nil {
		panic(err)
	}

	t, err := time.Parse("2006-01-02", toDate)
	if err != nil {
		panic(err)
	}

	days := int(t.Sub(f).Hours() / 24)

	for i := 0; i <= days; i++ {
		dates.list = append(dates.list, f.Format("2006-01-02"))
		f = f.AddDate(0, 0, 1)
	}

	return dates
}

// Days returns the number of days between dates
func (r dates) days() int {
	return len(r.list) - 1
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "between-dates",
	Short: "Action a task between two dates",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.between-dates.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&fromDate, "fromDate", "", "From Date e.g. 2019-01-01")
	rootCmd.MarkPersistentFlagRequired("fromDate")

	todayDate := time.Now().Format("2006-01-02")
	rootCmd.PersistentFlags().StringVar(&toDate, "toDate", todayDate, fmt.Sprintf("To Date e.g. %s", todayDate))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".between-dates" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".between-dates")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
