package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type LlamadaFuncion struct {
	Identificador Ast.Expresion
	Parametros    *arraylist.List
	Tipo          Ast.TipoDato
	Fila          int
	ScopeOriginal *Ast.Scope
	Columna       int
}

func NewLlamadaFuncion(id Ast.Expresion, parametros *arraylist.List, tipo Ast.TipoDato, fila, columna int) LlamadaFuncion {
	nF := LlamadaFuncion{
		Identificador: id,
		Parametros:    parametros,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nF
}

func (l LlamadaFuncion) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var funcion Funcion
	var parametrosCreados Ast.TipoRetornado
	var resultadoFuncion Ast.TipoRetornado
	//Crear el scope para la nueva función
	newScope := Ast.NewScope("funcion", scope)

	//Verificar que la función existe en el ámbilo global
	simbolo = newScope.Exist_fms_local(l.Identificador.(expresiones.Identificador).Valor)
	//Sino verificar que exista en el local

	if simbolo.Tipo != Ast.FUNCION ||
		simbolo.Tipo == Ast.ERROR_ACCESO_PRIVADO ||
		simbolo.Tipo == Ast.ERROR_NO_EXISTE {
		//La función no existe en el scope global, puede que exista en un módulo
		//simbolo = newScope.Exist_fms(l.Identificador.(Identificador).Valor)

		//Verificar si el símbolo es privado
		if simbolo.Tipo == Ast.ERROR_ACCESO_PRIVADO {
			//Error el símbolo tiene acceso privado
			msg := "Semantic error, the function is private." +
				" -- Line: " + strconv.Itoa(l.Fila) +
				" Column: " + strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}

		//Verificar si el símbolo existe
		if simbolo.Tipo == Ast.ERROR_NO_EXISTE {
			//Error el símbolo no existe
			msg := "Semantic error, the function doesn't exist." +
				" -- Line: " + strconv.Itoa(l.Fila) +
				" Column: " + strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

		//Verificar que el símbolo sea una función
		if simbolo.Tipo != Ast.FUNCION {
			//Error, el símbolo no es una función
			msg := "Semantic error, " + l.Identificador.(expresiones.Identificador).Valor + " is not a function." +
				" -- Line: " + strconv.Itoa(l.Fila) +
				" Column: " + strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
	} else {
		//simbolo = newScope.GetSimbolo(l.Identificador.(Identificador).Valor)
		funcion = simbolo.Valor.(Ast.TipoRetornado).Valor.(Funcion)
	}

	//Verificar que la función reciba o no parámetros y se estén enviando parámetros

	if l.Parametros.Len() > 0 && funcion.Parametros.Len() == 0 {
		//Error, se estan enviando parámetros y la función no pide parámetros
		msg := "Semantic error, " + l.Identificador.(expresiones.Identificador).Valor + " function doesn't expect parameters." +
			" -- Line: " + strconv.Itoa(l.Fila) +
			" Column: " + strconv.Itoa(l.Columna)
		nError := errores.NewError(l.Fila, l.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		newScope.Errores.Add(nError)
		newScope.Consola += msg + "\n"
		newScope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	if l.Parametros.Len() == 0 && funcion.Parametros.Len() > 0 {
		//Error, la función espera parámetros y no se están enviando parámetros
		msg := "Semantic error, " + l.Identificador.(expresiones.Identificador).Valor + " function expects parameters." +
			" -- Line: " + strconv.Itoa(l.Fila) +
			" Column: " + strconv.Itoa(l.Columna)
		nError := errores.NewError(l.Fila, l.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		newScope.Errores.Add(nError)
		newScope.Consola += msg + "\n"
		newScope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar que el scopeOriginal no sea null
	if l.ScopeOriginal == nil {
		l.ScopeOriginal = &newScope
	}

	//Crear los parámetros de las funciones
	parametrosCreados = funcion.RunParametros(&newScope, l.ScopeOriginal, l.Parametros)

	if parametrosCreados.Tipo == Ast.ERROR {
		newScope.UpdateScopeGlobal()
		return parametrosCreados
	}

	//Ejecutar la función
	resultadoFuncion = funcion.Run(&newScope).(Ast.TipoRetornado)
	newScope.UpdateScopeGlobal()
	/*
		if newScope.Errores.Len() > 0 {
			msg := "Semantic error, " + l.Identificador.(expresiones.Identificador).Valor + " function expects parameters." +
				" -- Line: " + strconv.Itoa(l.Fila) +
				" Column: " + strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}*/
	return resultadoFuncion
	/*
		return Ast.TipoRetornado{
			Tipo:  Ast.EJECUTADO,
			Valor: true,
		}*/
}

func (l LlamadaFuncion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, l.Tipo
}

func (l LlamadaFuncion) GetFila() int {
	return l.Fila
}
func (l LlamadaFuncion) GetColumna() int {
	return l.Columna
}
