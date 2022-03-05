package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
)

type Print struct {
	Expresiones Ast.Expresion
	Fila        int
	Columna     int
}

func NewPrint(val Ast.Expresion, fila, columna int) Print {
	nuevo := Print{Expresiones: val, Fila: fila, Columna: columna}
	return nuevo
}

func (i Print) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.PRINT
}

func (i Print) Run(scope *Ast.Scope) interface{} {
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

func (op Print) GetFila() int {
	return op.Fila
}
func (op Print) GetColumna() int {
	return op.Columna
}
