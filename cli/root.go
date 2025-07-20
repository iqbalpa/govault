package cli

import (
	"fmt"
	"govault/internal/repository"
	"govault/internal/service"
	"govault/internal/utils"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	vaultRepo *repository.SecretRepository
	vaultSvc  *service.VaultService
	authRepo  *repository.AuthRepository
	authSvc   *service.AuthService
)

var (
	masterPass string
)

var (
	rootCmd = &cobra.Command{
		Use:   "govault",
		Short: "A secure local password manager",
		Long:  "Govault is a secure local password manager using cryptographic encryption",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Govault -- Local Password Manager  v0.1.0")
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Govault",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Govault v0.1.0")
		},
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Init master password for govault",
		Long:  "Master password initialization to use govault",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			initServices()
			authSvc.InitMasterPass(masterPass)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	initCmd.PersistentFlags().StringVarP(&masterPass, "masterPass", "m", "", "the master password to initialize govault")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
}

func initServices() {
	// Initialize database
	db := utils.ConnectDb()
	utils.MigrateDb(db)

	// Initialize repositories
	vaultRepo = repository.New(db)
	authRepo = repository.NewAuthRepo(db)

	// Initialize services
	vaultSvc = service.New(*vaultRepo)
	authSvc = service.NewAuthService(*authRepo)
}
