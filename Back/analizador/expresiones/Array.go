package expresiones

import (
	"Back/analizador/Ast"

	"github.com/colegno/arraylist"
)

type Array struct {
	Tipo          Ast.TipoDato
	TipoArray     Ast.TipoDato
	Elementos     *arraylist.List
	Fila          int
	Columna       int
	Mutable       bool
	TipoDelVector Ast.TipoDato
	TipoDelArray  Ast.TipoRetornado
	TipoDelStruct string
	Size          int //Tamaño de la dimensión
}

func NewArray(elementos *arraylist.List, TipoArray Ast.TipoDato, size, fila, columna int) Array {
	nA := Array{
		Tipo:          Ast.ARRAY,
		Elementos:     elementos,
		Fila:          fila,
		Columna:       columna,
		TipoArray:     TipoArray,
		TipoDelVector: Ast.INDEFINIDO,
		TipoDelArray:  Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true},
		TipoDelStruct: "INDEFINIDO",
		Size:          size,
	}
	return nA
}

func (a Array) GetValue(scope *Ast.Scope) Ast.TipoRetornado {

	return Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: a,
	}

}

func (a Array) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, a.Tipo
}

func (a Array) GetFila() int {
	return a.Fila
}
func (a Array) GetColumna() int {
	return a.Columna
}

func (a Array) GetElement(index int) Ast.TipoRetornado {
	return Ast.TipoRetornado{}
}

/*
func GetTipoVector(vector instrucciones.Vector) Ast.TipoDato {
	if vector.TipoVector == Ast.VECTOR {
		return vector.TipoDelVector
	}
	if vector.TipoVector == Ast.ARRAY {
		return vector.TipoDelArray
	}
	return vector.TipoVector
}

func GetTipoArray(array Array) Ast.TipoDato {
	if array.TipoArray == Ast.VECTOR {
		return array.TipoDelVector
	}
	if array.TipoArray == Ast.ARRAY {
		return array.TipoDelArray
	}
	return array.TipoArray
}
*/

func EsArray(tipo Ast.TipoDato) Ast.TipoDato {
	switch tipo {
	case Ast.ARRAY, Ast.ARRAY_ELEMENTOS, Ast.ARRAY_FAC:
		return Ast.ARRAY
	default:
		return tipo
	}
}
