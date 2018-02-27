package runner

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hitochan777/monkey/evaluator"
	"github.com/hitochan777/monkey/lexer"
	"github.com/hitochan777/monkey/object"
	"github.com/hitochan777/monkey/parser"
)

const MONKEY_FACE = `
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-""" """-./, -' /
   "'-' /_   ^ ^   _\ '-'"
       |  \._   _./  |
       \   \ "~" /   /
        '._ '-=-' _.'
           '~---~'
`

const PROPMT = ">> "

func Start(in io.Reader, out io.Writer, isRepl bool) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		if isRepl {
			fmt.Printf(PROPMT)
		}

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if isRepl && evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parse errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
