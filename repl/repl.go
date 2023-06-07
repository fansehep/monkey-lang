package repl

import (
	"bufio"
	"fmt"
	"github.com/fansehep/monkey-lang/lexer"
	"github.com/fansehep/monkey-lang/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scaner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scaner.Scan()
		if !scanned {
			return
		}
		line := scaner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
