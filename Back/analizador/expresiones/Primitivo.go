package expresiones

import (
	"Back/analizador/Ast"
)

type Primitivo struct {
	Tipo  Ast.TipoDato
	Valor interface{}
}

func (p Primitivo) GetValue(entorno *Ast.Scope) Ast.TipoRetornado {
	return Ast.TipoRetornado{
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func (p Primitivo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, Ast.PRIMITIVO
}

func NewPrimitivo(val interface{}, tipo Ast.TipoDato) Primitivo {
	nuevo := Primitivo{Tipo: tipo, Valor: val}
	return nuevo
}
