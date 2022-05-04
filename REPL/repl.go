package repl

import (
	"bufio"
	"fmt"
	"io"

	evaluator "github.com/ZooeyLang/Evaluator"
	lexer "github.com/ZooeyLang/Lexer"
	object "github.com/ZooeyLang/Object"
	parser "github.com/ZooeyLang/Parser"
)

// READ EVAL PRINT LOOP

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {

		fmt.Printf(PROMPT)
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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

		// Read user input, until encounter a new line
	}
}

const Mellus = `
	▓█████  ██▀███   ██▀███   ▒█████   ██▀███  
	▓█   ▀ ▓██ ▒ ██▒▓██ ▒ ██▒▒██▒  ██▒▓██ ▒ ██▒
	▒███   ▓██ ░▄█ ▒▓██ ░▄█ ▒▒██░  ██▒▓██ ░▄█ ▒
	▒▓█  ▄ ▒██▀▀█▄  ▒██▀▀█▄  ▒██   ██░▒██▀▀█▄  
	░▒████▒░██▓ ▒██▒░██▓ ▒██▒░ ████▓▒░░██▓ ▒██▒
	░░ ▒░ ░░ ▒▓ ░▒▓░░ ▒▓ ░▒▓░░ ▒░▒░▒░ ░ ▒▓ ░▒▓░
	░ ░  ░  ░▒ ░ ▒░  ░▒ ░ ▒░  ░ ▒ ▒░   ░▒ ░ ▒░
	░     ░░   ░   ░░   ░ ░ ░ ░ ▒    ░░   ░ 
	░  ░   ░        ░         ░ ░     ░   
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, Mellus)
	io.WriteString(out, "Woops! We ran into some wrong business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
