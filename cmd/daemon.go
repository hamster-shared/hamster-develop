/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	engine "github.com/hamster-shared/aline-engine"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/controller"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run daemon web app",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("daemon called")

		passwordFlag := cmd.Flags().Lookup("db_password")
		userFlag := cmd.Flags().Lookup("db_user")
		hostFlag := cmd.Flags().Lookup("db_host")
		portFlag := cmd.Flags().Lookup("db_port")
		nameFlag := cmd.Flags().Lookup("db_name")

		DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			userFlag.Value, passwordFlag.Value, hostFlag.Value, portFlag.Value, nameFlag.Value)

		// go Engine.Start()

		port, _ = rootCmd.PersistentFlags().GetInt("port")
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       DSN,   // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
		})
		if err != nil {
			return
		}
		application.SetBean[*gorm.DB]("db", db)
		application.SetBean[engine.Engine]("engine", Engine)
		workflowService := service.NewWorkflowService()
		application.SetBean[*service.WorkflowService]("workflowService", workflowService)
		contractService := service.NewContractService()
		application.SetBean[*service.ContractService]("contractService", contractService)
		reportService := service.NewReportService()
		application.SetBean[*service.ReportService]("reportService", reportService)
		userService := service.NewUserService()
		application.SetBean[*service.UserService]("userService", userService)
		githubService := service.NewGithubService()
		application.SetBean[*service.GithubService]("githubService", githubService)
		loginService := service.NewLoginService()
		application.SetBean[*service.LoginService]("loginService", loginService)
		containerDeployService := service.NewContainerDeployService()
		application.SetBean[*service.ContainerDeployService]("containerDeployService", containerDeployService)
		frontendPackageService := service.NewFrontendPackageService()
		application.SetBean[*service.FrontendPackageService]("frontendPackageService", frontendPackageService)
		templateService.Init(db)
		projectService.Init(db)

		// 定时同步starknet 合约部署状态
		ticker := time.NewTicker(time.Minute * 2)
		go func() {
			for {
				<-ticker.C
				contractService.SyncStarkWareContract()
			}
		}()

		controller.NewHttpService(*handlerServer, port).StartHttpServer()

	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//daemonCmd.PersistentFlags().String("db_password", "123456", "database password")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	daemonCmd.Flags().String("db_password", "123456", "database password")
	daemonCmd.Flags().String("db_user", "root", "database user")
	daemonCmd.Flags().String("db_host", "127.0.0.1", "database host")
	daemonCmd.Flags().String("db_port", "3307", "database port")
	daemonCmd.Flags().String("db_name", "aline", "database name")
}
