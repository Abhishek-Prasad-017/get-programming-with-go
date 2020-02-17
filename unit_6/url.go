package unit_6

/*
In the Go standard library, there’s a function to parse web addresses (see golang.org/ pkg/net/url/#Parse).
Display the error that occurs when url.Parse is used with an invalid web address, such as one containing a space: https://a b.com/.
Use the %#v format verb with Printf to learn more about the error. Then perform a *url.Error type assertion to access and print the fields of the underlying structure.
 */

import(
	"fmt"
	"net/url"
	"os"
)

func Url() {
	u, err := url.Parse("https://a b.com/")

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%#v\n", err)
		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op:", e.Op)
			fmt.Println("URL:", e.URL)
			fmt.Println("Err:", e.Err)
		}

		os.Exit(1)
	}
	fmt.Println(u)
}

