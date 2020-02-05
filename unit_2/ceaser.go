package unit_2

/*
Decipher the quote from Julius Caesar:
L fdph, L vdz, L frqtxhuhg.
—Julius Caesar
Your program will need to shift uppercase and lowercase letters by –3. Remember that
'a' becomes 'x', 'b' becomes 'y', and 'c' becomes 'z', and likewise for uppercase letters.
*/
import (
	"fmt"
	//"strings"
)

func CeaserCipher(s string) {
	var l = len(s)
	for i := 0; i < l; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			if s[i] <= 'c' {
				fmt.Printf("%c", s[i]+23)
			} else {
				fmt.Printf("%c", s[i]-3)
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			if s[i] <= 'C' {
				fmt.Printf("%c", s[i]+23)
			} else {
				fmt.Printf("%c", s[i]-3)
			}
		} else {
			fmt.Printf("%c", s[i])
		}
	}
	fmt.Println()
}
