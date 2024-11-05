# Blockchain-Go

🚀 **Blockchain-Go** is a basic blockchain implementation written in Go, utilizing **BadgerDB** for data persistence. This project demonstrates essential blockchain functionality, including block creation, transaction handling, and wallet management, using a CLI interface.

## 📋 Prerequisites

- **Go** version 1.23.1 or higher.
- Basic understanding of blockchain concepts and experience with command-line interfaces is recommended.

## 📂 Project Structure

The project includes:

- `main.go`: Entry point for the CLI, where blockchain commands are run.
- `blockchain/`: Implements the blockchain, transaction, and proof-of-work functionality.
- `cli/cli.go`: Manages command-line commands and arguments.
- `wallet/`: Includes wallet creation, management, and address generation.
- `tmp/blocks/`: Directory for storing blockchain data.

## 🔧 Setup Instructions

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/danubiobwm/blockchain-go.git
   cd blockchain-go
   ```
2. **Install Dependencies**:

- BadgerDB is included in the module files; ensure you have internet access for Go to download dependencies.

3. **Run the Project**:

- Use the CLI to interact with the blockchain.

## 🚀 Commands and Usage

Here’s a guide to using the blockchain's main functions:

#### 1. **Create a New Blockchain**

Create a new blockchain and initialize it with an address to receive mining rewards.

```bash
go run main.go createBlockchain -address "your_address"
```

#### 2. **Print the Blockchain**

Prints all blocks in the current blockchain, showing hashes and transaction details.

```bash
go run main.go printchain
```

#### 3. **Check Balance**

Check the balance of an address.

```bash
go run main.go getBalance -address "your_address"
```

#### 4. **Send Tokens**

Send tokens from one address to another. Make sure there are enough funds in the sender's address.

```bash
go run main.go send -from "sender_address" -to "recipient_address" -amount amount
```

If successful, the transaction ID will be printed.

#### 5. **Create a New Wallet**

Generate a new wallet with a unique address.

```bash
go run main.go createWallet
```

#### 6. **List All Addresses**

List all addresses created in the wallet file.

```bash
go run main.go listAddresses
```

## 🌟 Example Workflow

**_Create a Blockchain:_**

```bash
go run main.go createBlockchain -address "jonh"
```

**_Check Balance:_**

```bash
go run main.go getBalance -address "jonh"

```

**_Send Tokens:_**

```bash
go run main.go send -from "jonh" -to "danubio" -amount 10
```

**_Confirm Balance Update:_**

```bash
go run main.go getBalance -address "danubio"
```

Each command above will interact with the blockchain, update balances, and save transactions in BadgerDB.

## 📄 Logging

BadgerDB logs can help trace operations, especially during database access or debugging. Logs include information like compaction priorities, replay actions, and other database-specific messages.

### 🧑‍💻 Developed by Danubio

#### Developed with ❤️ and Go! 🐹
