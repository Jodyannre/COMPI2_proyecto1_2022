package bucles

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type Loop struct {
	Tipo          Ast.TipoDato
	Instrucciones *arraylist.List
	Fila          int
	Columna       int
}

func (l Loop) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, l.Tipo
}

func NewLoop(tipo Ast.TipoDato, instrucciones *arraylist.List, fila, columna int) Loop {
	nL := Loop{
		Tipo:          tipo,
		Instrucciones: instrucciones,
		Fila:          fila,
		Columna:       columna,
	}
	return nL
}

func (l Loop) Run(scope *Ast.Scope) interface{} {
	newScope := Ast.NewScope("Loop", scope)
	var contadorSeguridad int = 0
	var retornarBreak bool = false
	var instruccion, resultado interface{}
	var tipoGeneral, tipoParticular Ast.TipoDato
	var i int = 0

	for {

		if contadorSeguridad == 1000 {
			msg := "Semantic error, infinite loop." +
				" -- Line:" + strconv.Itoa(l.Fila) + " Column: " +
				strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

		//Ejecutar las instrucciones
		for i = 0; i < l.Instrucciones.Len(); i++ {
			instruccion = l.Instrucciones.GetValue(i)
			//Recuperar el tipo de la instrucción
			tipoGeneral, tipoParticular = instruccion.(Ast.Abstracto).GetTipo()
			//Verificar tipos
			if tipoGeneral == Ast.EXPRESION {
				//Error, no pueden existir expresiones aisladas
				msg := "Semantic error, an instruction was expected." +
					" -- Line:" + strconv.Itoa(instruccion.(Ast.Abstracto).GetFila()) + " Column: " +
					strconv.Itoa(instruccion.(Ast.Abstracto).GetColumna())
				nError := errores.NewError(instruccion.(Ast.Abstracto).GetFila(),
					instruccion.(Ast.Abstracto).GetColumna(), msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				newScope.Errores.Add(nError)
				newScope.Consola += msg + "\n"
				newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR,
				}
			}
			if tipoGeneral == Ast.INSTRUCCION {
				//Verificar los breaks y los return

				//Error, solo puede venir un break expresion
				if Ast.EsTransferencia(tipoParticular) &&
					l.Tipo == Ast.LOOP_EXPRESION {
					if tipoParticular != Ast.BREAK_EXPRESION &&
						tipoParticular != Ast.CONTINUE {
						msg := "Semantic error," + Ast.ValorTipoDato[tipoParticular] + " statement not allowed inside this kind of loop." +
							" -- Line:" + strconv.Itoa(instruccion.(Ast.Abstracto).GetFila()) + " Column: " +
							strconv.Itoa(instruccion.(Ast.Abstracto).GetColumna())
						nError := errores.NewError(instruccion.(Ast.Abstracto).GetFila(),
							instruccion.(Ast.Abstracto).GetColumna(), msg)
						nError.Tipo = Ast.ERROR_SEMANTICO
						newScope.Errores.Add(nError)
						newScope.Consola += msg + "\n"
						newScope.UpdateScopeGlobal()
						return Ast.TipoRetornado{
							Valor: nError,
							Tipo:  Ast.ERROR,
						}
					} else if tipoParticular == Ast.BREAK_EXPRESION {
						//Si ambos son de tipo expresion
						retornarBreak = true
					}

				}

				//Ejecutar la instrucción
				resultado = instruccion.(Ast.Instruccion).Run(scope)

			}
			if retornarBreak {
				//Terminar loop y retornar el valor del break
				newScope.UpdateScopeGlobal()
				return resultado.(Ast.TipoRetornado).Valor
			}

			if Ast.EsTransferencia(tipoParticular) &&
				l.Tipo == Ast.LOOP_EXPRESION {
				if tipoParticular != Ast.BREAK_EXPRESION &&
					tipoParticular != Ast.CONTINUE {
					msg := "Semantic error," + Ast.ValorTipoDato[tipoParticular] + " statement not allowed inside this kind of loop." +
						" -- Line:" + strconv.Itoa(instruccion.(Ast.Abstracto).GetFila()) + " Column: " +
						strconv.Itoa(instruccion.(Ast.Abstracto).GetColumna())
					nError := errores.NewError(instruccion.(Ast.Abstracto).GetFila(),
						instruccion.(Ast.Abstracto).GetColumna(), msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					newScope.Errores.Add(nError)
					newScope.Consola += msg + "\n"
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: nError,
						Tipo:  Ast.ERROR,
					}
				} else if tipoParticular == Ast.BREAK_EXPRESION {
					//Si ambos son de tipo expresion
					retornarBreak = true
				}

			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.EJECUTADO {
				//Siguiente instrucción
				continue
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.CONTINUE {
				//Siguiente iteración
				break
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK {
				//Terminar el loop
				newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.EJECUTADO,
					Valor: true,
				}
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION {
				newScope.UpdateScopeGlobal()
				//Terminar loop y retornar el return
				return resultado
			}
		}
		contadorSeguridad++
	}
}

func (op Loop) GetFila() int {
	return op.Fila
}
func (op Loop) GetColumna() int {
	return op.Columna
}
