package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/matheusgb/marmota/lexer"
	"github.com/matheusgb/marmota/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lexerVar := lexer.New(line)
		for tokenVar := lexerVar.NextToken(); tokenVar.Type != token.EOF; tokenVar = lexerVar.NextToken() {
			fmt.Printf("%+v\n", tokenVar)
		}
	}
}
