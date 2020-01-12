// Package lexer is for lexical analysis of program text from the Monkey
// programming language. It is meant to produce lexical tokens.
//
// To use, create a Lexer with some program text and call NextToken to start
// producing tokens:
//  l := lexer.New(`let x = 5;
//  let y = 10;
//  let z = print(x + y);`)
//  for tok := l.NextToken; tok.Type != token.EOF; toke = l.NextToken() {
//      fmt.Printf("%s\n", tok)
//  }
package lexer
