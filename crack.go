package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"os"
	"sync"
)

// Generate hash for a password
func generateHash(password string, algorithm string) string {
	if algorithm == "sha1" {
		h := sha1.New()
		h.Write([]byte(password))
		return fmt.Sprintf("%x", h.Sum(nil))
	} else if algorithm == "md5" {
		h := md5.New()
		h.Write([]byte(password))
		return fmt.Sprintf("%x", h.Sum(nil))
	} else {
		panic(fmt.Sprintf("Unsupported algorithm: %s", algorithm))
	}
}

func main() {
	// Get user input for stored password hash and read user passwords from a file
	var storedPasswordHash string
	fmt.Print("Enter the stored password hash: ")
	fmt.Scan(&storedPasswordHash)

	var algorithm string
	if len(storedPasswordHash) == 32 {
		algorithm = "md5"
	} else if len(storedPasswordHash) == 40 {
		algorithm = "sha1"
	} else {
		panic("Unable to detect hash algorithm")
	}

	var passwordFile string
	fmt.Print("Enter the file name containing the list of passwords: ")
	fmt.Scan(&passwordFile)

	// Read the passwords from the file
	file, err := os.Open(passwordFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var passwords []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Check each password and output the result
	var wg sync.WaitGroup
	found := make(chan string, len(passwords))

	for _, password := range passwords {
		wg.Add(1)
		go func(password string) {
			passwordHash := generateHash(password, algorithm)
			if passwordHash == storedPasswordHash {
				found <- password
			}
			wg.Done()
		}(password)
	}

	wg.Wait()
	close(found)

	for password := range found {
		fmt.Printf("Password %s is correct\n", password)
		return
	}

	fmt.Println("None of the passwords in the list are correct")
}
