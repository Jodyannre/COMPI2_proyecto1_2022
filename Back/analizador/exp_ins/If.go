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
		if i.Tipo == Ast.ELSEIF {
			//Es un if
			//i = elemento
			scope.Nombre = "ELSEIF"
			condicion1 = i.Condicion.GetValue(scope)
		} else if i.Tipo == Ast.ELSE {
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
				if tipo_abstracto == Ast.EXPRESION && i.Tipo == Ast.IF_EXPRESION {
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
							Tipo:  Ast.ERROR_SEMANTICO,
						}
					}
					expresion := i.Instrucciones.GetValue(n).(Ast.Expresion)
					ultimaExpresion = expresion.GetValue(scope)
				} else if tipo_abstracto == Ast.INSTRUCCION {
					instruccion := i.Instrucciones.GetValue(n).(Ast.Instruccion)
					resultado := instruccion.Run(scope)
					//Verificar si viene un break o un return
					if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK ||
						resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION ||
						resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN ||
						resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
						return resultado.(Ast.TipoRetornado)
					}

					//Verificar si el resultado es un bool o un string
					if resultado.(Ast.TipoRetornado).Tipo == Ast.STRING {
						//Agregar a la consola
						scope.Consola += resultado.(Ast.TipoRetornado).Valor.(string) + "\n"
					}
					/*
						if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR_SEMANTICO {
							//No hace nada
						}
					*/
					if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR_SEMANTICO_NO {
						error := resultado.(Ast.TipoRetornado).Valor.(errores.CustomSyntaxError)
						scope.Errores.Add(error)
						scope.Consola += error.Msg + "\n"
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
						Tipo:  Ast.ERROR_SEMANTICO,
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
					Tipo:  Ast.ERROR_SEMANTICO,
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
					newScope.Errores.Add(resultado.Valor)
					newScope.Consola += resultado.Valor.(errores.CustomSyntaxError).Msg + "\n"
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: resultado.Valor,
						Tipo:  Ast.ERROR_SEMANTICO,
					}
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
