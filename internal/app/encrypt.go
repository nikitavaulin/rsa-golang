package app

import (
	"fmt"

	rsa_crypto "github.com/nikitavaulin/rsa-golang/internal/rsacrypto"
	"github.com/nikitavaulin/rsa-golang/internal/rwfile"
)

func (app *App) encryptData() {
	fmt.Println("\n--- Encrypt Data ---")

	if !app.hasKeys {
		fmt.Println("𐄂 No keys loaded. Please generate or load keys first.")
		return
	}

	plaintext, err := rwfile.Read(INPUT_DATA_FILE)
	if err != nil {
		fmt.Printf("𐄂 Error reading input file '%s': %v\n", INPUT_DATA_FILE, err)
		fmt.Println("Please create 'input.txt' with your plain text message.")
		return
	}

	if len(plaintext) == 0 {
		fmt.Printf("𐄂 Input file '%s' is empty.\n", INPUT_DATA_FILE)
		return
	}

	fmt.Printf("➔ Read plaintext from '%s' (%d bytes)\n", INPUT_DATA_FILE, len(plaintext))
	fmt.Printf("Plaintext preview: %s\n", string(plaintext[:min(100, len(plaintext))]))
	if len(plaintext) > 100 {
		fmt.Printf("... (%d more bytes)\n", len(plaintext)-100)
	}

	fmt.Println("\n➔ Encrypting...")
	cipher := rsa_crypto.Encrypt(plaintext, app.publicKey)

	fmt.Printf("🗸 Encryption completed!\n")
	fmt.Printf("   Original size: %d bytes\n", len(plaintext))
	fmt.Printf("   Encrypted size: %d bytes\n", len(cipher))

	if err := rwfile.OverWrite(OUTPUT_DATA_FILE, cipher); err != nil {
		fmt.Printf("𐄂 Error saving encrypted data: %v\n", err)
		return
	}

	fmt.Printf("🗸 Encrypted data saved to '%s' as byte array format\n", OUTPUT_DATA_FILE)
	fmt.Printf("\n  Cipher preview: %s\n", cipher[:min(200, len(cipher))])
	if len(cipher) > 200 {
		fmt.Printf("... (%d more characters)\n", len(cipher)-200)
	}
}
