package fn_array

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type ArrayElementos struct {
	Elementos *arraylist.List
	Tipo      Ast.TipoDato
	Fila      int
	Columna   int
}

func NewArrayElementos(elementos *arraylist.List, fila, columna int) ArrayElementos {
	//Crear el vector dependiendo de las banderas
	nV := ArrayElementos{
		Tipo:      Ast.ARRAY_ELEMENTOS,
		Fila:      fila,
		Columna:   columna,
		Elementos: elementos,
	}
	return nV
}

func (v ArrayElementos) GetValue(scope *Ast.Scope) Ast.TipoRetornado {

	elementos := arraylist.New()
	var elemento interface{}
	var tipoVector Ast.TipoDato
	valorElemento := Ast.TipoRetornado{}
	tipoAnterior := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
	tipoDelVectorAnterior := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
	tipoDelArrayAnterior := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
	var vector expresiones.Vector
	var array expresiones.Array
	size := 0
	vacio := true

	for i := 0; i < v.Elementos.Len(); i++ {
		elemento = v.Elementos.GetValue(i)
		//Calcular el valor del elemento
		valorElemento = elemento.(Ast.Expresion).GetValue(scope)
		//Si hay error solo lo retorno
		if valorElemento.Tipo == Ast.ERROR {
			return valorElemento
		}
		//Calcular los tipos del elemento
		tipoGeneral, tipoParticular := elemento.(Ast.Abstracto).GetTipo()
		if tipoParticular != Ast.IDENTIFICADOR && !EsFuncion(tipoParticular) {
			tipoAnterior.Tipo = tipoParticular
		} else {
			tipoParticular = valorElemento.Tipo
		}
		if tipoAnterior.Tipo != Ast.INDEFINIDO {
			if tipoAnterior.Tipo != tipoParticular {
				//Los tipos son diferentes, error
				fila := elemento.(Ast.Abstracto).GetFila()
				columna := elemento.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, can't store " + Ast.ValorTipoDato[tipoParticular] + " value" +
					" in a ARRAY[" + Ast.ValorTipoDato[tipoAnterior.Tipo] + "]." +
					" -- Line: " + strconv.Itoa(fila) +
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
		}
		if tipoGeneral != Ast.EXPRESION {
			//Error, no se puede guardar algo que no sea una expresión en el vector
			fila := elemento.(Ast.Abstracto).GetFila()
			columna := elemento.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, can't store " + Ast.ValorTipoDato[tipoParticular] + " to an ARRAY" +
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

		//Si es un array o un vector, verificar que los tipos sean correctos
		if tipoParticular == Ast.VECTOR {
			vector = valorElemento.Valor.(expresiones.Vector)
			if tipoDelVectorAnterior.Tipo == Ast.INDEFINIDO {
				tipoDelVectorAnterior.Tipo = GetTipoVector(vector)
			} else {
				if tipoDelVectorAnterior.Tipo != GetTipoVector(vector) {
					//Error no se pueden guardar 2 tipos de vectores diferentes
					fila := elemento.(Ast.Abstracto).GetFila()
					columna := elemento.(Ast.Abstracto).GetColumna()
					msg := "Semantic error, can't store ARRAY[" + Ast.ValorTipoDato[GetTipoVector(vector)] +
						"] to a ARRAY[ VECTOR<" + Ast.ValorTipoDato[tipoDelVectorAnterior.Tipo] + "> >" +
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
			}
		}

		//Si es un array, verificar que los arrays sean del mismo tipo
		if tipoParticular == Ast.ARRAY {
			array = valorElemento.Valor.(expresiones.Array)
			if tipoDelArrayAnterior.Tipo == Ast.INDEFINIDO {
				tipoDelArrayAnterior.Tipo = GetTipoArray(array)
			} else {
				if tipoDelArrayAnterior.Tipo != GetTipoArray(array) {
					//Error no se pueden guardar 2 tipos de vectores diferentes
					fila := elemento.(Ast.Abstracto).GetFila()
					columna := elemento.(Ast.Abstracto).GetColumna()
					msg := "Semantic error, can't store ARRAY[" + Ast.ValorTipoDato[GetTipoArray(array)] +
						"] to a ARRAY[" + Ast.ValorTipoDato[tipoDelArrayAnterior.Tipo] + "]" +
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
			}
		}

		//Todo bien, entonces agregar el elemento a la lista del vector
		elementos.Add(valorElemento)
		size++
		tipoAnterior.Tipo = valorElemento.Tipo
		if vacio {
			vacio = false
			tipoVector = tipoParticular
		}
	}
	//Actualizar capacity

	newArray := expresiones.NewArray(elementos, tipoVector, size, v.Fila, v.Columna)
	newArray.TipoDelArray = tipoDelArrayAnterior.Tipo
	newArray.TipoDelVector = tipoDelVectorAnterior.Tipo
	if newArray.TipoDelArray == Ast.INDEFINIDO {
		newArray.TipoDelArray = newArray.TipoArray
	}

	concordancia := ConcordanciaDimensiones(newArray)
	if concordancia.Tipo == Ast.ERROR {
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, ARRAY dimensions do not match." +
			" -- Line: " + strconv.Itoa(fila) +
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
	return Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: newArray,
	}

}

func (v ArrayElementos) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v ArrayElementos) GetFila() int {
	return v.Fila
}
func (v ArrayElementos) GetColumna() int {
	return v.Columna
}

func GetTipoVector(vector expresiones.Vector) Ast.TipoDato {
	if vector.TipoVector == Ast.VECTOR {
		return vector.TipoDelVector
	}
	if vector.TipoVector == Ast.ARRAY {
		return vector.TipoDelArray
	}
	return vector.TipoVector
}

func GetNivelesVector(vectorGuardado expresiones.Vector, vectorEntrante expresiones.Vector) bool {
	if vectorGuardado.TipoVector == Ast.VECTOR && vectorEntrante.TipoVector != Ast.VECTOR ||
		vectorGuardado.TipoVector != Ast.VECTOR && vectorEntrante.TipoVector == Ast.VECTOR {
		return false
	}
	if vectorGuardado.TipoVector == Ast.VECTOR && vectorEntrante.TipoVector == Ast.VECTOR {
		//Verificar si tiene elementos ambos
		if vectorGuardado.Valor.Len() > 0 && vectorEntrante.Valor.Len() > 0 {
			return GetNivelesVector(vectorGuardado.Valor.GetValue(0).(Ast.TipoRetornado).Valor.(expresiones.Vector),
				vectorEntrante.Valor.GetValue(0).(Ast.TipoRetornado).Valor.(expresiones.Vector))
		}
		return true
	}
	return true
}

func GetTipoArray(array expresiones.Array) Ast.TipoDato {
	if array.TipoArray == Ast.VECTOR {
		return array.TipoDelVector
	}
	if array.TipoArray == Ast.ARRAY {
		return array.TipoDelArray
	}
	return array.TipoArray
}

func EsFuncion(tipo interface{}) bool {
	validador := false

	switch tipo {
	case Ast.FUNCION, Ast.VEC_NEW,
		Ast.VEC_LEN, Ast.VEC_CONTAINS,
		Ast.VEC_CAPACITY, Ast.VEC_REMOVE, Ast.VEC_FAC,
		Ast.VEC_WITH_CAPACITY, Ast.VEC_ELEMENTOS, Ast.ARRAY_ELEMENTOS, Ast.ARRAY_FAC, Ast.LLAMADA_FUNCION:
		validador = true
	default:
		validador = false
	}

	return validador
}

func ConcordanciaDimensiones(arr interface{}) Ast.TipoRetornado {
	array := arr.(expresiones.Array)
	if array.TipoArray != Ast.ARRAY {
		lista := *arraylist.New()
		lista.Add(array.Size)
		return Ast.TipoRetornado{
			Tipo:  Ast.EJECUTADO,
			Valor: &lista,
		}
	}

	var resultadoAnterior Ast.TipoRetornado = Ast.TipoRetornado{Tipo: Ast.LIBRE, Valor: true}
	var resultadoActual Ast.TipoRetornado = Ast.TipoRetornado{Tipo: Ast.LIBRE, Valor: true}
	if array.Elementos.Len() > 0 {
		for i := 0; i < array.Elementos.Len(); i++ {
			resultadoActual = ConcordanciaDimensiones(array.Elementos.GetValue(i).(Ast.TipoRetornado).Valor)
			resultadoActual.Valor.(*arraylist.List).Add(array.Size)
			if resultadoActual.Tipo == Ast.ERROR {
				return resultadoActual
			}
			if resultadoAnterior.Tipo != Ast.LIBRE {
				if !CompararListas(resultadoActual.Valor.(*arraylist.List), resultadoAnterior.Valor.(*arraylist.List)) {
					return Ast.TipoRetornado{Tipo: Ast.ERROR, Valor: true}
				}
			}

			resultadoAnterior = resultadoActual
		}
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: resultadoAnterior.Valor,
	}
}

func CompararListas(listaActual *arraylist.List, listaAnterior *arraylist.List) bool {
	if listaActual.Len() != listaAnterior.Len() {
		return false
	}
	for i := 0; i < listaActual.Len(); i++ {
		elemento1 := listaActual.GetValue(i)
		elemento2 := listaAnterior.GetValue(i)

		if elemento1 != elemento2 {
			return false
		}
	}
	return true
}
