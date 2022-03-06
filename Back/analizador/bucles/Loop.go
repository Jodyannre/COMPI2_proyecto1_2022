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
	var instruccion, tipoGeneral, resultado interface{}
	var i int = 0

	for {

		if contadorSeguridad == 25000 {
			msg := "Semantic error, infinite loop." +
				" -- Line:" + strconv.Itoa(l.Fila) + " Column: " +
				strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR_SEMANTICO,
				Valor: nError,
			}
		}

		//Ejecutar las instrucciones
		for i = 0; i < l.Instrucciones.Len(); i++ {
			instruccion = l.Instrucciones.GetValue(i)
			//Recuperar el tipo de la instrucción
			tipoGeneral, _ = instruccion.(Ast.Abstracto).GetTipo()
			//Verificar tipos
			if tipoGeneral.(Ast.TipoDato) == Ast.EXPRESION {
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
					Tipo:  Ast.ERROR_SEMANTICO,
				}
			}
			if tipoGeneral.(Ast.TipoDato) == Ast.INSTRUCCION {
				//Ejecutar la instrucción
				resultado = instruccion.(Ast.Instruccion).Run(scope)
			}
			if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR_SEMANTICO {
				continue
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR_SEMANTICO_NO {
				error := resultado.(Ast.TipoRetornado).Valor.(errores.CustomSyntaxError)
				newScope.Errores.Add(error)
				newScope.Consola += error.Msg + "\n"
				continue
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.EJECUTADO {
				continue
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK && l.Tipo != Ast.LOOP_EXPRESION {
				newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Valor: true,
					Tipo:  Ast.EJECUTADO,
				}
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK && l.Tipo == Ast.LOOP_EXPRESION {
				//ERROR, tiene que retornar algo
				msg := "Semantic error, break statement is not returning any value." +
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
					Tipo:  Ast.ERROR_SEMANTICO,
				}
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION && l.Tipo == Ast.LOOP_EXPRESION {
				//Retorna un valor
				valor := resultado.(Ast.TipoRetornado).Valor
				newScope.UpdateScopeGlobal()
				return valor
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION && l.Tipo != Ast.LOOP_EXPRESION {
				//Error, está retornando un valor dentro de un loop instrucción
				msg := "Semantic error, break statement is returning a value within a LOOP statement." +
					" -- Line:" + strconv.Itoa(instruccion.(Ast.Abstracto).GetFila()) + " Column: " +
					strconv.Itoa(instruccion.(Ast.Abstracto).GetColumna())
				nError := errores.NewError(instruccion.(Ast.Abstracto).GetFila(),
					instruccion.(Ast.Abstracto).GetColumna(), msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				newScope.Errores.Add(nError)
				newScope.Consola += msg + "\n"
				newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR_SEMANTICO,
					Valor: nError,
				}
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
