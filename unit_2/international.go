package unit_2

/*
Cipher the Spanish message “Hola Estación Espacial Internacional” with ROT13. Mod- ify listing 9.7 to use the range keyword.
Now when you use ROT13 on Spanish text, char- acters with accents are preserved.
*/

import (
	"fmt"
)

func Internatioanl(s string) {
	//message := "uv vagreangvbany fcnpr fgngvba"
	for _, c := range s {

		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
