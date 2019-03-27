package cmd

import (
	"fmt"
	"github.com/celloctl/internal"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resource in cell service",
	Long:  `Get resource in cello service,
			supported kind: Agent`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please input type to get, agent")
			os.Exit(1)
		}
		types := strings.Split(args[0], ",")
		for _, resourceType := range types {
			switch resourceType {
			case "agent":
				token := fmt.Sprintf("JWT %s", viper.GetString("auth.token"))
				baseUrl := viper.GetString("server.url")
				agentListResponse, err := internal.ListAgents(baseUrl, token)
				if err != nil {
					panic(err)
				} else {
					t := table.NewWriter()
					t.SetOutputMirror(os.Stdout)
					t.AppendHeader(table.Row{"#", "Name", "Worker API", "Capacity", "Node Capacity", "Status", "Create Time", "Organization"})
					for _, value := range agentListResponse.Agents {
						t.AppendRow([]interface{}{value.ID, value.Name, value.WorkerAPI, value.Capacity, value.NodeCapacity, value.Status, value.CreateTime, value.OrgID})
					}
					t.AppendFooter(table.Row{"", "", "", "", "", "", "Total", agentListResponse.Total})
					t.Render()
				}
				break
			default:
				break
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
