// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"github.com/levigross/grequests"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type UserInfo struct {
	ID string `json:"pk"`
	Name string `json:"username"`
	Email string `json:"email"`
}

type TokenResponse struct {
	Token string `json:"token"`
	User UserInfo `json:"user"`
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cello-client",
	Short: "Cello client is a tool to manage cello api service",
	Long: `This application can manage all cello api service, to control
Networks of hyperledger deployment `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		authUrl := fmt.Sprintf("%s/auth/login/", viper.Get("server.url"))
		fmt.Printf("Pre run url %v\n", authUrl)
		resp, err := grequests.Post(authUrl, &grequests.RequestOptions{JSON: map[string]string{
			"username": viper.GetString("server.username"),
			"password": viper.GetString("server.password"),
		}})
		// You can modify the request by passing an optional RequestOptions struct

		if err != nil {
			log.Fatalln("Unable to make request: ", err)
		}

		if resp.Ok != true {
			log.Printf("Get token failed")
		} else {
			var token TokenResponse
			err := resp.JSON(&token)
			if err != nil {
				panic(err)
			}
			viper.Set("user.token", token.Token)
			viper.Set("user.name", token.User.Name)
			viper.Set("user.email", token.User.Email)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cello.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".cobra-example" (without extension).
		viper.AddConfigPath("$HOME/.cello")
		viper.AddConfigPath("/etc/cello/")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}