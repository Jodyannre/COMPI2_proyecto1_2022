package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type ContainsVec struct {
	Identificador interface{}
	Tipo          Ast.TipoDato
	Valor         interface{}
	Fila          int
	Columna       int
}

func NewContainsVec(id interface{}, valor interface{}, tipo Ast.TipoDato, fila, columna int) ContainsVec {
	nP := ContainsVec{
		Identificador: id,
		Tipo:          tipo,
		Valor:         valor,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p ContainsVec) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var vector expresiones.Vector
	var resultado = false
	var id string
	//Primero verificar que sea un identificador el id
	_, tipoParticular := p.Identificador.(Ast.Abstracto).GetTipo()
	if tipoParticular != Ast.IDENTIFICADOR {
		//Error se espera un identificador
		msg := "Semantic error, expected VECTOR, found " + Ast.ValorTipoDato[tipoParticular] +
			". -- Line: " + strconv.Itoa(p.Fila) +
			" Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Recuperar el id del identificador
	id = p.Identificador.(expresiones.Identificador).Valor

	//Verificar que el id exista
	if !scope.Exist(id) {
		//Error la variable no existe
		msg := "Semantic error, the element \"" + id + "\" doesn't exist in any scope." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Conseguir el simbolo y el vector
	simbolo = scope.GetSimbolo(id)
	//Verificar que sea un vector
	if simbolo.Tipo != Ast.VECTOR {
		msg := "Semantic error, expected Vector, found " + Ast.ValorTipoDato[simbolo.Tipo] + "." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)
	valor := p.Valor.(Ast.Expresion).GetValue(scope)
	//Verificar que sean del mismo tipo

	if vector.Vacio {
		resultado = false
	} else if valor.Tipo != vector.TipoVector {
		msg := "Semantic error, expected &" + Ast.ValorTipoDato[vector.TipoVector] +
			", found &" + Ast.ValorTipoDato[valor.Tipo] + "." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	} else {
		for i := 0; i < vector.Valor.Len(); i++ {
			//Primero verificar que sean del mismo tipo
			elemento := vector.Valor.GetValue(i).(Ast.TipoRetornado)
			if elemento.Valor == valor.Valor {
				resultado = true
				break
			}
		}
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: resultado,
	}
}

func (v ContainsVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v ContainsVec) GetFila() int {
	return v.Fila
}
func (v ContainsVec) GetColumna() int {
	return v.Columna
}
