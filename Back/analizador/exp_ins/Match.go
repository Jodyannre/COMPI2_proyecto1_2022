package exp_ins

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type Match struct {
	Expresion Ast.Expresion
	Cases     *arraylist.List
	Tipo      Ast.TipoDato
	Fila      int
	Columna   int
}

type Case struct {
	Expresion     *arraylist.List
	Instrucciones *arraylist.List
	Tipo          Ast.TipoDato
	Default       bool
	Fila          int
	Columna       int
}

func NewMatch(expresion Ast.Expresion, cases *arraylist.List, tipo Ast.TipoDato, fila, columna int) Match {
	nMatch := Match{
		Expresion: expresion,
		Cases:     cases,
		Tipo:      tipo,
		Fila:      fila,
		Columna:   columna,
	}
	return nMatch
}

func NewCase(expresion *arraylist.List, instrucciones *arraylist.List, tipo Ast.TipoDato,
	fila, columna int, Default bool) Case {
	nCase := Case{
		Expresion:     expresion,
		Instrucciones: instrucciones,
		Tipo:          tipo,
		Default:       Default,
		Fila:          fila,
		Columna:       columna,
	}
	return nCase
}

func (m Match) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, m.Tipo
}

func (c Case) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, c.Tipo
}

func (m Match) Run(scope *Ast.Scope) interface{} {
	//Primero conseguir el exp_match de la expresion
	buscar_default := false
	default_econtrado := false
	exp_match := m.Expresion.GetValue(scope)
	pos_expresion_correcta := -1
	var i, j int

	//Validar si es un boolean y si tiene todos los casos
	if exp_match.Tipo == Ast.BOOLEAN && m.Cases.Len() < 2 {
		//Error, no estan todos los casos incluidos
		msg := "Semantic error, an instruction was expected." +
			" -- Line:" + strconv.Itoa(m.Fila) + " Column: " + strconv.Itoa(m.Columna)
		nError := errores.NewError(m.Fila, m.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		scope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Valor: nError,
			Tipo:  Ast.ERROR_SEMANTICO,
		}
	}

	//Verificar si no es boolean y se espera un default
	if exp_match.Tipo != Ast.BOOLEAN {
		buscar_default = true
	}

	for i = 0; i < m.Cases.Len(); i++ {
		//Recorrer la lista de cases y verificar que los exp_matches de las expresiones concuerdan
		caso := m.Cases.GetValue(i).(Case)
		listaExpresiones := caso.Expresion

		for j = 0; j < listaExpresiones.Len(); j++ {
			expresion := listaExpresiones.GetValue(j).(Ast.Expresion).GetValue(scope)

			if expresion.Tipo != exp_match.Tipo && !caso.Default {
				//Error de tipos
				msg := "Semantic error, the expression is not a " + Ast.ValorTipoDato[exp_match.Tipo] +
					" value. -- Line:" + strconv.Itoa(caso.Fila) + " Column: " + strconv.Itoa(caso.Columna)
				nError := errores.NewError(caso.Fila, caso.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR_SEMANTICO,
				}
			}

			if exp_match.Valor == expresion.Valor {
				pos_expresion_correcta = i
			}
		}
		//Verificar si viene default
		if buscar_default {
			if caso.Default {
				default_econtrado = true
			}
		}

		if default_econtrado && i != m.Cases.Len()-1 {
			//Error, el default tiene que venir de último
			msg := "Semantic error, the default case was expected in the last position." +
				" -- Line:" + strconv.Itoa(caso.Fila) + " Column: " + strconv.Itoa(caso.Columna)
			nError := errores.NewError(caso.Fila, caso.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			scope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR_SEMANTICO,
			}
		}

	}
	//Verificar que se necesite default y que lo traiga
	if buscar_default {
		if !default_econtrado {
			msg := "Semantic error, default case not found." +
				" -- Line:" + strconv.Itoa(m.Fila) + " Column: " + strconv.Itoa(m.Columna)
			nError := errores.NewError(m.Fila, m.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			scope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR_SEMANTICO,
			}
		}
	}

	//Todos los casos son correctos
	//Recorrer los cases
	var resultado_retornado Ast.TipoRetornado
	//Recuperar el caso en donde se encontró el valor igual
	if pos_expresion_correcta != -1 {
		caso := m.Cases.GetValue(pos_expresion_correcta).(Case)
		resultado_retornado = caso.Run(scope).(Ast.TipoRetornado)
	} else {
		//Ejecutar default
		caso := m.Cases.GetValue(m.Cases.Len() - 1).(Case)
		resultado_retornado = caso.Run(scope).(Ast.TipoRetornado)
	}

	if resultado_retornado.Tipo == Ast.ERROR_SEMANTICO || m.Tipo == Ast.MATCH_EXPRESION {
		return resultado_retornado
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: nil,
	}
}

func (c Case) Run(scope *Ast.Scope) interface{} {
	//Crear un nuevo scope y otras variables
	newScope := Ast.NewScope("Case", scope)
	var expresion = false
	ultimaExpresion := Ast.TipoRetornado{Valor: nil, Tipo: Ast.NULL}
	var instruccion interface{}
	i := 0

	//Verificar el tipo del caso, si es exp o ins
	if c.Tipo == Ast.CASE_EXPRESION {
		expresion = true
	}

	//Ejecutar las instrucciones

	for i = 0; i < c.Instrucciones.Len(); i++ {
		//Verificar el tipo de entrada
		instruccion = c.Instrucciones.GetValue(i).(Ast.Abstracto)
		tipo_abstracto, _ := instruccion.(Ast.Abstracto).GetTipo()

		if tipo_abstracto == Ast.EXPRESION && c.Tipo == Ast.CASE_EXPRESION {
			//Si es un if expresión, tiene que retornar algo

			//Verificar que sea la última o error
			if i != c.Instrucciones.Len()-1 {
				//Error porque no puede venir una expresión sin ser la última
				msg := "Semantic error, an instruction was expected." +
					" -- Line:" + strconv.Itoa(c.Fila) + " Column: " + strconv.Itoa(c.Columna)
				nError := errores.NewError(c.Fila, c.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				newScope.Errores.Add(nError)
				newScope.Consola += msg + "\n"
				newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR_SEMANTICO,
				}
			}

			expresion := c.Instrucciones.GetValue(i).(Ast.Expresion)
			ultimaExpresion = expresion.GetValue(&newScope)
		} else if tipo_abstracto == Ast.INSTRUCCION {

			instruccion := c.Instrucciones.GetValue(i).(Ast.Instruccion)
			resultado := instruccion.Run(&newScope)

			//Verificar si es un break o un return
			if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION &&
					c.Tipo != Ast.CASE_EXPRESION {
				return resultado.(Ast.TipoRetornado)
			}

			//Error si viene un break o un return dentro de un case expresion
			if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION &&
					c.Tipo == Ast.CASE_EXPRESION {
				temp := resultado.(Ast.TipoRetornado).Valor.(Ast.Abstracto)
				msg := "Semantic error, transfer statements are not allowed within a case statement." +
					" -- Line:" + strconv.Itoa(temp.GetFila()) + " Column: " + strconv.Itoa(temp.GetColumna())
				nError := errores.NewError(temp.GetFila(), temp.GetColumna(), msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				newScope.Errores.Add(nError)
				newScope.Consola += msg + "\n"
				newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR_SEMANTICO,
				}
			}

			//Verificar si el resultado es un bool o un string
			if resultado.(Ast.TipoRetornado).Tipo == Ast.STRING {
				//Agregar a la consola
				newScope.Consola += resultado.(Ast.TipoRetornado).Valor.(string) + "\n"
			}
			/*
				if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR_SEMANTICO {
					//Agregar a errores
					continue
				}
			*/
			if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR_SEMANTICO_NO {
				error := resultado.(Ast.TipoRetornado).Valor.(errores.CustomSyntaxError)
				newScope.Errores.Add(error)
				newScope.Consola += error.Msg + "\n"
			}

		} else if tipo_abstracto == Ast.EXPRESION {
			msg := "Semantic error, an instruction was expected." +
				" -- Line:" + strconv.Itoa(c.Fila) + " Column: " + strconv.Itoa(c.Columna)
			nError := errores.NewError(c.Fila, c.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR_SEMANTICO,
			}
		}
		i++
	}

	//Termino el for, retornar la ultima expresion
	//Verificar si hay algun retorno o retornar un error
	if ultimaExpresion.Tipo == Ast.NULL && expresion {
		msg := "Semantic error, the if clause is not returning any value." +
			" -- Line:" + strconv.Itoa(c.Fila) + " Column: " + strconv.Itoa(c.Columna)
		nError := errores.NewError(c.Fila, c.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		newScope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Valor: nError,
			Tipo:  Ast.ERROR_SEMANTICO,
		}
	} else if expresion {
		//Si esta retornado algún valor
		newScope.UpdateScopeGlobal()
		return ultimaExpresion
	}

	newScope.UpdateScopeGlobal()
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (op Match) GetFila() int {
	return op.Fila
}
func (op Match) GetColumna() int {
	return op.Columna
}
