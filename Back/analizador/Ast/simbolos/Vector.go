package simbolos

import (
	"Back/analizador/Ast"

	"github.com/colegno/arraylist"
)

type Vector struct {
	Tipo       Ast.TipoDato
	Valor      *arraylist.List
	TipoVector Ast.TipoDato
	Fila       int
	Columna    int
	Mutable    bool
}

func NewVector(tipo Ast.TipoDato, valor *arraylist.List, tipoVector Ast.TipoDato, mutable bool, fila, columna int) Vector {
	nV := Vector{
		Tipo:       tipo,
		Fila:       fila,
		Columna:    columna,
		Valor:      valor,
		TipoVector: tipoVector,
		Mutable:    mutable,
	}
	return nV
}

func (v Vector) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	return Ast.TipoRetornado{
		Valor: v,
		Tipo:  v.Tipo,
	}
}

func (v Vector) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v Vector) GetFila() int {
	return v.Fila
}
func (v Vector) GetColumna() int {
	return v.Columna
}

func (v Vector) GetTipoVector() Ast.TipoDato {
	return v.TipoVector
}
