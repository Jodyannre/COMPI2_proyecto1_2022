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

func UpdatePosition(v *Vector, posicion int, valorEntrante interface{}, scope *Ast.Scope) Ast.TipoRetornado {
	valor := valorEntrante.(Ast.Expresion).GetValue(scope)
	listaNueva := *arraylist.New()
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	for i := 0; i < v.Valor.Len(); i++ {
		if i == posicion {
			listaNueva.Add(valor)
			continue
		}
		listaNueva.Add(v.Valor.GetValue(i))
	}
	v.Valor = &listaNueva
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func GetTipoFinal(tipo Ast.TipoRetornado) Ast.TipoRetornado {
	if EsTipoFinal(tipo.Tipo) {
		if tipo.Tipo != Ast.STRUCT {
			return Ast.TipoRetornado{
				Tipo:  tipo.Tipo,
				Valor: true,
			}
		}
		return Ast.TipoRetornado{
			Tipo:  tipo.Tipo,
			Valor: tipo.Valor,
		}
	} else {
		return GetTipoFinal(tipo.Valor.(Ast.TipoRetornado))
	}
}

func EsTipoFinal(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.I64, Ast.F64, Ast.CHAR, Ast.STRING, Ast.STR, Ast.USIZE, Ast.BOOLEAN, Ast.STRUCT, Ast.INDEFINIDO,
		Ast.DIMENSION_ARRAY:
		return true
	default:
		return false
	}
}

func CompararTipos(tipoA Ast.TipoRetornado, tipoB Ast.TipoRetornado) bool {
	if EsTipoFinal(tipoA.Tipo) && EsTipoFinal(tipoB.Tipo) {
		if tipoA.Tipo == tipoB.Tipo {
			//Verificar si son structs
			if tipoA.Tipo == Ast.STRUCT {
				if tipoA.Valor == tipoB.Valor {
					return true
				} else {
					return false
				}
			}
			return true
		} else {
			return false
		}
	}
	if EsTipoFinal(tipoA.Tipo) && !EsTipoFinal(tipoB.Tipo) ||
		!EsTipoFinal(tipoA.Tipo) && EsTipoFinal(tipoB.Tipo) {
		return false
	}
	return CompararTipos(tipoA.Valor.(Ast.TipoRetornado), tipoB.Valor.(Ast.TipoRetornado))
}

func Tipo_String(t Ast.TipoRetornado) string {
	if t.Tipo == Ast.VECTOR {
		return "Vec <" + Tipo_String(t.Valor.(Ast.TipoRetornado)) + ">"
	} else {
		if t.Tipo == Ast.STRUCT {
			return t.Valor.(string)
		} else {
			return Ast.ValorTipoDato[t.Tipo]
		}
	}
}

func EsVector(tipo Ast.TipoDato) Ast.TipoDato {
	switch tipo {
	case Ast.VEC_ELEMENTOS, Ast.VEC_FAC, Ast.VEC_WITH_CAPACITY, Ast.VEC_NEW:
		return Ast.VECTOR
	default:
		return tipo
	}
}
