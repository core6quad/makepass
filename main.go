package main

// imports
import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
)

var Reader io.Reader // init random

var charset = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM!@#$%&^*><?/1234567890_-+=" // init password chars

func generate(length int) (string, error) {
	var passwd = make([]byte, length)
	max := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err // generate characters
		}
		passwd[i] = charset[n.Int64()]
	}
	return string(passwd), nil
}

func main() {
	var nosym = flag.Bool("nosymbols", false, "Don't include symbols?")
	var leng = flag.Int("length", 32, "Length of the password")
	var file = flag.String("file", "", "Store in a file? where")
	var simpleout = flag.Bool("simpleout", false, "output just the password?")
	flag.Parse() // parse launch params
	if !*simpleout {
		fmt.Println("Length: " + strconv.Itoa(*leng))
		fmt.Println("Generating...") // if not -simpleout, show more info
		fmt.Println("")
	}
	result, _ := generate(*leng)
	if *file != "" {
		fmt.Println("Saving into a file...") // if -file file specified, write to a file
		writetofile(result, *file)
	}
	if *nosym {
		charset = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890" // if -nosymbols is true then remove symbols
	}
	if *file == "" {
		fmt.Println(result) // dont print password if saving to a file
	}
}

func writetofile(data string, path string) {
	err := os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		panic(err)
	} // function for writing password to a file
}
