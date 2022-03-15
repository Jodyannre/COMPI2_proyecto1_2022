package simbolos

import (
	"Back/analizador/Ast"
)

type Tipo struct {
	Tipo     Ast.TipoDato
	TipoDato Ast.TipoDato
	Valor    interface{}
	Fila     int
	Columna  int
}

func NewTipo(tipo Ast.TipoDato, valor interface{}, fila, columna int) Tipo {
	nT := Tipo{
		Tipo:     Ast.TIPO,
		TipoDato: tipo,
		Valor:    valor,
		Fila:     fila,
		Columna:  columna,
	}
	return nT
}

func (t Tipo) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	return Ast.TipoRetornado{
		Tipo:  Ast.TIPO,
		Valor: t,
	}
}

func (op Tipo) GetFila() int {
	return op.Fila
}
func (op Tipo) GetColumna() int {
	return op.Columna
}

func (d Tipo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, Ast.TIPO
}

func (d Tipo) GetTipoFinal() Ast.TipoRetornado {
	return getTipoFinal(d)
}

func getTipoFinal(tipo Tipo) Ast.TipoRetornado {
	if esTipoFinal(tipo.TipoDato) {
		if tipo.Tipo != Ast.STRUCT {
			return Ast.TipoRetornado{
				Tipo:  tipo.TipoDato,
				Valor: true,
			}
		}
		return Ast.TipoRetornado{
			Tipo:  tipo.TipoDato,
			Valor: tipo.Valor,
		}
	} else {
		return getTipoFinal(tipo.Valor.(Tipo))
	}
}

func esTipoFinal(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.I64, Ast.F64, Ast.CHAR, Ast.STRING, Ast.BOOLEAN, Ast.STRUCT:
		return true
	default:
		return false
	}
}

func CompararTipos(tipoA Tipo, tipoB Tipo) bool {
	if esTipoFinal(tipoA.TipoDato) && esTipoFinal(tipoB.TipoDato) {
		if tipoA.TipoDato == tipoB.TipoDato {
			//Verificar si son structs
			if tipoA.TipoDato == Ast.STRUCT {
				if tipoA.Valor == tipoB.Valor {
					return true
				} else {
					return false
				}
			}
			return true
		} else {
			return false
		}
	}
	if esTipoFinal(tipoA.TipoDato) && !esTipoFinal(tipoB.TipoDato) ||
		!esTipoFinal(tipoA.TipoDato) && esTipoFinal(tipoB.TipoDato) {
		return false
	}
	return CompararTipos(tipoA.Valor.(Tipo), tipoB.Valor.(Tipo))
}

func Tipo_String(t Tipo) string {
	if t.TipoDato == Ast.VECTOR {
		return "Vec <" + Tipo_String(t.Valor.(Tipo)) + ">"
	} else {
		if t.TipoDato == Ast.STRUCT {
			return t.Valor.(string)
		} else {
			return Ast.ValorTipoDato[t.TipoDato]
		}
	}
}
