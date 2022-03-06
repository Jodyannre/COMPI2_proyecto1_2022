package transferencia

import "Back/analizador/Ast"

type Break struct {
	Tipo      Ast.TipoDato
	Expresion Ast.Expresion
	Valor     Ast.TipoRetornado
	Fila      int
	Columna   int
}

func (b Break) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, b.Tipo
}

func NewBreak(tipo Ast.TipoDato, expresion Ast.Expresion, fila, columna int) Break {
	nB := Break{
		Tipo:      tipo,
		Expresion: expresion,
		Valor:     Ast.TipoRetornado{Valor: false, Tipo: Ast.VOID},
		Fila:      fila,
		Columna:   columna,
	}
	return nB
}

func (b Break) Run(scope *Ast.Scope) interface{} {

	if b.Tipo == Ast.BREAK {

		return Ast.TipoRetornado{
			Tipo:  Ast.BREAK,
			Valor: false,
		}
	}
	valor := b.Expresion.GetValue(scope)
	if valor.Tipo == Ast.ERROR_SEMANTICO {
		return valor
	}

	return Ast.TipoRetornado{
		Tipo: Ast.BREAK_EXPRESION,
		Valor: Ast.TipoRetornado{
			Tipo:  valor.Tipo,
			Valor: valor.Valor,
		},
	}
}

func (op Break) GetFila() int {
	return op.Fila
}
func (op Break) GetColumna() int {
	return op.Columna
}
