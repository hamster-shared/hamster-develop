/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"time"

	engine "github.com/hamster-shared/aline-engine"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/aline-engine/model"
	"github.com/hamster-shared/aline-engine/pipeline"
	"github.com/hamster-shared/hamster-develop/pkg/controller"
	service2 "github.com/hamster-shared/hamster-develop/pkg/service"

	"github.com/spf13/cobra"
)

func getEngine() engine.Engine {
	e, err := engine.NewMasterEngine(50001)
	if err != nil {
		logger.Error("启动引擎失败 err:", err)
		panic(err)
	}
	return e
}

// rootCmd represents the base command when called without any subcommands
var (
	port            = 8080
	pipelineFile    string
	templateService = service2.NewTemplateService()
	projectService  = service2.NewProjectService()
	Engine          = getEngine()
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

			// 启动 executor
			yamlByte, err := io.ReadAll(cicdFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			yaml := string(yamlByte)

			job, _ := pipeline.GetJobFromYaml(yaml)

			err = Engine.CreateJob(job.Name, yaml)

			jobDetail, err := Engine.ExecuteJob(job.Name, getRandomNumber())
			if err != nil {
				logger.Error("err:", err)
				return
			}

			for jobDetail.Status <= model.STATUS_RUNNING {
				time.Sleep(time.Second * 3)
				jobDetail, err = Engine.GetJobHistory(jobDetail.Name, jobDetail.Id)
				if err != nil {
					logger.Error("err:", err)
				}
			}
		},
	}
)

// 获取一个随机的数，大于 1000，小于 65535
func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(64535) + 1000
}

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
