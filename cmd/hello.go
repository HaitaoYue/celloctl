package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//var foo string
//var toggle bool

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "says hello",
	Long:  `This subcommand says hello`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Server url is %v\n", viper.Get("server.url"))
		fmt.Printf("user token is %s\n", viper.GetString("user.token"))
		//resp, err := grequests.Get("http://httpbin.org/get", nil)
		//// You can modify the request by passing an optional RequestOptions struct
		//
		//if err != nil {
		//	log.Fatalln("Unable to make request: ", err)
		//}
		//
		//fmt.Println(resp.String())
	},
}

func init() {
	RootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//helloCmd.Flags().StringVar(&foo,"foo", "foo", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//helloCmd.Flags().BoolP(&toggle,"toggle", "t", false, "Help message for toggle")
}
