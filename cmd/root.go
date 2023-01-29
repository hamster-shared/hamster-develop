/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	engine "github.com/hamster-shared/aline-engine"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/aline-engine/model"
	"github.com/hamster-shared/aline-engine/pipeline"
	"github.com/hamster-shared/hamster-develop/pkg/controller"
	service2 "github.com/hamster-shared/hamster-develop/pkg/service"
	"io"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	port            = 8080
	pipelineFile    string
	templateService = service2.NewTemplateService()
	projectService  = service2.NewProjectService()
	Engine          = engine.NewEngine()
	handlerServer   = controller.NewHandlerServer(Engine, templateService, projectService)
	rootCmd         = &cobra.Command{
		Use:   "aline",
		Short: "aline is ci tool that can build and deploy",
		Long: `Aline is the core execution engine of hamster. 
It has a separate command line entry and can execute the 
hamster pipeline file in the local environment.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			//wd, _ := os.Getwd()
			cicdFile, err := os.Open(path.Join(pipelineFile))
			defer cicdFile.Close()
			if err != nil {
				fmt.Println("file error")
				return
			}

			// 启动executor
			yamlByte, err := io.ReadAll(cicdFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			yaml := string(yamlByte)

			job, _ := pipeline.GetJobFromYaml(yaml)

			go Engine.Start()

			err = Engine.CreateJob(job.Name, yaml)

			jobDetail, err := Engine.ExecuteJob(job.Name)
			if err != nil {
				logger.Error("err:", err)
			}

			for jobDetail.Status <= model.STATUS_RUNNING {
				time.Sleep(time.Second * 3)
				jobDetail = Engine.GetJobHistory(jobDetail.Name, jobDetail.Id)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.a-line-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVar(&pipelineFile, "file", "cicd.yml", "pipeline file")

	rootCmd.PersistentFlags().IntP("port", "p", 8080, "http port")
}
