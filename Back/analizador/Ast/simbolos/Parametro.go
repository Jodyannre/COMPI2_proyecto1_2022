package simbolos

import "Back/analizador/Ast"

type Parametro struct {
	Identificador   string
	Tipo            Ast.TipoDato
	Fila            int
	Columna         int
	Mutable         bool
	TipoDeclaracion Ast.TipoDato
}

func NewParametro(id string, tipo Ast.TipoDato, tipoD Ast.TipoDato, mutable bool, fila, columna int) Parametro {
	nP := Parametro{
		Identificador:   id,
		Tipo:            tipo,
		Mutable:         mutable,
		Fila:            fila,
		Columna:         columna,
		TipoDeclaracion: tipoD,
	}
	return nP
}

func (p Parametro) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	return Ast.TipoRetornado{
		Tipo:  p.TipoDeclaracion,
		Valor: p.Identificador,
	}
}

func (p Parametro) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, p.Tipo
}

func (p Parametro) GetFila() int {
	return p.Fila
}
func (p Parametro) GetColumna() int {
	return p.Columna
}
