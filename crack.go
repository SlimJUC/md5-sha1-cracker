package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

func generateHash(password string, algorithm string) string {
	switch algorithm {
	case "sha1":
		h := sha1.New()
		h.Write([]byte(password))
		return fmt.Sprintf("%x", h.Sum(nil))
	case "md5":
		h := md5.New()
		h.Write([]byte(password))
		return fmt.Sprintf("%x", h.Sum(nil))
	default:
		panic(fmt.Sprintf("Unsupported algorithm: %s", algorithm))
	}
}

func sanitize(input string) string {
	var output []rune
	for _, r := range input {
		if unicode.IsPrint(r) && !unicode.IsSpace(r) {
			output = append(output, r)
		}
		if len(output) >= 30 {
			break
		}
	}
	return string(output)
}

func main() {
	var storedPasswordHash, passwordFile string
	fmt.Print("Enter the stored password hash: ")
	fmt.Scan(&storedPasswordHash)

	storedPasswordHash = strings.ToLower(storedPasswordHash)

	var algorithm string
	switch len(storedPasswordHash) {
	case 32:
		algorithm = "md5"
	case 40:
		algorithm = "sha1"
	default:
		panic("Unable to detect hash algorithm based on hash length")
	}

	fmt.Print("Enter the file name containing the list of passwords: ")
	fmt.Scan(&passwordFile)

	file, err := os.Open(passwordFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup
	found := make(chan bool, 1)

	go func() {
		for scanner.Scan() {
			password := scanner.Text()
			password = strings.TrimSpace(password)
			displayPassword := sanitize(password)
			fmt.Printf("\r\033[KTesting password: %s", displayPassword)
			wg.Add(1)
			go func(password string) {
				defer wg.Done()
				passwordHash := generateHash(password, algorithm)
				if passwordHash == storedPasswordHash {
					fmt.Printf("\r\033[KPassword %s is correct\n", displayPassword)
					found <- true
				}
			}(password)
		}
		wg.Wait()
		close(found)
	}()

	if _, ok := <-found; !ok {
		fmt.Printf("\r\033[KNone of the passwords in the list are correct\n")
	}
}
