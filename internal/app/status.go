package app

import (
	"fmt"
	"os"
	"strings"
)

func (app *App) showCurrentStatus() {
	fmt.Println("\n--- Current Status ---")
	fmt.Println(strings.Repeat("-", 50))

	// Статус ключей
	fmt.Println("Keys Status:")
	if app.hasKeys {
		fmt.Println("  🗸 Keys are loaded")
		keySize := app.publicKey.N.BitLen()
		fmt.Printf("     Key size: %d bits\n", keySize)
	} else {
		fmt.Println("  𐄂 No keys loaded")
	}

	fmt.Println()

	// Статус файлов
	fmt.Println("Files Status:")
	checkFileContent(INPUT_DATA_FILE, "Input file (plaintext)")
	checkFileContent(OUTPUT_DATA_FILE, "Output file (cipher array)")
	checkFile(PUBLIC_KEY_FILE, "Public key file")
	checkFile(SECRET_KEY_FILE, "Secret key file")

	fmt.Println(strings.Repeat("-", 50))
}

func checkFile(filename, description string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("  𐄂 %s: %s - not found\n", description, filename)
	} else {
		info, _ := os.Stat(filename)
		fmt.Printf("  🗸 %s: %s (%d bytes)\n", description, filename, info.Size())
	}
}

func checkFileContent(filename, description string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("  𐄂 %s: %s - not found\n", description, filename)
	} else {
		info, _ := os.Stat(filename)
		if info.Size() == 0 {
			fmt.Printf("  !  %s: %s (empty)\n", description, filename)
		} else {
			fmt.Printf("  🗸 %s: %s (%d bytes)\n", description, filename, info.Size())
		}
	}
}
