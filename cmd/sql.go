package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/chouqimin/Go-tour/internal/sql2struct"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql轉換和處理",
	Long:  "sql轉換和處理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql轉換",
	Long:  "sql轉換",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			Username: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		template := sql2struct.NewStructTemplate()
		templageColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templageColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)

	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "請輸入資料庫的帳號")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "請輸入資料庫的密碼")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "請輸入資料庫的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "請輸入資料庫的編碼")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "請輸入資料庫實例類型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "請輸入資料庫名稱")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "請輸入表名稱")
}
