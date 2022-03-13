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
	TipoDelArray  Ast.TipoDato
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
		TipoDelArray:  Ast.INDEFINIDO,
		Size:          size,
	}
	return nA
}

func (a Array) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Verificar todos los elementos de adentro y luego cambiar el tipo de array

	//Si el array ya ha sido declarado, solo retornarlo junto con su lista de elementos preparada

	//Sino esta declarada, ejecutar sus elementos y declararla

	//Crear la nueva lista de elementos
	/*
		elementos := arraylist.New()
		var elemento interface{}
		var tipoArray Ast.TipoDato
		valorElemento := Ast.TipoRetornado{}
		tipoAnterior := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
		tipoDelVectorAnterior := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
		tipoDelArrayAnterior := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
		var vector Vector
		var array Array
		size := 0
		vacio := true

		for i := 0; i < a.Elementos.Len(); i++ {
			elemento = a.Elementos.GetValue(i)
			//Calcular el valor del elemento
			valorElemento = elemento.(Ast.Expresion).GetValue(scope)
			//Si hay error solo lo retorno
			if valorElemento.Tipo == Ast.ERROR {
				return valorElemento
			}
			//Calcular los tipos del elemento
			tipoGeneral, tipoParticular := elemento.(Ast.Abstracto).GetTipo()
			if tipoParticular != Ast.IDENTIFICADOR {
				tipoAnterior.Tipo = tipoParticular
			} else {
				tipoAnterior.Tipo = valorElemento.Tipo
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
				vector = valorElemento.Valor.(Vector)
				if tipoDelVectorAnterior.Tipo == Ast.INDEFINIDO {
					tipoDelVectorAnterior.Tipo = GetTipoVector(vector)
				} else {
					if tipoDelVectorAnterior.Tipo != GetTipoVector(vector) {
						//Error no se pueden guardar 2 tipos de vectores diferentes
						fila := elemento.(Ast.Abstracto).GetFila()
						columna := elemento.(Ast.Abstracto).GetColumna()
						msg := "Semantic error, can't store VECTOR<" + Ast.ValorTipoDato[GetTipoVector(vector)] +
							"> to a VECTOR<VECTOR<" + Ast.ValorTipoDato[tipoDelVectorAnterior.Tipo] + ">>" +
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
				array = valorElemento.Valor.(Array)
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
			if vacio {
				vacio = false
				tipoArray = tipoParticular
			}
		}

		a.Elementos = nil
		a.Elementos = elementos
		a.Size = a.Elementos.Len()
		a.TipoArray = tipoArray
		a.TipoDelArray = tipoDelArrayAnterior.Tipo
		a.TipoDelVector = tipoDelVectorAnterior.Tipo
	*/
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

func GetTipoVector(vector Vector) Ast.TipoDato {
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
