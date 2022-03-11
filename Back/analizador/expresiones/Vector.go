package expresiones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type Vector struct {
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

func NewVector(tipo Ast.TipoDato, valor *arraylist.List, tipoVector Ast.TipoDato, tamaño interface{}, mutable bool, factorial bool, vacio bool, fila, columna int) Vector {
	//Crear el vector dependiendo de las banderas
	nV := Vector{
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

func (v Vector) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Crear los valores del vector
	if v.Vacio {
		//Es un vector vacio, pero tiene que tener el tipo, entonces solo crear la lista y retornarlo
		v.Valor = arraylist.New()
		tamaño := v.Tamaño.(Ast.Expresion).GetValue(scope)
		_, tipoParticular := v.Tamaño.(Ast.Abstracto).GetTipo()

		if (tipoParticular == Ast.USIZE || tamaño.Tipo == Ast.USIZE) ||
			(tipoParticular == Ast.I64 && tipoParticular != Ast.IDENTIFICADOR) {
			//Verificar que sea uzise
			nElemento := Ast.TipoRetornado{
				Tipo:  Ast.LIBRE,
				Valor: true,
			}
			tamaño := v.Tamaño.(Ast.Expresion).GetValue(scope).Valor.(int)
			for i := 0; i < tamaño; i++ {
				elemento := nElemento
				v.Valor.Add(elemento)
				if v.Vacio {
					v.Vacio = false
				}
			}
			v.Size = tamaño

		} else if tipoParticular != Ast.LIBRE {
			if tamaño.Tipo == Ast.ERROR {
				return tamaño
			}
			//ERROR NO ES UN USIZE
			fila := v.Tamaño.(Ast.Abstracto).GetFila()
			columna := v.Tamaño.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, expected usize, found " + Ast.ValorTipoDato[tamaño.Tipo] +
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

		return Ast.TipoRetornado{
			Valor: v,
			Tipo:  v.Tipo,
		}
	} else if !v.Vacio && !v.Factorial {
		//Trae elementos
		//Copio elementos y reinicio el del vector para agregar los elementos ya ejecutados
		listaTemp := v.Valor.Clone()
		v.Valor = arraylist.New()
		valorElemento := Ast.TipoRetornado{}
		tipoAnterior := Ast.TipoRetornado{Tipo: Ast.LIBRE, Valor: true}
		for i := 0; i < listaTemp.Len(); i++ {
			elemento := listaTemp.GetValue(i)
			tipoGeneral, tipoParticular := elemento.(Ast.Abstracto).GetTipo()
			if tipoAnterior.Tipo != Ast.LIBRE {
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
			} else {
				tipoAnterior.Tipo = tipoParticular
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
			valorElemento = elemento.(Ast.Expresion).GetValue(scope)
			//Si hay error solo lo retorno
			if valorElemento.Tipo == Ast.ERROR {
				return valorElemento
			}
			//Todo bien, entonces agregarlo al vector, cambiar el estado de vacio y aumentar el tamaño
			v.Valor.Add(valorElemento)
			v.Size++
			v.TipoVector = tipoAnterior.Tipo
			if v.Vacio {
				v.Vacio = false
			}
		}

	} else {
		//Crear la cantidad de elementos que se solicita
		//conseguir la cantidad de veces que se va a repetir el valor
		veces := v.Valor.GetValue(1).(Ast.Expresion).GetValue(scope)
		elemento := v.Valor.GetValue(0).(Ast.Expresion).GetValue(scope)

		if veces.Tipo == Ast.ERROR {
			return veces
		}
		if elemento.Tipo == Ast.ERROR {
			return elemento
		}
		if veces.Tipo != Ast.USIZE {
			//Error, se esperaba un USIZE
			fila := v.Valor.GetValue(1).(Ast.Abstracto).GetFila()
			columna := v.Valor.GetValue(1).(Ast.Abstracto).GetColumna()
			msg := "Semantic error, expected usize, found. " + Ast.ValorTipoDato[veces.Tipo] +
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
		//Reiniciamos el valor del vector
		v.Valor = arraylist.New()
		for i := 0; i < veces.Valor.(int); i++ {
			nElemento := elemento
			v.Valor.Add(nElemento)
			v.Size++
		}
	}

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

func (v Vector) GetSize() int {
	return v.Size
}
