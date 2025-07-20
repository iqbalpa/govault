package cli

import (
	"fmt"
	"govault/internal/crypto"
	"govault/internal/repository"
	"govault/internal/service"
	"govault/internal/utils"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// Services
var (
	db        *gorm.DB
	vaultRepo *repository.SecretRepository
	vaultSvc  *service.VaultService
	authRepo  *repository.AuthRepository
	authSvc   *service.AuthService
)

// Flags
var (
	masterPass string
	name       string
	username   string
	password   string
	note       string
	id         string
)

// CLI Commands
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

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all available secrets",
		Long:  "List all available secrets (name and username)",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			initServices()
			res, err := vaultSvc.GetAllSecrets()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)
		},
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add new secret",
		Long:  "Add new secret to the govault",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			initServices()
			salt := crypto.GenerateRandomSalt()
			s, err := vaultSvc.CreateSecret(masterPass, name, username, password, note, salt)
			if err != nil {
				fmt.Println("Failed to add new secret")
			}
			fmt.Println("Succesfully added new secret!\n", s)
		},
	}

	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete secret by id",
		Long:  "Delete specified secret",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			initServices()
			s, err := vaultSvc.DeleteSecretById(id)
			if err != nil {
				fmt.Println("Failed to delete secret")
			}
			fmt.Println("Succesfully delete secret!\n", s)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	initCmd.PersistentFlags().StringVarP(&masterPass, "masterPass", "m", "", "the master password to initialize govault")
	addCmd.PersistentFlags().StringVarP(&masterPass, "masterPass", "m", "", "the master password to initialize govault")
	addCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "the secret name")
	addCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "the secret username")
	addCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "the password you want to store")
	addCmd.PersistentFlags().StringVarP(&note, "note", "", "", "additional notes")
	deleteCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "secret id to delete")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
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
