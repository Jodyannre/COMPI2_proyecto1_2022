package Ast

type Expresion interface {
	GetValue(entorno Scope) TipoRetornado
}

type Instruccion interface {
	Run(entorno *Scope) interface{}
}

type Abstracto interface {
	GetTipo() (TipoDato, TipoDato)
}
