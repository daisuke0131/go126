package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func evalBasicLit(bl *ast.BasicLit) float64  {
	switch bl.Kind {
	case token.INT:
		iv, err := strconv.ParseInt(bl.Value,10,64)
		if err != nil{
			panic(err)
		}
		return float64(iv)
	case token.FLOAT:
		fv, err := strconv.ParseFloat(bl.Value,64)
		if err != nil{
			panic(err)
		}
		return fv
	default:
		panic("error")
	}
}

func evalBinaryExpr(be *ast.BinaryExpr) float64{
	x := evalExpr(be.X)
	y := evalExpr(be.Y)

	switch be.Op {
	case token.ADD:
		return x + y
	case token.SUB:
		return x - y
	case token.MUL:
		return x * y
	case token.QUO:
		return x / y
	case token.REM:
		return float64(int64(x) % int64(y))
	default:
		panic("error")
	}
}


func evalUnaryExpr(ue *ast.UnaryExpr) float64{
	x := evalExpr(ue.X)

	switch ue.Op {
	case token.ADD:
		return x
	case token.SUB:
		return -x
	default:
		panic("error")
	}
}

func evalExpr(expr ast.Expr) float64{
	switch e := expr.(type) {
	case *ast.BasicLit:
		return evalBasicLit(e)
	case *ast.BinaryExpr:
		return evalBinaryExpr(e)
	case *ast.ParenExpr:
		return evalExpr(e.X)
	case *ast.UnaryExpr:
		return evalUnaryExpr(e)
	default:
		panic("error")
	}
}

func main(){
	val := "6 + 2 * 3 + 3"

	expr, err := parser.ParseExpr(val)
	if err != nil{
		panic("error")
	}
	v := evalExpr(expr)
	fmt.Println(v)
}