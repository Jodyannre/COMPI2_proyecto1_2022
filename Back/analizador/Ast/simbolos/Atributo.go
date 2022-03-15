package simbolos

import (
	"Back/analizador/Ast"
)

type Atributo struct {
	Nombre       string
	Tipo         Ast.TipoDato
	Valor        interface{}
	TipoAtributo interface{}
	Fila         int
	Columna      int
	Publico      bool
	Mutable      bool
}

func NewAtributoTemplate(nombre string, tipo interface{}, publico bool, fila, columna int) Atributo {
	nuevo := Atributo{
		Nombre:       nombre,
		Tipo:         Ast.ATRIBUTO,
		TipoAtributo: tipo,
		Fila:         fila,
		Columna:      columna,
		Publico:      publico,
	}
	return nuevo
}

func NewAtributo(nombre string, valor interface{}, mutable bool, fila, columna int) Atributo {
	nuevo := Atributo{
		Tipo:    Ast.ATRIBUTO,
		Nombre:  nombre,
		Fila:    fila,
		Valor:   valor,
		Columna: columna,
		Mutable: mutable,
	}
	return nuevo
}

func (a Atributo) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//calcular el valor y los tipos del atributo
	valor := a.Valor.(Ast.Expresion).GetValue(scope)

	if valor.Tipo == Ast.ERROR {
		return valor
	}

	a.Valor = valor
	if esTipoFinal(valor.Tipo) && valor.Tipo != Ast.STRUCT {
		a.TipoAtributo = NewTipo(valor.Tipo, true, a.Fila, a.Columna)
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.ATRIBUTO,
		Valor: a,
	}
}

func (p Atributo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, p.Tipo
}

func (p Atributo) GetFila() int {
	return p.Fila
}
func (p Atributo) GetColumna() int {
	return p.Columna
}
