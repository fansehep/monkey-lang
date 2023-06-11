package repl

import (
	"bufio"
	"fmt"
	"github.com/fansehep/monkey-lang/executor"
	"github.com/fansehep/monkey-lang/lexer"
	"github.com/fansehep/monkey-lang/parser"
	"io"
)

const PROMPT = ">> "

// func Start(in io.Reader, out io.Writer) {
// 	scaner := bufio.NewScanner(in)

// 	for {
// 		fmt.Fprintf(out, PROMPT)
// 		scanned := scaner.Scan()
// 		if !scanned {
// 			return
// 		}
// 		line := scaner.Text()
// 		l := lexer.New(line)
// 		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
// 			fmt.Fprintf(out, "%+v\n", tok)
// 		}
// 	}
// }

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := executor.Eval(program)
		if evaluated != nil {
			io.WriteString(out, program.String())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
