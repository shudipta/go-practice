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
	"os"

	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"Go-Practice/Book-Server-by-glide/book_server"
	//"github.com/spf13/viper"
)

//var cfgFile string

var port string
var loggedIn bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bookServer",
	Short: "A simple Book Server",
	Long: `A simple Book Server`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		book_server.HandleRequests()
		//defer book_server.ShutdownServer()
		book_server.StartServer(port, loggedIn)
	},
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
	rootCmd.PersistentFlags().StringVar(&port, "port", ":8080", "it is Port no.(default: 8080)")
	rootCmd.PersistentFlags().BoolVar(&loggedIn, "logIn", false, "it is for checking whether the user is logged in or not")
}
