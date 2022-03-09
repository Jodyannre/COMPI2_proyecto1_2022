package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/instrucciones"
	"strconv"

	"github.com/colegno/arraylist"
)

type Funcion struct {
	Nombre        string
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
	Instrucciones *arraylist.List
	Publica       bool
	Parametros    *arraylist.List
	Retorno       Ast.TipoDato
	Entorno       *Ast.Scope
}

func NewFuncion(nombre string, tipo Ast.TipoDato, instrucciones *arraylist.List,
	parametros *arraylist.List, retorno Ast.TipoDato, publica bool, fila, columna int) Funcion {
	nF := Funcion{
		Nombre:        nombre,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
		Instrucciones: instrucciones,
		Publica:       publica,
		Parametros:    parametros,
		Retorno:       retorno,
	}
	return nF
}

func (f Funcion) Run(scope *Ast.Scope) interface{} {
	var actual interface{}
	var tipoGeneral interface{}
	var respuesta interface{}
	//Ejecutar instrucciones

	for i := 0; i < f.Instrucciones.Len(); i++ {
		actual = f.Instrucciones.GetValue(i)
		if actual != nil {
			tipoGeneral, _ = actual.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}
		//Error, no puede haber expresiones sueltas
		if tipoGeneral == Ast.EXPRESION {
			fila := respuesta.(Ast.Abstracto).GetFila()
			columna := respuesta.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, an instruction was expected." +
				" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			continue
		}

		respuesta = actual.(Ast.Instruccion).Run(scope)
		if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
			//Hay error, pero ya esta agregado
			continue
		}

		if Ast.EsTransferencia(respuesta.(Ast.TipoRetornado).Tipo) {
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.BREAK ||
				respuesta.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION ||
				respuesta.(Ast.TipoRetornado).Tipo == Ast.CONTINUE {
				//Error de break, break_expresion
				valor := actual.(Ast.Abstracto)
				fila := valor.GetFila()
				columna := valor.GetColumna()
				msg := "Semantic error," + Ast.ValorTipoDato[respuesta.(Ast.TipoRetornado).Tipo] +
					" statement not allowed outside a loop." +
					" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
			}
			if f.Retorno == Ast.VOID && respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
				//Error de break, break_expresion
				valor := actual.(Ast.Abstracto)
				fila := valor.GetFila()
				columna := valor.GetColumna()
				msg := "Semantic error, this function can't return a value." +
					" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
			}
			if f.Retorno != Ast.VOID && respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN {
				//Error de break, break_expresion
				valor := actual.(Ast.Abstracto)
				fila := valor.GetFila()
				columna := valor.GetColumna()
				msg := "Semantic error, this function is not returning any value." +
					" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
			}

			if f.Retorno != Ast.VOID && respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
				//Verificar que los tipos sean correctos
				if f.Retorno != respuesta.(Ast.TipoRetornado).Valor.(Ast.TipoRetornado).Tipo {
					//Error de break, break_expresion
					valor := actual.(Ast.Abstracto)
					fila := valor.GetFila()
					columna := valor.GetColumna()
					msg := "Semantic error, this function is not returning any value." +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
					nError := errores.NewError(fila, columna, msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					scope.Errores.Add(nError)
					scope.Consola += msg + "\n"
				}
				return respuesta.(Ast.TipoRetornado).Valor.(Ast.TipoRetornado).Valor
			}
		}
	}

	//Verificar que la funcion no sea de return y no este retornando nada
	if f.Retorno != Ast.VOID {
		//Error la funcion debe retornar algo
		valor := actual.(Ast.Abstracto)
		fila := valor.GetFila()
		columna := valor.GetColumna()
		msg := "Semantic error, expected " + Ast.ValorTipoDato[f.Retorno] + " , found ()." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (f Funcion) RunParametros(scope *Ast.Scope, parametrosIN *arraylist.List) Ast.TipoRetornado {
	var tipos Ast.TipoRetornado             //Variable para verificar que los tipos son correctos
	var parametrosCreados Ast.TipoRetornado //Variaable para verificar que los parametros fueron creados
	// Primero revisar que la cantidad de parámetros sea la misma
	if parametrosIN.Len() != f.Parametros.Len() {
		//Error, la cantidad de parámetros no es la esperada
		msg := "Semantic error, wrong number of parameters in function." +
			" type. -- Line: " + strconv.Itoa(f.Fila) +
			" Column: " + strconv.Itoa(f.Columna)
		nError := errores.NewError(f.Fila, f.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	// Revisar que los tipos sean correctos
	tipos = TiposCorrectos(scope, f.Parametros, parametrosIN)
	if tipos.Tipo != Ast.BOOLEAN {
		return tipos
	}

	// crear los parámetros
	parametrosCreados = CrearParametros(scope, f.Parametros, parametrosIN)

	if parametrosCreados.Tipo == Ast.ERROR {
		return parametrosCreados
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: true,
	}
}

func TiposCorrectos(scope *Ast.Scope, parametros, parametrosIN *arraylist.List) Ast.TipoRetornado {
	var iterador int
	var resultadoParametroIN, resultadoParametro Ast.TipoRetornado
	var parametroIN, parametro interface{}
	for iterador = 0; iterador < parametros.Len(); iterador++ {
		parametro = parametros.GetValue(iterador)
		parametroIN = parametrosIN.GetValue(iterador)
		resultadoParametroIN = parametroIN.(Ast.Expresion).GetValue(scope)
		resultadoParametro = parametro.(Ast.Expresion).GetValue(scope)
		if resultadoParametroIN.Tipo != resultadoParametro.Tipo {
			//Error no son iguales los tipos
			fila := parametroIN.(Ast.Abstracto).GetFila()
			columna := parametroIN.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, " + Ast.ValorTipoDato[resultadoParametro.Tipo] +
				" type expected, " + Ast.ValorTipoDato[resultadoParametroIN.Tipo] + " type found." +
				" type. -- Line: " + strconv.Itoa(fila) +
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
	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: true,
	}
}

func CrearParametros(scope *Ast.Scope, parametros, parametrosIN *arraylist.List) Ast.TipoRetornado {
	var iterador int
	var resultadoParametro Ast.TipoRetornado
	var parametroIN, parametro, resultadoDeclaracion interface{}
	for iterador = 0; iterador < parametros.Len(); iterador++ {
		parametro = parametros.GetValue(iterador)
		parametroIN = parametrosIN.GetValue(iterador)
		resultadoParametro = parametro.(Ast.Expresion).GetValue(scope)
		//Crear un objeto declaración
		nuevaDeclaracion := instrucciones.NewDeclaracion(resultadoParametro.Valor.(string),
			resultadoParametro.Tipo, parametro.(Parametro).Mutable, false, Ast.VOID, parametroIN,
			parametro.(Ast.Abstracto).GetFila(), parametro.(Ast.Abstracto).GetColumna())
		//Ejecutar declaración
		resultadoDeclaracion = nuevaDeclaracion.Run(scope)

		if resultadoDeclaracion.(Ast.TipoRetornado).Tipo == Ast.ERROR {
			return resultadoDeclaracion.(Ast.TipoRetornado)
		}
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: true,
	}
}

func (f Funcion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, f.Tipo
}

func (f Funcion) GetFila() int {
	return f.Fila
}
func (f Funcion) GetColumna() int {
	return f.Columna
}