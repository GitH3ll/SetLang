#SetLang compiler
A compiler for SetLang PL.
##Structure and packages
###gocc
Package gocc contains grammar.bnf and gocc generated files.
###ast
Package ast contains functions and related types that are invoked while parsing
the input file.
###checker
Package checker contains typecheck logic that is invoked after ast generation is
complete.
###gen
Package gen contains logic of translating AST into Go code.

##Usage
To build and use the compiler you must have Go installed.
* See examples directory. It contains code that is guaranteed to work.
* Create input.txt file and put the source code there.
* Run compile.sh and see the result
