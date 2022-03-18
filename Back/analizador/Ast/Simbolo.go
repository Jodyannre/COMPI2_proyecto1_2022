package Ast

type Simbolo struct {
	Identificador      string
	Valor              interface{}
	Fila               int
	Columna            int
	Tipo               TipoDato
	TipoEspecial       TipoRetornado
	Mutable            bool
	Publico            bool
	Entorno            *Scope
	Referencia         bool
	Referencia_puntero *Simbolo
}

type SimboloReporte struct {
	Identificador string
	TipoSimbolo   string
	TipoDato      string
	Scope         string
	Fila          int
	Columna       int
}

func NewSimbolo(identificador string, valor interface{}, fila int, columna int,
	tipo TipoDato, mutable bool, publico bool) Simbolo {
	simbolo := Simbolo{
		Identificador:      identificador,
		Valor:              valor,
		Fila:               fila,
		Columna:            columna,
		Tipo:               tipo,
		Mutable:            mutable,
		Publico:            publico,
		Referencia:         false,
		Referencia_puntero: nil,
		Entorno:            nil,
		TipoEspecial:       TipoRetornado{Valor: true, Tipo: INDEFINIDO},
	}
	return simbolo
}

func (s Simbolo) NewSimboloReporte(scope *Scope) SimboloReporte {
	var tipo string
	var nombreScope string
	if s.Tipo == FUNCION {
		tipo = ValorTipoDato[FUNCION]
	} else if s.Tipo == MODULO {
		tipo = ValorTipoDato[MODULO]
	} else {
		tipo = ValorTipoDato[VARIABLE]
	}

	if scope.Global {
		nombreScope = "Global"
	} else {
		nombreScope = "Local"
	}

	return SimboloReporte{
		Identificador: s.Identificador,
		TipoSimbolo:   tipo,
		TipoDato:      ValorTipoDato[s.Valor.(TipoRetornado).Tipo],
		Scope:         nombreScope,
		Fila:          s.Fila,
		Columna:       s.Columna,
	}
}
