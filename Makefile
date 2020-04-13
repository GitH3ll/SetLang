.PHONY: gocc
gocc:
	gocc -p "github.com/GitH3ll/SetLang/pkg/gocc/cc" pkg/gocc/grammar.bnf
	rm -rf pkg/gocc/cc
	mkdir pkg/gocc/cc
	mv -f -t pkg/gocc/cc errors parser lexer token util