package errores

import (
	"Back/analizador/Ast"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CustomSyntaxError struct {
	Linea   int
	Columna int
	Msg     string
	Tipo    Ast.TipoDato
	Ambito  string
}

type CustomError struct {
	Linea   int
	Columna int
	Msg     string
	Tipo    Ast.TipoDato
	Ambito  string
}

func NewError(linea int, columna int, msg string) CustomSyntaxError {
	return CustomSyntaxError{
		Linea:   linea,
		Columna: columna,
		Msg:     msg,
	}
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errores []CustomSyntaxError
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errores = append(c.Errores, CustomSyntaxError{
		Linea:   line,
		Columna: column,
		Msg:     msg,
	})
}
