package unit_2
/*
Write a program to decipher the ciphered text shown in table 11.2. To keep it simple, all characters are uppercase English letters for both the text and keyword:
cipherText := "CSOITEUIWUIZNSROCNKFD" keyword := "GOLANG"
 The strings.Repeat function may come in handy. Give it a try, but also complete this exercise without importing any packages other than fmt to print the deci- phered message.
 Try this exercise using range in a loop and again without it. Remember that the range keyword splits a string into runes, whereas an index like keyword[0] results in a byte.
TIP You can only perform operations on values of the same type, but you can convert one type to the other (string, byte, rune).
 To wrap around at the edges of the alphabet, the Caesar cipher exercise made use of a comparison. Solve this exercise without any if statements by using modulus (%).

 */

import (
	"fmt"
)

func DecipherMe() {
	message := ""
	cipher := "CSOITEUIWUIZNSROCNKFD"
	key := "GOLANG"
	index := 0
	lengthCipher := len(cipher)
	for i := 0; i < lengthCipher; i++ {
		c := cipher[i] - 'A'
		k := key[index] - 'A'
		c = (c-k+26)%26 + 'A'
		message = message + string(c)
		index++
		index = index%len(key)
	}
	fmt.Println(message)
}