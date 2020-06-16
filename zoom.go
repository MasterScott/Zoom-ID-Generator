package main

import (
	"bufio"
	"crypto/rand"
	"math/big"
	"os"
	"strconv"

	"github.com/gookit/color"
)

func main() {
	color.FgYellow.Println(`
				███████  ██████   ██████  ███    ███ ███████ ██████  
				   ███  ██    ██ ██    ██ ████  ████ ██      ██   ██ 
				  ███   ██    ██ ██    ██ ██ ████ ██ █████   ██   ██ 
				 ███    ██    ██ ██    ██ ██  ██  ██ ██      ██   ██ 
				███████  ██████   ██████  ██      ██ ███████ ██████
				
`)
	color.FgWhite.Printf("Enter amount of codes to generate: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	amountCodes, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < amountCodes; i++ {
		randomint, _ := rand.Int(rand.Reader, big.NewInt(100))
		randomint1, _ := rand.Int(rand.Reader, big.NewInt(100))
		randomint10, _ := rand.Int(rand.Reader, big.NewInt(10))
		randomint101, _ := rand.Int(rand.Reader, big.NewInt(10))
		color.FgMagenta.Printf("[ZOOM] ---> %d%d-%d%d-%d%d%d\n", randomint1.Int64(), randomint101.Int64(), randomint.Int64(), randomint10.Int64(), randomint1.Int64(), randomint101.Int64(), randomint10.Int64())
	}
	color.FgGreen.Println("Press any key to exit...")
	scanner.Scan()
	os.Exit(0)
}
