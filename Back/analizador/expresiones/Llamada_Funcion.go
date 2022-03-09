package expresiones

import (
	"Back/analizador/Ast"
	"fmt"

	"github.com/colegno/arraylist"
)

type LlamadaFuncion struct {
	Identificador Ast.Expresion
	Parametros    *arraylist.List
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewLlamadaFuncion(id Ast.Expresion, parametros *arraylist.List, tipo Ast.TipoDato, fila, columna int) LlamadaFuncion {
	nF := LlamadaFuncion{
		Identificador: id,
		Parametros:    parametros,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nF
}

func (l LlamadaFuncion) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	existe := false

	//Crear el scope para la nueva función
	newScope := Ast.NewScope("funcion", scope)

	//Verificar que la función existe
	existe = newScope.Exist(l.Identificador.(Identificador).Valor)
	if !existe {
		//Error la función no existe
		fmt.Println()
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (l LlamadaFuncion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, l.Tipo
}

func (l LlamadaFuncion) GetFila() int {
	return l.Fila
}
func (l LlamadaFuncion) GetColumna() int {
	return l.Columna
}
