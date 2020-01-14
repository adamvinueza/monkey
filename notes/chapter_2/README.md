# Notes on Chapter Two

## godoc does not currently support Go modules
I want to generate documentation for my code, but `godoc` does not currently
work with code organized via Go modules. This is a bummer, because I've tried to
document my code appropriately and it would be nice to see the generated
documentation. I think it's in v1.14, which isn't stable yet.

## My first exposure to abstract syntax trees
Spoiler alert: abstract syntax trees are trees!

I only belatedly realized that the return value of the parser's `ParseProgram`
function is not an abstract syntax tree, but instead a _slice_ of abstract
syntax trees. That's because a computer program is in general a list of
statements, and each statement is representable as an abstract syntax tree.

So statements are abstract syntax trees, and there's nothing especially
complicated about trees. So far I've worked with only LET and RETURN statements
and wth expressions, although I peeked ahead and saw BLOCK statements, which are
used when parsing IF expressions. But what I'm trying to say is that there
aren't that many different kinds of statements, and I think everything that
isn't a LET statement is an expression, meaning it produces a value. This vastly
simplifies parsing, by the way. A surprising consequence is that languages with
functions as expressions ("first-class objects") are simpler and easier to write
parsers for than languages without.

## The LOWEST precedence pseudo-operator

Trying to understand how this works. More soon.

