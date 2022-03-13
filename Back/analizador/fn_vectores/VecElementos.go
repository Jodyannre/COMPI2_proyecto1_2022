package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type VecElementos struct {
	Elementos *arraylist.List
	Tipo      Ast.TipoDato
	Fila      int
	Columna   int
}

func NewVecElementos(elementos *arraylist.List, fila, columna int) VecElementos {
	//Crear el vector dependiendo de las banderas
	nV := VecElementos{
		Tipo:      Ast.VEC_ELEMENTOS,
		Fila:      fila,
		Columna:   columna,
		Elementos: elementos,
	}
	return nV
}

func (v VecElementos) GetValue(scope *Ast.Scope) Ast.TipoRetornado {

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
	capacity := 0
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
					" in a Vec<" + Ast.ValorTipoDato[tipoAnterior.Tipo] + ">." +
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
			msg := "Semantic error, can't store " + Ast.ValorTipoDato[tipoParticular] + " to a vector" +
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
					msg := "Semantic error, can't store VECTOR<" + Ast.ValorTipoDato[GetTipoVector(vector)] +
						"> to a VECTOR< VECTOR<" + Ast.ValorTipoDato[tipoDelVectorAnterior.Tipo] + "> >" +
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
						"] to a VECTOR<ARRAY[" + Ast.ValorTipoDato[tipoDelArrayAnterior.Tipo] + "]>" +
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
	capacity = size

	newVector := expresiones.NewVector(elementos, tipoVector, size, capacity, vacio, v.Fila, v.Columna)
	newVector.TipoDelArray = tipoDelArrayAnterior.Tipo
	newVector.TipoDelVector = tipoDelVectorAnterior.Tipo
	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: newVector,
	}

}

func (v VecElementos) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v VecElementos) GetFila() int {
	return v.Fila
}
func (v VecElementos) GetColumna() int {
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
		Ast.VEC_WITH_CAPACITY, Ast.VEC_ELEMENTOS:
		validador = true
	default:
		validador = false
	}

	return validador
}
