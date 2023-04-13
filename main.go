package main

import (
	"fmt"
	"algorithm/encrypt"
	"algorithm/decrypt"
)

func main(){
	encryptedString := encrypt.Hashing("SecretCode")
	fmt.Println("Encrypted string:", encryptedString)
	fmt.Println("Decrypted string:",decrypt.Hashing(encryptedString))
}

//to generate executable and run it from anywhere
//Algorithm % go env GOPATH
//Algorithm %  export PATH=$PATH:<//Users/mitul/go/bin>
// Algorithm % go build 
// Algorithm % go install

// cd /Users/mitul/go/bin and check executable present or not 

