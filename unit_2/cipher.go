package unit_2

/*
To send ciphered messages, write a program that ciphers plain text using a keyword:
plainText := "your message goes here" keyword := "GOLANG"
Bonus: rather than write your plain text message in uppercase letters with no spaces, use the strings.Replace and strings.
ToUpper functions to remove spaces and uppercase the string before you cipher it.
Once you’ve ciphered a plain text message, check your work by deciphering the ciphered text with the same keyword.
Use the keyword "GOLANG" to cipher a message and post it to the forums for Get Program- ming with Go at forums.manning.com/forums/get-programming-with-go.
 */

import(
	"fmt"
	"strings"
)

func CipherMe(){
	key := "golang"
	cipher := ""
	message := "Very Important Message!!!"
	index:= 0
	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	key = strings.ToUpper(strings.Replace(key, " ", "", -1))
	lengthMessage := len(message)
	for i := 0; i < lengthMessage; i++ {
		c := message[i]
		if c >= 'A' && c <= 'Z' {
			c = c - 'A'
			k := key[index] - 'A'
			c = (c+k)%26 + 'A'
			index = index % len(key)
		}
		cipher = cipher + string(c)
	}
	fmt.Println(cipher)
}
