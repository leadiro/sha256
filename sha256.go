package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify an input file")
		fmt.Println("sha256 in.txt")
		return
	}
	//Open the file
	name := os.Args[1]
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	//Scan through line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		hasher := sha256.New()
		hasher.Write([]byte(line))
		sha := hex.EncodeToString(hasher.Sum(nil))
		fmt.Println(sha)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}
