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

## Pratt parsing

Started reading Ball's discussion of Pratt parsing, and walked through the
example he gave, which I understand _mostly_ but not completely: I get all the
individual steps, but I don't see the overall picture clearly just yet.

I've started reading about Pratt parsing by way of familiarizing myself with it,
hoping that will help. Eli Bendersky [has an article about
it](https://eli.thegreenplace.net/2010/01/02/top-down-operator-precedence-parsing),
that's based on [another](https://effbot.org/zone/simple-top-down-parsing.htm)
by Frederick Lundh; the problem with those articles is that both involve what
looks to me like fairly convoluted and hacky Python 2 code; I want code I can
read fairly easily.

Today I happened on Bob Nystrom's [article on Pratt parsing]() which is much, much
better. The code is in Java, which is super verbose but in Nystrom's hand is
very easy to read; and his discussion seems to be a _lot_ cleaner.

