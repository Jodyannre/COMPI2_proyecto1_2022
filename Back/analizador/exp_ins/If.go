package exp_ins

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type IF struct {
	Condicion     Ast.Expresion
	Instrucciones *arraylist.List
	Lista_if_else *arraylist.List
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
	Expresion     bool
}

func NewIF(condicion Ast.Expresion, instrucciones *arraylist.List, lista *arraylist.List, tipo Ast.TipoDato,
	fila, columna int, expresion bool) IF {
	nif := IF{
		Condicion:     condicion,
		Instrucciones: instrucciones,
		Lista_if_else: lista,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
		Expresion:     expresion,
	}
	return nif
}

func (i IF) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, i.Tipo
}

func (i IF) Run(scope *Ast.Scope) interface{} {
	//Crear el nuevo scope
	newScope := Ast.NewScope("if", scope)
	//Inicializar la lista de respuestas
	//Ejecutar la instrucción if
	resultado := GetResultado(i, &newScope, -1, i.Expresion)

	//actualizar el scope global con los resultados
	newScope.UpdateScopeGlobal()
	return resultado
}

func GetResultado(i IF, scope *Ast.Scope, pos int, expresion bool) Ast.TipoRetornado {
	var condicion1 Ast.TipoRetornado
	if pos == -1 {
		condicion1 = i.Condicion.GetValue(scope)
	} else {
		//Conseguir el if/entonces
		//elemento := i.Lista_if_else.GetValue(pos).(IF)
		//Evaluar si es if o entonces
		if i.Tipo == Ast.ELSEIF || i.Tipo == Ast.ELSEIF_EXPRESION {
			//Es un if
			//i = elemento
			scope.Nombre = "ELSEIF"
			condicion1 = i.Condicion.GetValue(scope)
		} else if i.Tipo == Ast.ELSE || i.Tipo == Ast.ELSE_EXPRESION {
			//Es un else
			//i = elemento
			scope.Nombre = "ELSE"
			condicion1 = Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: true,
			}
		}
	}

	if condicion1.Tipo == Ast.BOOLEAN {
		if condicion1.Valor.(bool) {
			//Es verdadera; ejecutar las instrucciones
			n := 0
			var ultimaExpresion interface{}
			//Dar un valor por defecto a la expresión, por si nunca retorna nada
			ultimaExpresion = Ast.TipoRetornado{
				Valor: nil,
				Tipo:  Ast.NULL,
			}
			for n < i.Instrucciones.Len() {

				//Verificar que la instrucción no sea null
				if i.Instrucciones.GetValue(n) == nil {
					n++
					continue
				}

				elemento2 := i.Instrucciones.GetValue(n).(Ast.Abstracto)
				tipo_abstracto, _ := elemento2.GetTipo()
				if tipo_abstracto == Ast.EXPRESION && (i.Tipo == Ast.IF_EXPRESION ||
					i.Tipo == Ast.ELSE_EXPRESION || i.Tipo == Ast.ELSEIF_EXPRESION) {
					//Si es un if expresión, tiene que retornar algo

					//Verificar que sea la última o error
					if n != i.Instrucciones.Len()-1 {
						//Error porque no puede venir una expresión sin ser la última
						msg := "Semantic error, an instruction was expected." +
							" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
						nError := errores.NewError(i.Fila, i.Columna, msg)
						nError.Tipo = Ast.ERROR_SEMANTICO
						scope.Errores.Add(nError)
						scope.Consola += msg + "\n"
						return Ast.TipoRetornado{
							Valor: nError,
							Tipo:  Ast.ERROR,
						}
					}
					expresion := i.Instrucciones.GetValue(n).(Ast.Expresion)
					ultimaExpresion = expresion.GetValue(scope)
					if Ast.EsPrimitivo(ultimaExpresion.(Ast.TipoRetornado).Tipo) {
						scope.UpdateScopeGlobal()
						return ultimaExpresion.(Ast.TipoRetornado)
					}
				} else if tipo_abstracto == Ast.INSTRUCCION {
					instruccion := i.Instrucciones.GetValue(n).(Ast.Instruccion)
					//Es transferencia pero if es normal, solo retornar
					/*
						if Ast.EsTransferencia(tipoParticular) &&
							(i.Tipo == Ast.IF_EXPRESION ||
								i.Tipo == Ast.ELSE_EXPRESION || i.Tipo == Ast.ELSEIF_EXPRESION) {
							msg := "Semantic error," + Ast.ValorTipoDato[tipoParticular] + " statement not allowed inside this kind of IF." +
								" -- Line:" + strconv.Itoa(instruccion.(Ast.Abstracto).GetFila()) + " Column: " +
								strconv.Itoa(instruccion.(Ast.Abstracto).GetColumna())
							nError := errores.NewError(instruccion.(Ast.Abstracto).GetFila(),
								instruccion.(Ast.Abstracto).GetColumna(), msg)
							nError.Tipo = Ast.ERROR_SEMANTICO
							scope.Errores.Add(nError)
							scope.Consola += msg + "\n"
							return Ast.TipoRetornado{
								Valor: nError,
								Tipo:  Ast.ERROR,
							}
						}
					*/
					resultado := instruccion.Run(scope)
					/*
						if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR ||
							resultado.(Ast.TipoRetornado).Tipo == Ast.EJECUTADO {
							//Continuar a la siguiente instruccion
							//fmt.Println(resultado)
						}
					*/

					if Ast.EsTransferencia(resultado.(Ast.TipoRetornado).Tipo) {
						//Si es transferencia, terminar con el if y retornarlo
						return resultado.(Ast.TipoRetornado)
					}

				} else if tipo_abstracto == Ast.EXPRESION {
					msg := "Semantic error, an instruction was expected." +
						" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
					nError := errores.NewError(i.Fila, i.Columna, msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					scope.Errores.Add(nError)
					scope.Consola += msg + "\n"
					return Ast.TipoRetornado{
						Valor: nError,
						Tipo:  Ast.ERROR,
					}
				}
				n++
			}
			//Termino el for, retornar la ultima expresion
			//Verificar si hay algun retorno o retornar un error
			if ultimaExpresion.(Ast.TipoRetornado).Tipo == Ast.NULL && expresion {
				msg := "Semantic error, the if clause is not returning any value." +
					" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
				nError := errores.NewError(i.Fila, i.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR,
				}

			} else if expresion {
				//Si esta retornado algún valor
				return ultimaExpresion.(Ast.TipoRetornado)
			}

		} else {
			//Es falsa, buscar en la lista si hay otras
			//Llamada recursiva o fin
			//Recorrer la lista de ifs y else
			for j := 0; j < i.Lista_if_else.Len(); j++ {
				newScope := Ast.NewScope("if", scope)
				resultado := GetResultado(i.Lista_if_else.GetValue(j).(IF), &newScope, 0, expresion)
				if resultado.Tipo == Ast.EJECUTADO && resultado.Valor == true && !expresion {
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: true,
						Tipo:  Ast.EJECUTADO,
					}
				}
				if resultado.Tipo == Ast.ERROR {
					//newScope.Errores.Add(resultado.Valor)
					//newScope.Consola += resultado.Valor.(errores.CustomSyntaxError).Msg + "\n"
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: resultado.Valor,
						Tipo:  Ast.ERROR,
					}
				}

				if Ast.EsTransferencia(resultado.Tipo) {
					newScope.UpdateScopeGlobal()
					return resultado
				}

				if Ast.EsPrimitivo(resultado.Tipo) && (i.Tipo == Ast.IF_EXPRESION ||
					i.Tipo == Ast.ELSE_EXPRESION || i.Tipo == Ast.ELSEIF_EXPRESION) {
					scope.UpdateScopeGlobal()
					return resultado
				}

			}
			return Ast.TipoRetornado{
				Valor: false,
				Tipo:  Ast.EJECUTADO,
			}
		}
	} else {
		//No es booleano, entonces generar un error semántico
		//fmt.Println("Error semántico, la expresión no es un booleano")
		msg := "Semantic error, the condition of the expression is not a boolean expression." +
			" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
		nError := errores.NewError(i.Fila, i.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Se acabo todo y retornar un true de finalizado
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.EJECUTADO,
	}
}

func (op IF) GetFila() int {
	return op.Fila
}
func (op IF) GetColumna() int {
	return op.Columna
}
