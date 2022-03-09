package simbolos

import "Back/analizador/Ast"

type Valor struct {
	Valor      Ast.Expresion
	Tipo       Ast.TipoDato
	Fila       int
	Columna    int
	Referencia bool
	Mutable    bool
}

func NewValor(valor Ast.Expresion, tipo Ast.TipoDato, referencia bool, mutable bool, fila, columna int) Valor {
	nV := Valor{
		Tipo:       tipo,
		Valor:      valor,
		Referencia: referencia,
		Mutable:    mutable,
		Fila:       fila,
		Columna:    columna,
	}
	return nV
}

func (v Valor) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	return v.Valor.GetValue(scope)
}

func (p Valor) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, p.Tipo
}

func (p Valor) GetFila() int {
	return p.Fila
}
func (p Valor) GetColumna() int {
	return p.Columna
}
