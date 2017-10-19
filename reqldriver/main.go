package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
)

func main() {
	filename := "" // A filename is optional
	var src string

	/*
		r.db("blog").table("users").filter({name: "Michel"});

		FILTER = 39     // from ql2.proto
		TABLE = 15
		DB = 14

		r.db("blog") =>
			[14, ["blog"]]

		r.db("blog").table("users") =>
			[15, [[14, ["blog"]], "users"]]

		r.db("blog").table("users").filter({name: "Michel"}) =>
			[39, [[15, [[14, ["blog"]], "users"]], {"name": "Michel"}]]
	*/

	src = `r.db("blog").table("users").filter({name: "Michel"});`

	// Parse some JavaScript, yielding a *ast.Program and/or an ErrorList
	program, err := parser.ParseFile(nil, filename, src, 0)
	if err != nil {
		fmt.Printf("oh dear: %s\n", err) // TODO:
	}

	var f func(n ast.Node)
	f = func(n ast.Node) {

		switch t := interface{}(n).(type) {
		case *ast.BlockStatement:
			for _, s := range t.List {
				f(s)
			}
		case *ast.CaseStatement:
			for _, c := range t.Consequent {
				f(c)
			}
		case *ast.CatchStatement:
			f(t.Body)
		case *ast.DoWhileStatement:
			f(t.Test)
			f(t.Body)
		case *ast.ExpressionStatement:
			f(t.Expression)
		case *ast.ForInStatement:
			f(t.Into)
			f(t.Source)
			f(t.Body)
		case *ast.ForStatement:
			f(t.Initializer)
			f(t.Update)
			f(t.Test)
			f(t.Body)
		case *ast.IfStatement:
			f(t.Test)
			f(t.Consequent)
			f(t.Alternate)
		case *ast.LabelledStatement:
			f(t.Statement)
		case *ast.ReturnStatement:
			f(t.Argument)
		case *ast.SwitchStatement:
			f(t.Discriminant)
			for _, s := range t.Body {
				f(s)
			}
		case *ast.ThrowStatement:
			f(t.Argument)
		case *ast.TryStatement:
			f(t.Body)
			f(t.Catch)
			f(t.Finally)
		case *ast.VariableStatement:
			for _, e := range t.List {
				f(e)
			}
		case *ast.WhileStatement:
			f(t.Test)
			f(t.Body)
		case *ast.WithStatement:
			f(t.Object)
			f(t.Body)
		case *ast.ArrayLiteral:
			for _, v := range t.Value {
				f(v)
			}
		case *ast.AssignExpression:
			f(t.Left)
			f(t.Right)
		case *ast.BinaryExpression:
			f(t.Left)
			f(t.Right)
		case *ast.BracketExpression:
			f(t.Left)
			f(t.Member)
		case *ast.CallExpression:
			fmt.Println("call")
			f(t.Callee)
			for _, a := range t.ArgumentList {
				f(a)
			}
		case *ast.ConditionalExpression:
			f(t.Test)
			f(t.Consequent)
			f(t.Alternate)
		case *ast.DotExpression:
			f(t.Left)
		case *ast.FunctionLiteral:
			f(t.Body)
		case *ast.NewExpression:
			f(t.Callee)
			for _, e := range t.ArgumentList {
				f(e)
			}
		case *ast.ObjectLiteral:
			for _, v := range t.Value {
				f(v.Value)
			}
		case *ast.Property:
			f(t.Value)

		case *ast.SequenceExpression:
			for _, e := range t.Sequence {
				f(e)
			}
		case *ast.UnaryExpression:
			f(t.Operand)
		case *ast.VariableExpression:
			f(t.Initializer)
		case *ast.StringLiteral:
			fmt.Println("*** yeah, I got a string: ", t.Value)

		default:
			fmt.Printf("uninteresting type %T\n", t) // %T prints whatever type t has
		}

	}

	for _, stmt := range program.Body {
		f(stmt)
	}

	spew.Dump(program)

}
