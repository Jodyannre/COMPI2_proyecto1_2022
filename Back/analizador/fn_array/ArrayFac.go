package fn_array

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type ArrayFactorial struct {
	Tipo      Ast.TipoDato
	Elementos *arraylist.List
	TipoArray Ast.TipoDato
	Fila      int
	Columna   int
}

func NewArrayFactorial(elementos *arraylist.List, fila, columna int) ArrayFactorial {
	//Crear el vector dependiendo de las banderas
	nV := ArrayFactorial{
		Tipo:      Ast.ARRAY_FAC,
		Fila:      fila,
		Columna:   columna,
		TipoArray: Ast.INDEFINIDO,
		Elementos: elementos,
	}
	return nV
}

func (v ArrayFactorial) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Se crea como factorial
	//Crear la cantidad de elementos que se solicita
	//conseguir la cantidad de veces que se va a repetir el valor
	elementoVeces := v.Elementos.GetValue(1)
	_, tipoParticular := elementoVeces.(Ast.Abstracto).GetTipo()
	cantidad := elementoVeces.(Ast.Expresion).GetValue(scope)
	elemento := v.Elementos.GetValue(0).(Ast.Expresion).GetValue(scope)
	sizeArray := 0
	tipoDelVector := Ast.INDEFINIDO
	tipoDelArray := Ast.INDEFINIDO
	vacio := true
	var tipoArray Ast.TipoDato

	if cantidad.Tipo == Ast.ERROR {
		return cantidad
	}
	if elemento.Tipo == Ast.ERROR {
		return elemento
	}
	if cantidad.Tipo != Ast.USIZE && tipoParticular != Ast.I64 {
		//Error, se esperaba un USIZE
		fila := v.Elementos.GetValue(1).(Ast.Abstracto).GetFila()
		columna := v.Elementos.GetValue(1).(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected usize, found. " + Ast.ValorTipoDato[cantidad.Tipo] +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Reiniciamos el valor del array
	elementos := arraylist.New()
	for i := 0; i < cantidad.Valor.(int); i++ {
		nElemento := elemento
		if nElemento.Tipo == Ast.ERROR {
			return nElemento
		}
		if nElemento.Tipo == Ast.ARRAY {
			tipoDelArray = expresiones.GetTipoArray(nElemento.Valor.(expresiones.Array))
		}
		if nElemento.Tipo == Ast.VECTOR {
			tipoDelVector = expresiones.GetTipoVector(nElemento.Valor.(expresiones.Vector))
		}
		elementos.Add(nElemento)
		sizeArray++
		if vacio {
			vacio = false
			tipoArray = elemento.Tipo
		}
	}

	array := expresiones.NewArray(elementos, tipoArray, sizeArray, v.Fila, v.Columna)
	array.TipoDelVector = tipoDelVector
	array.TipoDelArray = tipoDelArray

	return Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: array,
	}
}

func (v ArrayFactorial) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v ArrayFactorial) GetFila() int {
	return v.Fila
}
func (v ArrayFactorial) GetColumna() int {
	return v.Columna
}
