package main

import (
	filewalker "Prince-Decryptor/iterator"
	"os/exec"
	"syscall"
)

func main() {
	filewalker.DecryptDirectory("/")
	setWallpaper()
}

func setWallpaper() {
	setWallpaperCmd := exec.Command("powershell", "-Command", `Add-Type -TypeDefinition 'using System; using System.Runtime.InteropServices; public class Wallpaper { [DllImport("user32.dll", CharSet = CharSet.Auto)] public static extern int SystemParametersInfo(int uAction, int uParam, string lpvParam, int fuWinIni); public static void Set(string path) { SystemParametersInfo(20, 0, path, 3); } }'; [Wallpaper]::Set('C:\Windows\Web\Wallpaper\Windows\img19.jpg')`)
	setWallpaperCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := setWallpaperCmd.Run()
	if err != nil {
		return
	}
}
