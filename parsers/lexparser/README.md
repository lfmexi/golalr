golalr/parser/lexparser
=======================

This parser extends the functionallity of `golalr/prattparser`. It recognizes the following BNF ambiguous grammar:

```
<REGEXP> := '(' <REGEXP> ')'
        | '['char '-' char']'
        | <REGEXP><REGEXP>
        | <REGEXP> '|' <REGEXP>
        | ^<REGEXP>
        | <REGEXP>*
        | <REGEXP>+
        | <REGEXP>?
        | char 
```
It adds precedence values to the operators in order to break the ambiguous derivation as shown in the following table:

|   Operators             | Precedence Type    | Precedence value |
|-------------------------|:------------------:|-----------------:|
|       `-`               |    Range           |        1         |
| `'('` and `')'`         |    Grouping        |        2         |
| `'['` and `']'`         |    Grouping Braces |        3         |
|      char               |    Concat          |        4         |
|      ``\|``             |    Or              |        5         |
|      `'^'`              |    Prefix          |        6         |
|  `'?'`, `'*'` and `'+'` |    Postfix         |        7         |

## Requirements

* Go 1.7+

## API

Check the Godocs of this package at godoc.org:

https://godoc.org/github.com/lfmexi/golalr/parsers/lexparser

## How to use it?

### Install it

```bash
go get github.com/lfmexi/golalr/parsers/lexparser
```

### Get a Simple Lexer for your input

```go
import "github.com/lfmexi/golalr/parsers/lexparser"

...

input := "(a|b)*abb"

scanner := lexparser.NewLexScanner(input)
```

### Create a parser and parse your input

```go
parser := lexparser.NewLexerParser(scanner)

expression, err := parser.Parse()

if err != nil {
    // Manage your ParseError
}

// This expression is a lexparser.LexerExpression struct
// This will represent the AST in a JSON string
fmt.Println(expression.ToString())
```
