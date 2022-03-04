package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
)

type Imprimir struct {
	Expresiones Ast.Expresion
	Fila        int
	Columna     int
}

func NewImprimir(val Ast.Expresion, fila, columna int) Imprimir {
	nuevo := Imprimir{Expresiones: val, Fila: fila, Columna: columna}
	return nuevo
}

func (i Imprimir) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.PRINT
}

func (i Imprimir) Run(scope *Ast.Scope) interface{} {
	resultado_expresion := i.Expresiones.GetValue(scope)
	retorno := Ast.TipoRetornado{}
	retorno.Tipo = Ast.STRING
	switch resultado_expresion.Tipo {
	case Ast.BOOLEAN:
		retorno.Valor = strconv.FormatBool(resultado_expresion.Valor.(bool))
	case Ast.I64:
		retorno.Valor = strconv.Itoa(resultado_expresion.Valor.(int))
	case Ast.F64:
		retorno.Valor = strconv.FormatFloat(resultado_expresion.Valor.(float64), 'E', -1, 64)
	case Ast.STRING:
		retorno.Valor = resultado_expresion.Valor.(string)
	default:
		//Error, no es un tipo que se pueda imprimir
		//O es una operación que dio como resultado null
		//No existe, generar un error semántico
		msg := "Semantic error, error in the parameters of the \"Print\" function." +
			" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
		nError := errores.NewError(i.Fila, i.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	return retorno
}
