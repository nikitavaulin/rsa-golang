package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nikitavaulin/rsa-golang/internal/keys"
)

const (
	INPUT_DATA_FILE  string = "input.txt"
	OUTPUT_DATA_FILE string = "output.txt"
	PUBLIC_KEY_FILE  string = "key_public.txt"
	SECRET_KEY_FILE  string = "key_secret.txt"
)

type App struct {
	publicKey keys.PublicKey
	secretKey keys.SecretKey
	hasKeys   bool
	scanner   *bufio.Scanner
}

func NewApp() *App {
	return &App{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (app *App) Run() {
	fmt.Println("===== RSA Encryption/Decryption Tool =====")

	for {
		app.showMenu()
		choice := app.getUserInput("Select option: ")

		switch choice {
		case "1":
			app.generateNewKeys()
		case "2":
			app.readKeyFiles()
		case "3":
			app.encryptData()
		case "4":
			app.decryptData()
		case "5":
			app.saveKeysToFiles()
		case "6":
			app.showCurrentStatus()
		case "0", "q", "Q", "exit", "quit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}

		fmt.Println("\nPress Enter to continue...")
		app.scanner.Scan()
	}
}

func (app *App) showMenu() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("MAIN MENU")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("1. Generate new key pair")
	fmt.Println("2. Read keys from files")
	fmt.Println("3. Encrypt data (reads input.txt as plain text)")
	fmt.Println("4. Decrypt data (reads input.txt as cipher array)")
	fmt.Println("5. Save current keys to files")
	fmt.Println("6. Show current status")
	fmt.Println("0. Exit")
	fmt.Println(strings.Repeat("-", 50))
}

func (app *App) getUserInput(prompt string) string {
	fmt.Print(prompt)
	app.scanner.Scan()
	return strings.TrimSpace(app.scanner.Text())
}
