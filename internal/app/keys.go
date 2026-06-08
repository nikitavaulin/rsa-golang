package app

import (
	"fmt"
	"strings"

	"github.com/nikitavaulin/rsa-golang/internal/keys"
	"github.com/nikitavaulin/rsa-golang/internal/rwfile"
)

func (app *App) generateNewKeys() {
	fmt.Println("\n--- Generate New Key Pair ---")

	fmt.Println("Generating new RSA keys...")
	public, secret, err := keys.GenerateKeys()
	if err != nil {
		fmt.Printf("Error generating keys: %v\n", err)
		return
	}

	app.publicKey = public
	app.secretKey = secret
	app.hasKeys = true

	fmt.Println("🗸 New keys generated successfully!")
	fmt.Printf("   Public Key (e,n): e=%s, n=%s\n", app.publicKey.E.String(), app.publicKey.N.String())
	fmt.Printf("   Secret Key (d,n): d=***, n=%s\n", app.secretKey.N.String())

	saveNow := app.getUserInput("Do you want to save these keys to files? (y/n): ")
	if strings.ToLower(saveNow) == "y" {
		app.saveKeysToFiles()
	}
}

func (app *App) readKeyFiles() {
	fmt.Println("\n--- Read Keys from Files ---")

	publicData, err := rwfile.Read(PUBLIC_KEY_FILE)
	if err != nil {
		fmt.Printf("Error reading public key file: %v\n", err)
		return
	}

	secretData, err := rwfile.Read(SECRET_KEY_FILE)
	if err != nil {
		fmt.Printf("Error reading secret key file: %v\n", err)
		return
	}

	publicKey, err := keys.ParsePublicKey(string(publicData))
	if err != nil {
		fmt.Printf("Error parsing public key: %v\n", err)
		return
	}

	secretKey, err := keys.ParseSecretKey(string(secretData))
	if err != nil {
		fmt.Printf("Error parsing secret key: %v\n", err)
		return
	}

	app.publicKey = publicKey
	app.secretKey = secretKey
	app.hasKeys = true

	fmt.Println("🗸 Keys successfully loaded")
}

func (app *App) saveKeysToFiles() {
	fmt.Println("\n--- Save Keys to Files ---")

	if !app.hasKeys {
		fmt.Println("𐄂 No keys to save. Please generate or load keys first.")
		return
	}

	publicData := []byte(app.publicKey.String())
	if err := rwfile.OverWrite(PUBLIC_KEY_FILE, publicData); err != nil {
		fmt.Printf("Error saving public key: %v\n", err)
		return
	}

	secretData := []byte(app.secretKey.String())
	if err := rwfile.OverWrite(SECRET_KEY_FILE, secretData); err != nil {
		fmt.Printf("Error saving secret key: %v\n", err)
		return
	}

	fmt.Printf("🗸 Keys saved to:\n")
	fmt.Printf("   Public key: %s\n", PUBLIC_KEY_FILE)
	fmt.Printf("   Secret key: %s\n", SECRET_KEY_FILE)
}
