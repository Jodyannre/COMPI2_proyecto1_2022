package Ast

type Simbolo struct {
	Identificador string
	Valor         interface{}
	Fila          int
	Columna       int
	Tipo          TipoDato
	Funcion       bool
	Mutable       bool
	Publico       bool
}

type SimboloReporte struct {
	Identificador string
	TipoSimbolo   string
	TipoDato      string
	Scope         string
	Fila          string
	Columna       string
}

func NewSimbolo(identificador string, valor interface{}, fila int, columna int,
	tipo TipoDato, funcion bool, mutable bool, publico bool) Simbolo {
	simbolo := Simbolo{
		Identificador: identificador,
		Valor:         valor,
		Fila:          fila,
		Columna:       columna,
		Tipo:          tipo,
		Funcion:       funcion,
		Mutable:       mutable,
		Publico:       publico,
	}
	return simbolo
}
