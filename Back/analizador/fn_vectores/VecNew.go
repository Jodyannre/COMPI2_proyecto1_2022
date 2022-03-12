package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/expresiones"

	"github.com/colegno/arraylist"
)

type VecNew struct {
	Tipo       Ast.TipoDato
	Valor      *arraylist.List
	TipoVector Ast.TipoDato
	Fila       int
	Columna    int
	Mutable    bool
	Factorial  bool
	Vacio      bool
	Tamaño     interface{}
	Referencia bool
	Size       int
}

func NewVecNew(tipo Ast.TipoDato, valor *arraylist.List, tipoVector Ast.TipoDato, tamaño interface{},
	mutable bool, factorial bool, vacio bool, fila, columna int) VecNew {
	//Crear el vector dependiendo de las banderas
	nV := VecNew{
		Tipo:       tipo,
		Fila:       fila,
		Columna:    columna,
		Valor:      valor,
		TipoVector: tipoVector,
		Mutable:    mutable,
		Factorial:  factorial,
		Vacio:      vacio,
		Tamaño:     tamaño,
		Referencia: false,
		Size:       0,
	}
	return nV
}

func (w VecNew) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	fila := w.Fila
	columna := w.Columna
	vacio := w.Vacio
	listaTemp := w.Valor
	nV := expresiones.NewVector(Ast.VECTOR, listaTemp, w.TipoVector,
		w.Tamaño, w.Mutable, w.Factorial, vacio, fila, columna)
	resultado := nV.GetValue(scope)
	return resultado
}

func (v VecNew) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v VecNew) GetFila() int {
	return v.Fila
}
func (v VecNew) GetColumna() int {
	return v.Columna
}
