package cli

import (
	"fmt"
	"govault/internal/crypto"
	"govault/internal/repository"
	"govault/internal/service"
	"govault/internal/utils"
	"os"

	"github.com/k0kubun/pp/v3"
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
	name     string
	username string
	password string
	note     string
	id       string
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
			masterPass := promptPassword()
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
			masterPass := promptPassword()
			if _, err := authSvc.Login(masterPass); err != nil {
				fmt.Println("master password is incorrect")
				os.Exit(1)
			}
			res, err := vaultSvc.GetAllSecrets()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			utils.PPrint.Println(res)
		},
	}

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get secret by id",
		Long:  "Get specified secret",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			initServices()
			masterPass := promptPassword()
			if _, err := authSvc.Login(masterPass); err != nil {
				fmt.Println("master password is incorrect")
				os.Exit(1)
			}
			s, err := vaultSvc.GetSecretById(masterPass, id)
			if err != nil {
				fmt.Println("Failed to get secret")
				os.Exit(1)
			}
			fmt.Println("Succesfully retrieve secret!")
			pp.Println(s)
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
			masterPass := promptPassword()
			if _, err := authSvc.Login(masterPass); err != nil {
				fmt.Println("master password is incorrect")
				os.Exit(1)
			}
			s, err := vaultSvc.CreateSecret(masterPass, name, username, password, note, salt)
			if err != nil {
				fmt.Println("Failed to add new secret")
				os.Exit(1)
			}
			fmt.Println("Succesfully added new secret!")
			utils.PPrint.Println(s)
		},
	}

	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete secret by id",
		Long:  "Delete specified secret",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			initServices()
			masterPass := promptPassword()
			if _, err := authSvc.Login(masterPass); err != nil {
				fmt.Println("master password is incorrect")
				os.Exit(1)
			}
			s, err := vaultSvc.DeleteSecretById(id)
			if err != nil {
				fmt.Println("Failed to delete secret")
				os.Exit(1)
			}
			fmt.Println("Succesfully delete secret!")
			utils.PPrint.Println(s)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	addCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "the secret name")
	addCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "the secret username")
	addCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "the password you want to store")
	addCmd.PersistentFlags().StringVarP(&note, "note", "", "", "additional notes")
	deleteCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "secret id to delete")
	getCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "secret id to delete")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(getCmd)
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

func promptPassword() string {
	var masterPass string
	fmt.Print("Enter the master password: ")
	fmt.Scan(&masterPass)
	return masterPass
}
