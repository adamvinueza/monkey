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
