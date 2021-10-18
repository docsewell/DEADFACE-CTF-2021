package main

import (
	"crypto/md5"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func findFilesWithExt(ext string) []string {
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}

// EncryptDecrypt runs a XOR encryption on the input string, encrypting it if it hasn't already been,
// and decrypting it if it has, using the key provided.
func encryptDecryptXor(input string, key byte) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key)
	}
	return output
}

func encryptCaesar(plaintext string, key int) string {
	return rotateText(plaintext, key)
}

func decryptCaesar(ciphertext string, key int) string {
	return rotateText(ciphertext, -key)
}

// Takes a string and rotates each character by the provided amount.
func rotateText(inputText string, rot int) string {
	rot %= 26
	rotatedText := []byte(inputText)

	for index, byteValue := range rotatedText {
		if byteValue >= 'a' && byteValue <= 'z' {
			rotatedText[index] = lowerCaseAlphabet[(int((26+(byteValue-'a')))+rot)%26]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			rotatedText[index] = upperCaseAlphabet[(int((26+(byteValue-'A')))+rot)%26]
		}
	}
	return string(rotatedText)
}

func getMd5HashByte(data []byte) []byte {
	hash := md5.Sum(data)
	return (hash[:])
}

func fetchKey(url string) []byte {
	//resp, err := http.Get(url)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "DEADFACE_LLABS_CRYPTOWARE/6.69")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	hash_key := getMd5HashByte(body)

	return  hash_key
}

func encryptAes(text []byte, key[]byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	encryptedData := gcm.Seal(nonce, nonce, text, nil)

	return encryptedData, nil
}

func decryptAes(ciphertext []byte, key[]byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		panic(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return plaintext, nil
}

func main() {
	// Get the decryption key from the commandn line
	if len(os.Args) != 2 {
		fmt.Println("To decrypt your files, enter the key as a 32-character hex string.")
		os.Exit(-2)
	}

	// Decode the hex string
	fileKey, err := hex.DecodeString(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Find all files in the current directory that match the
	// extension pattern
	fileList := findFilesWithExt(".oodev")

	// Loop through the *.oodev files.  Encrypt them, then write out
	// the data with a CAESAR(3) decrypted file name.
	// Remove the original file.
	for _, fileName := range fileList {
		encFileName := decryptCaesar(fileName, 3)

		fileData, _ := os.ReadFile(fileName)
		decryptedFileData, _ := decryptAes(fileData, fileKey)
		os.WriteFile(encFileName, decryptedFileData, 0644)
		os.Remove(fileName)
	}

	fmt.Println("LYTTON LABS: You got lucky this time...")

}