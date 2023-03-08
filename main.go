package main

import (
	"Jade/Core"
	jade "Jade/Core/parser"
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"time"
)

func main() {
	start := time.Now()

	input, _ := antlr.NewFileStream("Examples/test.jd")

	lexer := jade.NewJadeLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)

	p := jade.NewJadeParser(tokens)
	p.BuildParseTrees = true
	tree := p.Parse()

	visitor := Core.NewVisitor(Core.NewEnv())
	visitor.Visit(tree)

	elapsed := time.Since(start)
	fmt.Printf("Time to execute: %v\n", elapsed)

	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
