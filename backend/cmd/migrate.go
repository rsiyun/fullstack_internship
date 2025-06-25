package cmd

import (
	"api-money-management/internal/models"
	"api-money-management/pkg/database"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "use bun migration tool",
	Long:  `run bun migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.DBConn()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		if err := db.Migrator().DropTable(&models.User{}); err != nil {
			log.Printf("Warning: Failed to drop table: %v", err)
		}
		err = db.AutoMigrate(&models.User{})
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}

		log.Println("Migration completed successfully!")

	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", "5432")
	viper.SetDefault("db_name", "api_money_management")
	viper.SetDefault("db_user", "postgres")
	viper.SetDefault("db_password", "postgres")
}
