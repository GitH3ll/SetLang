package main

import (
	"io/ioutil"
	"log"

	"github.com/GitH3ll/SetLang/pkg/ast"
	"github.com/GitH3ll/SetLang/pkg/checker"
	"github.com/GitH3ll/SetLang/pkg/gen"
	"github.com/GitH3ll/SetLang/pkg/gocc/cc/lexer"
	"github.com/GitH3ll/SetLang/pkg/gocc/cc/parser"
)

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Panicf("failed to read input.txt: %s", err)
	}

	prog := Parse(string(buf))
	TypeCheck(prog)
	code := gen.GenWrapper(prog)
	output := code.String()
	err = ioutil.WriteFile("output.go", []byte(output), 0777)
	if err != nil {
		log.Println(err)
	}
}

func Parse(input string) *ast.Program {
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	node, err := p.Parse(l)
	check(err)
	program, _ := node.(*ast.Program)
	return program
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TypeCheck(program *ast.Program) {
	err := checker.Checker(program)
	check(err)
}
