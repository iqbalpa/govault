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

4.  **Build the binary (optional):**
    You can build the binary to run the CLI directly.
    ```sh
    go build -o govault cmd/main.go
    # or 
    bash build.sh
    ```

5.  **Run the application:**
    ```sh
    go run cmd/main.go
    ```

## Usage (CLI)

GoVault provides a command-line interface (CLI) to manage your secrets.

You can either run the application using `go run` or by building the binary.

### Using `go run`

#### `init`

Initialize the vault with a master password. This is the first command you need to run.

```sh
go run cmd/main.go init --masterPass <your-master-password>
```

#### `add`

Add a new secret to the vault.

```sh
go run cmd/main.go add --masterPass <your-master-password> --name <secret-name> --username <secret-username> --password <secret-password> --note <some-note>
```

#### `list`

List all secrets in the vault.

```sh
go run cmd/main.go list
```

#### `delete`

Delete a secret by its ID. You can get the ID from the `list` command.

```sh
go run cmd/main.go delete --id <secret-id>
```

#### `export`

Export the secrets into JSON file.

```sh
go run cmd/main.go export --filepath <filepath>
```

### Using the built binary

If you have built the binary, you can use it directly:

#### `init`

```sh
./govault init --masterPass <your-master-password>
```

#### `add`

```sh
./govault add --masterPass <your-master-password> --name <secret-name> --username <secret-username> --password <secret-password> --note <some-note>
```

#### `list`

```sh
./govault list
```

#### `delete`

```sh
./govault delete --id <secret-id>
```

#### `export`

```sh
./govault export --filepath <filepath>
```

## Tech Stack

*   [Go](https://golang.org/)
*   [PostgreSQL](https://www.postgresql.org/)
*   [GORM](https://gorm.io/) - The fantastic ORM library for Go
*   [golang-jwt](https://github.com/golang-jwt/jwt) - For handling JSON Web Tokens
*   [go-crypto](https://golang.org/x/crypto) - For encryption and hashing
