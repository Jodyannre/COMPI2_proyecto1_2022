package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
)

type DeclaracionVector struct {
	Id            string
	Tipo          Ast.TipoDato
	TipoVector    Ast.TipoDato
	TipoDelVector Ast.TipoDato
	TipoDelStruct string
	TipoDelArray  Ast.TipoDato
	Mutable       bool
	Publico       bool
	Valor         interface{}
	Fila          int
	Columna       int
}

func NewDeclaracionVector(id string, mutable, publico bool, tipoVector Ast.TipoDato, tipoStruct string,
	valor interface{}, fila int, columna int) DeclaracionVector {
	nd := DeclaracionVector{
		Id:            id,
		Tipo:          Ast.DECLARACION,
		TipoDelVector: tipoVector,
		TipoDelStruct: tipoStruct,
		Mutable:       mutable,
		Publico:       publico,
		Valor:         valor,
		Fila:          fila,
		Columna:       columna,
	}
	return nd
}

func (d DeclaracionVector) Run(scope *Ast.Scope) interface{} {
	//Verificar que no exista

	existe := scope.Exist_actual(d.Id)
	//Calcular el valor del elemento a asignar
	valor := d.Valor.(Ast.Expresion).GetValue(scope)

	//Primero verificar que no es un if expresion
	tipoIn := valor.Tipo

	//Si es diferente de vector error
	if tipoIn != Ast.VECTOR {
		msg := "Semantic error, can't initialize a" + Ast.ValorTipoDato[d.Tipo] + "with " + Ast.ValorTipoDato[tipoIn] + " value." +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar si ya existe
	if existe {
		msg := "Semantic error, the element \"" + d.Id + "\" already exist in this scope." +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar que los tipos sean correctos

	return nil
}

func (op DeclaracionVector) GetFila() int {
	return op.Fila
}
func (op DeclaracionVector) GetColumna() int {
	return op.Columna
}

func (d DeclaracionVector) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}
