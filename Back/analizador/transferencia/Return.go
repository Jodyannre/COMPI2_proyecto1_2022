package transferencia

import "Back/analizador/Ast"

type Return struct {
	Tipo      Ast.TipoDato
	Expresion Ast.Expresion
	Valor     Ast.TipoRetornado
	Fila      int
	Columna   int
}

func (r Return) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, r.Tipo
}

func NewReturn(tipo Ast.TipoDato, expresion Ast.Expresion, fila, columna int) Return {
	nR := Return{
		Tipo:      tipo,
		Expresion: expresion,
		Valor:     Ast.TipoRetornado{Valor: false, Tipo: Ast.VOID},
		Fila:      fila,
		Columna:   columna,
	}
	return nR
}

func (r Return) Run(scope *Ast.Scope) interface{} {
	valor := r.Expresion.GetValue(scope)

	if valor.Tipo == Ast.ERROR_SEMANTICO {
		return valor
	}

	if r.Tipo == Ast.RETURN {

		return Ast.TipoRetornado{
			Tipo:  r.Tipo,
			Valor: false,
		}
	}

	return Ast.TipoRetornado{
		Tipo: r.Tipo,
		Valor: Ast.TipoRetornado{
			Tipo:  valor.Tipo,
			Valor: valor.Valor,
		},
	}
}

func (op Return) GetFila() int {
	return op.Fila
}
func (op Return) GetColumna() int {
	return op.Columna
}
