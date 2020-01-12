// Package parser is for parsing program text from the Monkey programming
// language. It is meant to produce abstract syntax trees.
//
// To use, create a Parser and call ParseProgram:
//  p := parser.New(lexer.New(`let x = 5;
//  let y = 10;
//  let z = print(x + y);`))
//  s := p.ParseProgram()
package parser
