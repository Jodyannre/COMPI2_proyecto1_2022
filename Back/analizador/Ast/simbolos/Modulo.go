package simbolos

import (
	"Back/analizador/Ast"

	"github.com/colegno/arraylist"
)

type Modulo struct {
	Identificador interface{}
	Tipo          Ast.TipoDato
	Instrucciones *arraylist.List
	Fila          int
	Columna       int
	Publico       bool
	Entorno       *Ast.Scope
}
