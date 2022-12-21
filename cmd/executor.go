/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hamster-shared/a-line/pkg/controller"
	"github.com/spf13/cobra"
)

var (
	registry string
)

// executorCmd represents the executor command
var executorCmd = &cobra.Command{
	Use:   "executor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		go Engine.Start()

		port, _ = rootCmd.PersistentFlags().GetInt("port")
		controller.NewHttpService(*handlerServer, port).StartHttpServer()
	},
}

func init() {
	rootCmd.AddCommand(executorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	executorCmd.Flags().StringVarP(&registry, "registry", "r", "", "registry url, example: 192.168.0.4:8080")
}
