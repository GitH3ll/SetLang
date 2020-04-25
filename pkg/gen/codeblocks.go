package gen

import "fmt"

const (
	prog = `package main

import (
	"fmt"
)

type Set struct {
	elems map[string]Value
}

type Value struct {
	val string
}

type Bool struct {
	val string
}

func NewBool(v string) Bool {
	return Bool{val: v}
}

func NewValue(v string) Value {
	return Value{val: v}
}

func NewSet(vals ...Value) Set {
	elems := make(map[string]Value)
	for _, v := range vals {
		elems[v.val] = v
	}
	return Set{elems: elems}
}

func (v Value) get() string {
	return v.val
}

func (v Value) PRINT() Value {
	fmt.Println(v.val)
	return v
}

func (s Set) PRINT() Set {
	fmt.Print("{")
	for _, v := range s.elems {
		fmt.Print(v.val + ", ")
	}
	fmt.Println("}")
	return s
}

func (s Set) PLUS(val Value) Set {
	s.elems[val.get()] = val
	return s
}

func dump(a interface{}) {

}

`
)

type Set struct {
	elems map[string]Value
}

type Value struct {
	val string
}

type Bool struct {
	val string
}

func NewBool(v string) Bool {
	return Bool{val: v}
}

func NewValue(v string) Value {
	return Value{val: v}
}

func NewSet(vals ...Value) Set {
	elems := make(map[string]Value)
	for _, v := range vals {
		elems[v.val] = v
	}
	return Set{elems: elems}
}

func (v Value) get() string {
	return v.val
}

func (v Value) PRINT() Value {
	fmt.Println(v.val)
	return v
}

func (s Set) PRINT() Set {
	fmt.Print("{")
	for _, v := range s.elems {
		fmt.Print(v.val + ", ")
	}
	fmt.Println("}")
	return s
}

func (s Set) PLUS(val Value) Set {
	s.elems[val.get()] = val
	return s
}

func dump(a interface{}) {

}
