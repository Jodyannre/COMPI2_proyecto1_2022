package Ast

type Expresion interface {
	GetValue(entorno *Scope) TipoRetornado
}

type Instruccion interface {
	Run(entorno *Scope) interface{}
}

type Abstracto interface {
	GetTipo() (TipoDato, TipoDato)
	GetFila() int
	GetColumna() int
}

type Structs interface {
	GetPlantilla(scope *Scope) string
	SetMutabilidad(mutable bool)
}

type AccesosM interface {
	GetTipoFromAccesoModulo(tipo TipoRetornado, scope *Scope) TipoRetornado
	GetTipoEstructura(tipo TipoRetornado, scope *Scope) TipoRetornado
}

type AbstractoM interface {
	GetMutable() bool
}

type AccesoVectorAbstracto interface {
	GetIdentificador() string
}
