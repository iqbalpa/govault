# GoVault

GoVault is a simple and secure password manager built with Go. It allows you to store and manage your secrets, with a master password to encrypt and protect your data. All secrets are encrypted using AES.

## Features

*   **Vault Initialization**: Securely initialize the vault with a master password.
*   **Login**: Authenticate with your master password to access your secrets.
*   **Secret Management**:
    *   Create new secrets (e.g., username, password, notes).
    *   Retrieve a specific secret.
    *   List all stored secrets.
    *   Delete secrets you no longer need.
*   **Encryption**: Secrets are encrypted using AES to ensure your data is secure.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

*   [Go](https://golang.org/doc/install) (version 1.24 or later)
*   [PostgreSQL](https://www.postgresql.org/download/)

### Installation

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/your-username/govault.git
    cd govault
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Set up the database:**
    *   Create a PostgreSQL database.
    *   Create a `.env` file in the root of the project by copying `.env.example`

4.  **Run the application:**
    ```sh
    go run cmd/main.go
    ```

## Usage (CLI - Coming Soon)

The following commands will be available in the CLI version of GoVault:

*   `govault init`: Initialize the vault with a master password.
*   `govault login`: Log in to the vault.
*   `govault create`: Create a new secret.
*   `govault get <secret-name>`: Retrieve a secret.
*   `govault list`: List all secrets.
*   `govault delete <secret-name>`: Delete a secret.

## Tech Stack

*   [Go](https://golang.org/)
*   [PostgreSQL](https://www.postgresql.org/)
*   [GORM](https://gorm.io/) - The fantastic ORM library for Go
*   [golang-jwt](https://github.com/golang-jwt/jwt) - For handling JSON Web Tokens
*   [go-crypto](https://golang.org/x/crypto) - For encryption and hashing
