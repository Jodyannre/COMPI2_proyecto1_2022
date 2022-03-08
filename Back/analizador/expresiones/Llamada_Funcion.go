package expresiones

import (
	"Back/analizador/Ast"

	"github.com/colegno/arraylist"
)

type LlamadaFuncion struct {
	Identificador string
	Parametros    *arraylist.List
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func (l LlamadaFuncion) Run(scope *Ast.Scope) interface{} {
	existe := false

	//Crear el scope para la nueva función
	newScope := Ast.NewScope("funcion", scope)

	//Verificar que la función existe
	existe = newScope.Exist(l.Identificador)
	if !existe {
		//Error la función no existe
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
