package filewalker

import (
	"Prince-Decryptor/configuration"
	"Prince-Decryptor/decryption"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func DecryptDirectory(dirPath string) {
	var wg sync.WaitGroup
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			for _, excluded := range configuration.ExcludedDirectories {
				if strings.EqualFold(strings.ToLower(filepath.Base(path)), excluded) {
					return filepath.SkipDir
				}
			}
		}

		if !info.IsDir() {
			fileExt := filepath.Ext(path)
			fileName := strings.ToLower(filepath.Base(path))

			if fileName == "decryption instructions.txt" {
				_ = os.Remove(path)
			} else if strings.HasSuffix(fileExt, configuration.EncryptedExtension) {
				wg.Add(1)
				go func() {
					defer wg.Done()
					decryption.DecryptFile(path)
					_ = os.Rename(path, strings.TrimSuffix(path, configuration.EncryptedExtension))
				}()
			}
		}

		return nil
	})

	wg.Wait()
}
