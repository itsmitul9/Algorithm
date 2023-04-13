package encrypt

func Hashing(s string) string{
	
	var encryptedStr string
	
	for _, c := range s{
		asciiStr := int(c)
		char := (asciiStr + 3)
		encryptedStr = encryptedStr + string(char)
	}
	return encryptedStr
}