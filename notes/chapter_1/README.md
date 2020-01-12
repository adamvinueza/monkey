# Notes on Chapter One

## Stateful structs require pointer-receiver methods
I hadn't understood [the semantics of struct methods and receivers in
Go](https://golang.org/doc/faq#methods_on_values_or_pointers), so my first
attempt at modifying the `NextToken` method on `Lexer` to read identifiers
caused my tests to fail. This was because the `Lexer` read position wasn't
changing after reading all the characters in an identifier, and _that_ was
because `readIdentifier` took a `Lexer` value as a receiver. Simply put:
`readIdentifier` was modifying the state of a copy of the `Lexer` instance, not
the instance itself.

## Generalizing reading multiple-character tokens
If we are reading multiple-character tokens, we just want to read until we
encounter an unfamiliar character and return the string we find. So detecting an
identifier or a number can be generalized so that the same function can be used
to update the `Lexer`: just pass in the character-identifying function. If we
want to identify a number, pass in `isNumber`; to identify an identifier, pass
in `isLetter`.

## Finished!

Extended the language in the requisite ways and wrote the preliminary REPL
program. Now on to Chapter 2!

## What I've built so far

I've built a lexer that transforms a string of text into tokens of the Monkey
language, in order of occurrence in the text. The key data structures for tokens
are:
```
type Token struct {
    // Type represents the type of token, e.g. INT, TRUE, IDENT, IF, LET.
    Type    TokenType

    // Literal represents the string value of the token as it occurs in the
    // program text, e.g. "int","true", "x", "if", "let".
    Literal string
}

// TokeType is just an alias for a string, shorthand for the enumeration of
// token type names.
type TokenType string
```
The lexer is a stateful struct designed to consume text and produce tokens.
```
type Lexer struct {
    input        string // program text input
    position     int    // current index of the lexer in the program text
    readPosition int    // position + 1
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    // [ consumes the input text and produces a token from it ... ]
    return tok
}
```
The lexer is essentially a string iterator that produces tokens along the way.
This is how it produces a token:
   1. Skips any white space it sees.
   2. Looks at the current character:
      - if it's a single-character Token, creates that Token.
      - if it's a two-character non-identifier Token, peeks at the next
        character and creates the right one, then consumes a character.
        (Consuming a character bumps the internal position of
        the lexer.)
      - if it's a multi-character identifier Token (such as a keyword or
        variable identifier), consumes characters until it comes to a
        non-identifier character, then creates an identifier or keyword token.
      - if it can't identify the type of Token in any of these processes, it
        creates a special type of Token indicating it's illegal.
    3. Consumes the current character.
I also created a primitive REPL, with a prompt and an ability to read what's
typed in and print out the lexed output.
