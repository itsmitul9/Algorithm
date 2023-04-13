package decrypt

func Hashing(s string) string{
	decryptedStr := ""
	
	for _, c := range s{
		asciiStr := int(c)
		char := string(asciiStr - 3)
		decryptedStr += char
	}
	return decryptedStr
}