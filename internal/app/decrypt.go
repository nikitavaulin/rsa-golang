package app

import (
	"fmt"
	"strings"

	rsa_crypto "github.com/nikitavaulin/rsa-golang/internal/rsacrypto"
	"github.com/nikitavaulin/rsa-golang/internal/rwfile"
)

func (app *App) decryptData() {
	fmt.Println("\n--- Decrypt Data ---")

	if !app.hasKeys {
		fmt.Println("𐄂 No keys loaded. Please generate or load keys first.")
		return
	}

	// reading as cipher
	cipherStr, err := rwfile.Read(INPUT_DATA_FILE)
	if err != nil {
		fmt.Printf("𐄂 Error reading input file '%s': %v\n", INPUT_DATA_FILE, err)
		return
	}

	if len(cipherStr) == 0 {
		fmt.Printf("𐄂 Input file '%s' is empty.\n", INPUT_DATA_FILE)
		return
	}

	fmt.Printf("➔ Read ciphertext from '%s' (%d bytes)\n", INPUT_DATA_FILE, len(cipherStr))
	fmt.Printf("Cipher preview: %s\n", string(cipherStr[:min(200, len(cipherStr))]))
	if len(cipherStr) > 200 {
		fmt.Printf("... (%d more characters)\n", len(cipherStr)-200)
	}

	cipher := cipherStr

	fmt.Printf("🗸 Parsed ciphertext: %d bytes\n", len(cipher))

	fmt.Println("\n➔ Decrypting...")
	plaintext := rsa_crypto.Decrypt(cipher, app.secretKey)

	fmt.Printf("🗸 Decryption completed!\n")
	fmt.Printf("   Encrypted size: %d bytes\n", len(cipher))
	fmt.Printf("   Decrypted size: %d bytes\n", len(plaintext))

	fmt.Println("\n--- Decrypted Content ---")
	fmt.Println(string(plaintext))
	fmt.Println(strings.Repeat("-", 50))

	err = rwfile.OverWrite(OUTPUT_DATA_FILE, []byte(string(plaintext)))
	if err != nil {
		fmt.Printf("𐄂 Failed to save decrypted data to '%s'\n", OUTPUT_DATA_FILE)
	}
	fmt.Printf("🗸 Decrypted data saved to '%s'\n", OUTPUT_DATA_FILE)

	saveOutput := app.getUserInput("Do you want to save decrypted data to another file? (y/n): ")
	if strings.ToLower(saveOutput) == "y" {
		outputFile := app.getUserInput("Enter output filename (default: decrypted.txt): ")
		if outputFile == "" {
			outputFile = "decrypted.txt"
		}
		if err := rwfile.OverWrite(outputFile, plaintext); err != nil {
			fmt.Printf("𐄂 Error saving decrypted data: %v\n", err)
		} else {
			fmt.Printf("🗸 Decrypted data saved to '%s'\n", outputFile)
		}
	}
}
