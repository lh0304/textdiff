package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"
)

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func hashFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil)), nil
}

func commit(filename string, message string) {
	if err := os.MkdirAll("commits", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	hash, err := hashFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS
	commitFile := fmt.Sprintf("commits/%s_%s.txt", timestamp, hash)
	content := fmt.Sprintf("Message: %s \nTime:%s\nFile:%s\nData:\n%s\n", message, timestamp, filename, data)
	if err := os.WriteFile(commitFile, []byte(content), 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Committed: %s\n", commitFile)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <filename><commit message>", os.Args[0])
	}
	filename := os.Args[1]
	message := os.Args[2]
	commit(filename, message)
}
