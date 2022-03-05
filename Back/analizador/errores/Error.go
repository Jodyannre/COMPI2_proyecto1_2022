package errores

import (
	"Back/analizador/Ast"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CustomSyntaxError struct {
	Fila    int
	Columna int
	Msg     string
	Tipo    Ast.TipoDato
	Ambito  string
}

type CustomError struct {
	Fila    int
	Columna int
	Msg     string
	Tipo    Ast.TipoDato
	Ambito  string
}

func NewError(fila int, columna int, msg string) CustomSyntaxError {
	return CustomSyntaxError{
		Fila:    fila,
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
		Fila:    line,
		Columna: column,
		Msg:     msg,
	})
}

func (op CustomSyntaxError) GetFila() int {
	return op.Fila
}
func (op CustomSyntaxError) GetColumna() int {
	return op.Columna
}
