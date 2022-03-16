package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"Back/analizador/fn_array"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type DeclaracionArray struct {
	Id        string
	Tipo      Ast.TipoDato
	TipoArray Ast.TipoRetornado
	Dimension interface{}
	Mutable   bool
	Publico   bool
	Valor     interface{}
	Fila      int
	Columna   int
}

func NewDeclaracionArray(id string, dimension interface{},
	mutable, publico bool, valor interface{}, fila int, columna int) DeclaracionArray {
	nd := DeclaracionArray{
		Id:        id,
		Tipo:      Ast.DECLARACION_ARRAY,
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
	var validacionDimensiones string
	existe := scope.Exist_actual(d.Id)
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()
	valor := d.Valor.(Ast.Expresion).GetValue(scope)
	dimension := d.Dimension.(Ast.Expresion).GetValue(scope)

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
	//Verificar error en dimension
	if dimension.Tipo == Ast.ERROR {
		return dimension
	}
	//Recuperar el tipo del array que se espera desde dimension
	//d.TipoArray = d.Dimension.(expresiones.DimensionArray).TipoArray

	//Primero que vengan arrays
	if !EsArray(tipoIn) {
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
	validacionDimensiones = fn_array.ConcordanciaArray(valor.Valor.(expresiones.Array))
	//Conseguir la lista
	split := strings.Split(validacionDimensiones, ",")
	//Crear la lista con las posiciones
	listaDimensiones := arraylist.New()
	for _, num := range split {
		listaDimensiones.Add(num)
	}

	//Comparar las lista de dimensiones
	//Get primitivos del array de dimension
	arrayDimension := arraylist.New()
	for i := 0; i < dimension.Valor.(*arraylist.List).Len(); i++ {
		arrayDimension.Add(dimension.Valor.(*arraylist.List).GetValue(i).(Ast.TipoRetornado).Valor)
	}

	if !fn_array.CompararListas(listaDimensiones, arrayDimension) {
		msg := "Semantic error, ARRAY dimension does not match" +
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

	//Validar el tipo del array
	if d.TipoArray.Tipo != valor.Valor.(expresiones.Array).TipoDelArray.Tipo {
		fila := valor.Valor.(expresiones.Array).GetFila()
		columna := valor.Valor.(expresiones.Array).GetColumna()
		var tipoDelArray Ast.TipoDato
		if valor.Valor.(expresiones.Array).TipoDelArray.Tipo == Ast.INDEFINIDO {
			tipoDelArray = valor.Valor.(expresiones.Array).TipoArray
		} else {
			tipoDelArray = valor.Valor.(expresiones.Array).TipoDelArray.Tipo
		}
		msg := "Semantic error, can't initialize ARRAY[" + Ast.ValorTipoDato[tipoDelArray] +
			"] with a ARRAY[" + Ast.ValorTipoDato[d.TipoArray.Tipo] + "]" +
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

	//Crear el símbolo
	nSimbolo := Ast.Simbolo{
		Identificador: d.Id,
		Valor:         valor,
		Fila:          d.Fila,
		Columna:       d.Columna,
		Tipo:          valor.Tipo,
		Mutable:       d.Mutable,
		Publico:       d.Publico,
	}
	scope.Add(nSimbolo)
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (op DeclaracionArray) GetFila() int {
	return op.Fila
}
func (op DeclaracionArray) GetColumna() int {
	return op.Columna
}

func DimensionesCorrectas(dimensiones arraylist.List, array interface{}, scope *Ast.Scope) Ast.TipoRetornado {
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
		if dimensiones.Len() == 0 {
			return Ast.TipoRetornado{
				Valor: true,
				Tipo:  Ast.BOOLEAN,
			}
		}
		validacion := DimensionesCorrectas(dimensiones, arreglo.Elementos.GetValue(i).(Ast.TipoRetornado).Valor, scope)
		if validacion.Tipo == Ast.ERROR {
			return validacion
		}
	}
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.BOOLEAN,
	}

}

func EsArray(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.ARRAY, Ast.ARRAY_ELEMENTOS, Ast.ARRAY_FAC:
		return true
	default:
		return false
	}
}
