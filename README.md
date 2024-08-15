# Prince Ransomware

## Detection Update (15/08/24)
Prince now has a Windows Defender flag, namely "Ransom:Win64/PrinceRansom.YAA!MTB". This means that Prince Ransomware will no longer bypass Windows Defender without modifications to remove the signature.
![image](https://github.com/user-attachments/assets/d686558c-acb9-4354-9b38-e7442f2bf0dc)

If, for whatever reason, bypassing Windows Defender is a priority for you, contact me on [Telegram](https://t.me/secdbg) and I will accept payment for any changes you may require.

## Brief Overview
Prince is a ransomware written from scratch in Go. It uses a mixture of ChaCha20 and ECIES cryptography in order to encrypt files securely so that they cannot be recovered by traditional recovery tools. Files which have been encrypted by Prince can only be decrypted using the corresponding decryptor.

## Installation & Setup
### Pre-requisites:
- [The Go Programming Language](https://go.dev)

### Compiling the Builder
- In order to compile the builder program, you must run the `Build.bat` file.
- This will automatically download the dependencies and build the `Builder.exe` file in the current directory.

### Building the Ransomware
- In order to build the encryptor and decryptor, you must run the `Builder.exe` program.
- Ensure that the builder is in the same directory as the `Encryptor` and `Decryptor` directories, as it will not be able to build them otherwise.
- The builder will generate a unique ECIES key pair and output the compiled executables to the current directory.
- The `Prince-Built.exe` file is the encryptor. Use caution when handling it as it can cause a lot of damage to your system.
- The `Decryptor-Built.exe` file is the decryptor. It will only decrypt files which were decrypted by the corresponding encryptor.

## Showcase
https://github.com/SecDbg/Prince-Ransomware/assets/73649897/433e6e4e-bc92-4553-a4d8-68745591058d

## Encryption Process
- The encryptor enumerates all drives on the system, and proceeds to iterate through each directory recursively.
- It ignores blacklisted files, directories and extensions.
- It generates a unique ChaCha20 key and nonce for each file, and encrypts the file using a pattern of 1 byte encrypted, 2 bytes unencrypted.
- It encrypts the ChaCha20 key and nonce using the ECIES public key, and prepends them to the start of the file.

## Benefits of ChaCha20 and ECIES
I chose this unique combination of encryption methods for several reasons:
- ChaCha20's stream-based approach allows for byte-by-byte encryption, enabling the pattern of 1 byte encrypted, 2 bytes unencrypted.
- ECIES offers similar security to RSA with shorter key lengths, making it a more efficient choice.

## Ethical Considerations
Releasing an open-source ransomware tool like Prince on GitHub presents ethical considerations, but it also offers significant benefits, particularly for security researchers:

- Open-source ransomware projects such as Prince can provide researchers with valuable insights into the techniques used by threat actors. This is critical for developing countermeasures and improving cybersecurity practices.

- Open-source ransomware projects such as Prince can provide security professionals with an easy-to-use tool to simulate real-world scenarios in a safe and ethical manner. This can help in identifying vulnerabilities and weaknesses in existing defenses.

- Open-source ransomware projects such as Prince can promote collaboration within the cybersecurity community. Researchers can share their findings, and collectively work towards developer more robust defenses against threat actors utilising ransomware.

## Disclaimer

### Important Notice: This tool is intended for educational purposes only.

- This software, referred to as Prince Ransomware, is provided strictly for educational and research purposes. Under no circumstances should this tool be used for any malicious activities, including but not limited to unauthorized access, data theft, or any other harmful actions.

### Usage Responsibility:

- By accessing and using this tool, you acknowledge that you are solely responsible for your actions. Any misuse of this software is strictly prohibited, and the creator (SecDbg) disclaims any responsibility for how this tool is utilized. You are fully accountable for ensuring that your usage complies with all applicable laws and regulations in your jurisdiction.

### No Liability:

- The creator (SecDbg) of this tool shall not be held responsible for any damages or legal consequences resulting from the use or misuse of this software. This includes, but is not limited to, direct, indirect, incidental, consequential, or punitive damages arising out of your access, use, or inability to use the tool.

### No Support:

- The creator (SecDbg) will not provide any support, guidance, or assistance related to the misuse of this tool. Any inquiries regarding malicious activities will be ignored.

### Acceptance of Terms:

- By using this tool, you signify your acceptance of this disclaimer. If you do not agree with the terms stated in this disclaimer, do not use the software.
