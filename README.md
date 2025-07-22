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

## Architecture

GoVault follows a layered architecture that separates concerns and promotes modularity. The main layers are:

-   **`cmd`**: The entry point of the application, responsible for parsing command-line arguments and initializing the CLI.
-   **`cli`**: Contains the core CLI logic, including command definitions and flags. It interacts with the `service` layer to execute user commands.
-   **`service`**: Implements the business logic of the application. It coordinates the interaction between the `repository` and `crypto` layers.
-   **`repository`**: Handles data access and persistence. It interacts with the PostgreSQL database using the GORM library.
-   **`model`**: Defines the data structures used throughout the application, such as `Secret` and `User`.
-   **`crypto`**: Manages all cryptographic operations, including encryption, decryption, and key derivation.
-   **`utils`**: Provides utility functions for tasks like password hashing, configuration management, and database connections.

This layered approach makes the codebase easier to maintain, test, and extend.

## Password Manager Design

The password manager is designed with security as the top priority. Here’s a breakdown of the key design principles:

### Master Password

-   **Single Point of Entry**: The entire vault is protected by a single master password. This password is used to encrypt and decrypt all your secrets.
-   **Hashing**: The master password is not stored directly. Instead, it is hashed using `bcrypt`.

### Encryption

-   **AES-GCM**: All secrets are encrypted using AES-256 in Galois/Counter Mode (GCM). AES is a widely trusted encryption standard, and GCM provides both confidentiality and authenticity.
-   **Key Derivation**: The encryption key is derived from the master password using `scrypt`. This adds an extra layer of security by making it computationally expensive to generate the key.
-   **Unique Salts**: Each secret is encrypted with a unique salt. This ensures that even if two secrets have the same password, their encrypted values will be different.

### Data Storage

-   **Encrypted at Rest**: All secrets are stored in the database in their encrypted form. This means that even if an attacker gains access to the database, they will not be able to read your secrets without the master password.
-   **No Plaintext**: The master password and unencrypted secrets are never stored on disk. They are only held in memory during the execution of a command.

This design ensures that your secrets are protected at all times, both in transit and at rest.
