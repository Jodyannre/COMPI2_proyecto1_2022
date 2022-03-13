package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type DeclaracionArray struct {
	Id        string
	Tipo      Ast.TipoDato
	TipoArray Ast.TipoDato
	Dimension expresiones.DimensionArray
	Mutable   bool
	Publico   bool
	Valor     interface{}
	Fila      int
	Columna   int
}

func NewDeclaracionArray(id string, dimension expresiones.DimensionArray,
	mutable, publico bool, tipoArray Ast.TipoDato,
	valor interface{}, fila int, columna int) DeclaracionArray {
	nd := DeclaracionArray{
		Id:        id,
		Tipo:      Ast.DECLARACION_ARRAY,
		TipoArray: tipoArray,
		Publico:   publico,
		Mutable:   mutable,
		Valor:     valor,
		Fila:      fila,
		Columna:   columna,
		Dimension: dimension,
	}
	return nd
}

func (d DeclaracionArray) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d DeclaracionArray) Run(scope *Ast.Scope) interface{} {
	//Verificar que exista, recuperar los arryas y los tipos
	var validacionDimensiones Ast.TipoRetornado
	existe := scope.Exist_actual(d.Id)
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()
	valor := d.Valor.(Ast.Expresion).GetValue(scope)
	dimension := d.Dimension.GetValue(scope)

	if existe {
		msg := "Semantic error, the element \"" + d.Id + "\" already exist in this scope." +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Verificar que no venga ningún error
	if valor.Tipo == Ast.ERROR {
		return valor
	}
	if dimension.Tipo == Ast.ERROR {
		return dimension
	}
	//Primero que vengan arrays
	if tipoIn != Ast.ARRAY {
		//Error, no se estan asignado arrays al array
		msg := "Semantic error, can't initialize an ARRAY with " + Ast.ValorTipoDato[tipoIn] + " type" +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar que las dimensiones concuerda con la lista de arrays
	validacionDimensiones = dimensionesCorrectas(dimension.Valor.(arraylist.List), valor, scope)
	return Ast.TipoRetornado{Valor: validacionDimensiones}
}

func (op DeclaracionArray) GetFila() int {
	return op.Fila
}
func (op DeclaracionArray) GetColumna() int {
	return op.Columna
}

func dimensionesCorrectas(dimensiones arraylist.List, array interface{}, scope *Ast.Scope) Ast.TipoRetornado {
	dimension := dimensiones.GetValue(0).(Ast.TipoRetornado)
	arreglo := array.(expresiones.Array)
	//Validar dimensiones
	if arreglo.Size != dimension.Valor.(int) {
		fila := arreglo.GetFila()
		columna := arreglo.GetColumna()
		msg := "Semantic error, Array dimensions don't match " +
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

	for i := 0; i < arreglo.Size; i++ {
		dimensiones.RemoveAtIndex(0)
		validacion := dimensionesCorrectas(dimensiones, arreglo.Elementos.GetValue(i), scope)
		if validacion.Tipo == Ast.ERROR {
			return validacion
		}
	}
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.BOOLEAN,
	}

}
