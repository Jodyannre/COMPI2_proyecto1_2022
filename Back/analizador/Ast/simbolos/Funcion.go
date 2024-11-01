package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
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
	Retorno       Ast.TipoRetornado
	Entorno       *Ast.Scope
}

func NewFuncion(nombre string, tipo Ast.TipoDato, instrucciones *arraylist.List,
	parametros *arraylist.List, retorno Ast.TipoRetornado, publica bool, fila, columna int) Funcion {
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

		if tipoGeneral == Ast.INSTRUCCION {
			//Declarar variables globales
			respuesta = actual.(Ast.Instruccion).Run(scope)
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				return respuesta
			}
		} else if tipoGeneral == Ast.EXPRESION {
			respuesta = actual.(Ast.Expresion).GetValue(scope)
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				return respuesta
			}
		}
		//scope.UpdateReferencias()

		if Ast.EsTransferencia(respuesta.(Ast.TipoRetornado).Tipo) {
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.BREAK ||
				respuesta.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION ||
				respuesta.(Ast.TipoRetornado).Tipo == Ast.CONTINUE {
				//Error de break, break_expresion y continue
				valor := actual.(Ast.Abstracto)
				fila := valor.GetFila()
				columna := valor.GetColumna()
				msg := "Semantic error," + Ast.ValorTipoDato[respuesta.(Ast.TipoRetornado).Tipo] +
					" statement not allowed outside a loop." +
					" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			if f.Retorno.Tipo == Ast.VOID && respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
				//Error de break, break_expresion
				valor := actual.(Ast.Abstracto)
				fila := valor.GetFila()
				columna := valor.GetColumna()
				msg := "Semantic error, this function can't return a value." +
					" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			if f.Retorno.Tipo != Ast.VOID && respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN {
				//Error, la función espera retornar algo y no esta retornando nada
				valor := actual.(Ast.Abstracto)
				fila := valor.GetFila()
				columna := valor.GetColumna()
				msg := "Semantic error, this function is not returning any value." +
					" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}

			if f.Retorno.Tipo != Ast.VOID && respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
				//Verificar que los tipos sean correctos
				var tipoEntrante Ast.TipoRetornado
				var valorRespueta Ast.TipoRetornado = respuesta.(Ast.TipoRetornado).Valor.(Ast.TipoRetornado)
				if valorRespueta.Tipo == Ast.STRUCT {
					tipoEntrante.Tipo = Ast.STRUCT
					tipoEntrante.Valor = valorRespueta.Valor.(StructInstancia).GetPlantilla(scope)
				} else {
					if valorRespueta.Tipo == Ast.VECTOR {
						tipoEntrante = Ast.TipoRetornado{
							Tipo:  Ast.VECTOR,
							Valor: valorRespueta.Valor.(expresiones.Vector).TipoVector,
						}
					}
					if valorRespueta.Tipo == Ast.ARRAY {
						tipoEntrante = Ast.TipoRetornado{
							Tipo:  Ast.ARRAY,
							Valor: valorRespueta.Valor.(expresiones.Array).TipoDelArray,
						}
					}
				}
				if f.Retorno.Tipo == Ast.DIMENSION_ARRAY && tipoEntrante.Tipo == Ast.ARRAY {
					//Comparar las dimensiones solicitadas con el array de salida
					//Tengo que comparar el dimension de retorno y el valorRespuesta del array
					resultado := instrucciones.CompararDimensiones(f.Retorno.Valor.(expresiones.DimensionArray),
						valorRespueta.Valor.(expresiones.Array), scope)
					if resultado.Tipo == Ast.ERROR {
						return resultado
					}

				} else if !CompararTipos(f.Retorno, tipoEntrante) {
					//Error, retorna un tipo diferente
					valor := actual.(Ast.Abstracto)
					fila := valor.GetFila()
					columna := valor.GetColumna()
					msg := "Semantic error, expected" + expresiones.Tipo_String(f.Retorno) + " found " +
						expresiones.Tipo_String(respuesta.(Ast.TipoRetornado).Valor.(Ast.TipoRetornado)) +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
					nError := errores.NewError(fila, columna, msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					nError.Ambito = scope.GetTipoScope()
					scope.Errores.Add(nError)
					scope.Consola += msg + "\n"
					return Ast.TipoRetornado{
						Tipo:  Ast.ERROR,
						Valor: nError,
					}
				}
				//Ejecutar el return y retornar el valor que trae
				return respuesta.(Ast.TipoRetornado).Valor.(Ast.TipoRetornado)
			}
		}
	}

	//Verificar que la funcion no sea de return y no este retornando nada
	if f.Retorno.Tipo != Ast.VOID {
		//Error la funcion debe retornar algo
		valor := actual.(Ast.Abstracto)
		fila := valor.GetFila()
		columna := valor.GetColumna()
		msg := "Semantic error, expected " + expresiones.Tipo_String(f.Retorno) + " , found ()." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (f Funcion) RunParametros(scope *Ast.Scope, scopeOrigen *Ast.Scope, parametrosIN *arraylist.List) Ast.TipoRetornado {
	//var tipos Ast.TipoRetornado             //Variable para verificar que los tipos son correctos
	var parametrosCreados Ast.TipoRetornado //Variaable para verificar que los parametros fueron creados
	// Primero revisar que la cantidad de parámetros sea la misma
	if parametrosIN.Len() != f.Parametros.Len() {
		//Error, la cantidad de parámetros no es la esperada
		msg := "Semantic error, wrong number of parameters in function." +
			" type. -- Line: " + strconv.Itoa(f.Fila) +
			" Column: " + strconv.Itoa(f.Columna)
		nError := errores.NewError(f.Fila, f.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	// Revisar que los tipos sean correctos
	/*
		tipos = TiposCorrectos(scope, f.Parametros, parametrosIN)
		if tipos.Tipo != Ast.BOOLEAN {
			return tipos
		}
	*/

	// crear los parámetros
	parametrosCreados = CrearParametros(scope, scopeOrigen, f.Parametros, parametrosIN)

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

		//Verificar errores en los resultados
		if resultadoParametroIN.Tipo == Ast.ERROR {
			return resultadoParametroIN
		}
		if resultadoParametro.Tipo == Ast.ERROR {
			return resultadoParametro
		}

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
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

		//Verificar la mutabilidad en los valores que se mandan por valor
		_, tipoParticular := parametroIN.(Valor).Valor.(Ast.Abstracto).GetTipo()
		if tipoParticular == Ast.IDENTIFICADOR {

			//Verificar la mutabilidad en los valores que se mandan por referencia

			//Conseguir el símbolo del identificador
			simbolo := scope.GetSimbolo(parametroIN.(Valor).Valor.(expresiones.Identificador).Valor)
			//Primero verificar que el símbolo sea un struct, un vector o un array
			if simbolo.Tipo == Ast.STRUCT ||
				simbolo.Tipo == Ast.ARRAY ||
				simbolo.Tipo == Ast.VECTOR {
				if parametro.(Parametro).Mutable != parametroIN.(Valor).Mutable {
					fila := parametroIN.(Ast.Abstracto).GetFila()
					columna := parametroIN.(Ast.Abstracto).GetColumna()
					var mut1, mut2 string
					if parametro.(Parametro).Mutable {
						mut1 = "Mutable"
					} else {
						mut1 = "Not-Mutable"
					}
					if parametroIN.(Valor).Mutable {
						mut2 = "Mutable"
					} else {
						mut2 = "Not-Mutable"
					}
					msg := "Semantic error, " + mut1 + " value expected, " +
						mut2 + " value found." +
						" -- Line: " + strconv.Itoa(fila) +
						" Column: " + strconv.Itoa(columna)
					nError := errores.NewError(fila, columna, msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					nError.Ambito = scope.GetTipoScope()
					scope.Errores.Add(nError)
					scope.Consola += msg + "\n"
					return Ast.TipoRetornado{
						Tipo:  Ast.ERROR,
						Valor: nError,
					}
				}
				if simbolo.Mutable != parametroIN.(Valor).Mutable {
					fila := parametroIN.(Ast.Abstracto).GetFila()
					columna := parametroIN.(Ast.Abstracto).GetColumna()
					var mut1, mut2 string
					if simbolo.Mutable {
						mut1 = "Mutable"
					} else {
						mut1 = "Not-Mutable"
					}
					if parametroIN.(Valor).Mutable {
						mut2 = "Mutable"
					} else {
						mut2 = "Not-Mutable"
					}
					msg := "Semantic error, " + mut2 + " value expected, " +
						mut1 + " value found." +
						" -- Line: " + strconv.Itoa(fila) +
						" Column: " + strconv.Itoa(columna)
					nError := errores.NewError(fila, columna, msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					nError.Ambito = scope.GetTipoScope()
					scope.Errores.Add(nError)
					scope.Consola += msg + "\n"
					return Ast.TipoRetornado{
						Tipo:  Ast.ERROR,
						Valor: nError,
					}
				}

			} else if simbolo.Tipo == Ast.MODULO {
				fila := parametroIN.(Ast.Abstracto).GetFila()
				columna := parametroIN.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, a module can't be a parameter." +
					" -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			} else

			//Es cualquier otra variable
			if parametroIN.(Valor).Mutable {
				// Error, las variables no pueden ser mut
				fila := parametroIN.(Ast.Abstracto).GetFila()
				columna := parametroIN.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, " + Ast.ValorTipoDato[simbolo.Tipo] + " variable can´t be mut." +
					" -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}

		}

	}
	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: true,
	}
}

func CrearParametros(scope *Ast.Scope, scopeOrigen *Ast.Scope, parametros, parametrosIN *arraylist.List) Ast.TipoRetornado {
	var iterador int
	var resultadoParametro, tipoParametro, tipoParametroIN Ast.TipoRetornado
	var parametroIN, parametro, resultadoDeclaracion interface{}
	var paramTemp Parametro
	var verificacionDeReferencia Ast.TipoRetornado
	for iterador = 0; iterador < parametros.Len(); iterador++ {
		parametro = parametros.GetValue(iterador)
		tipoParametro = parametro.(Parametro).FormatearTipo(scope)
		paramTemp = parametro.(Parametro)
		paramTemp.TipoDeclaracion = tipoParametro
		parametro = paramTemp
		parametroIN = parametrosIN.GetValue(iterador)
		tipoParametroIN = parametroIN.(Ast.Expresion).GetValue(scopeOrigen) //**************
		resultadoParametro = parametro.(Ast.Expresion).GetValue(scope)
		if resultadoParametro.Tipo == Ast.ERROR {
			return resultadoParametro
		}
		if tipoParametroIN.Tipo == Ast.ERROR {
			return tipoParametroIN
		}
		//Crear un objeto declaración
		//Obtener el tipo del parámetro
		/*
			nuevaDeclaracion := instrucciones.NewDeclaracion(resultadoParametro.Valor.(string),
				resultadoParametro.Tipo, parametro.(Parametro).Mutable, false, Ast.VOID, parametroIN,
				parametro.(Ast.Abstracto).GetFila(), parametro.(Ast.Abstracto).GetColumna())
		*/
		//Verificar si es referencia y que todo este correcto
		verificacionDeReferencia = VerificarReferencia(parametro, parametroIN, scope, scopeOrigen)
		//Verificar algún posible error
		if verificacionDeReferencia.Tipo == Ast.ERROR {
			return verificacionDeReferencia
		}

		if parametroIN.(Valor).Referencia || !expresiones.EsVAS(tipoParametroIN.Tipo) {
			if tipoParametro.Tipo == Ast.DIMENSION_ARRAY {
				nuevaDeclaracion := instrucciones.NewDeclaracionArray(
					resultadoParametro.Valor.(string), tipoParametro.Valor, parametro.(Parametro).Mutable,
					false, tipoParametroIN.Valor, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)
			} else {
				nuevaDeclaracion := instrucciones.NewDeclaracionTotal(resultadoParametro.Valor.(string), parametroIN, tipoParametro,
					parametro.(Parametro).Mutable, false, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)

			}

		} else if expresiones.EsVAS(tipoParametroIN.Tipo) {

			if tipoParametro.Tipo == Ast.DIMENSION_ARRAY {
				nuevaDeclaracion := instrucciones.NewDeclaracionArrayNoRef(
					resultadoParametro.Valor.(string), tipoParametro.Valor, parametro.(Parametro).Mutable,
					false, tipoParametroIN.Valor, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)

			} else {
				nuevaDeclaracion := instrucciones.NewDeclaracionNoRef(resultadoParametro.Valor.(string), parametroIN, tipoParametro,
					parametro.(Parametro).Mutable, false, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)
			}

		}

		if resultadoDeclaracion.(Ast.TipoRetornado).Tipo == Ast.ERROR {
			return resultadoDeclaracion.(Ast.TipoRetornado)
		}
		//De forma sucia asignar la referencia si el valor es por referencia
		if parametroIN.(Valor).Referencia {
			simbolo := scope.GetSimbolo(resultadoParametro.Valor.(string))
			referencia := parametroIN.(Valor).Valor.(expresiones.Identificador).Valor
			simboloRef := scope.GetSimboloReferencia(referencia)
			simbolo.Referencia = true
			simbolo.Referencia_puntero = &simboloRef
			scope.UpdateSimbolo(resultadoParametro.Valor.(string), simbolo)
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

func (f Funcion) GetTipoRetornado(scope *Ast.Scope) string {
	tipo := GetTipoEstructura(f.Retorno, scope, f)
	cadena := Tipo_String(tipo)
	return cadena
}

func EsAVelementos(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.DIMENSION_ARRAY:
		return true
	default:
		return false
	}
}
