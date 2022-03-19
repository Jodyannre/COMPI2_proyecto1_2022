package expresiones

import (
	"Back/analizador/Ast"
	//"Back/analizador/instrucciones"

	"github.com/colegno/arraylist"
)

type Vector struct {
	Tipo       Ast.TipoDato
	Valor      *arraylist.List
	TipoVector Ast.TipoRetornado
	Fila       int
	Columna    int
	Mutable    bool
	Vacio      bool
	Size       int
	Capacity   int
}

func NewVector(valor *arraylist.List, tipoVector Ast.TipoRetornado, size, capacity int, vacio bool, fila, columna int) Vector {
	//Crear el vector dependiendo de las banderas
	nV := Vector{
		Tipo:       Ast.VECTOR,
		Fila:       fila,
		Columna:    columna,
		Valor:      valor,
		TipoVector: tipoVector,
		Mutable:    false,
		Size:       size,
		Capacity:   capacity,
		Vacio:      vacio,
	}
	return nV
}

func (v Vector) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Crear los valores del vector
	return Ast.TipoRetornado{
		Valor: v,
		Tipo:  Ast.VECTOR,
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

func (v Vector) GetTipoVector() Ast.TipoRetornado {
	return v.TipoVector
}

func (v Vector) GetSize() int {
	return v.Size
}

func (v Vector) GetMutable() bool {
	return v.Mutable
}

func (v Vector) CalcularCapacity(size int, capacity int) int {
	if size == 1 && capacity == 0 {
		return 4
	}
	if size == 0 && capacity == 0 {
		return 0
	}
	if capacity <= size {
		if capacity == 0 {
			return v.CalcularCapacity(size, capacity+4)
		}
		return v.CalcularCapacity(size, capacity*2)
	} else {
		return capacity
	}
}
