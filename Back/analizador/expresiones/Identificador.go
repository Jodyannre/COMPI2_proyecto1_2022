package expresiones

import (
	"Back/analizador/Ast"
)

type Identificador struct {
	Tipo    Ast.TipoDato
	Valor   string
	Fila    int
	Columna int
}

func (p Identificador) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, Ast.IDENTIFICADOR
}

func (p Identificador) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Buscar el símbolo en la tabla de símbolos y retornar el valor
	//Verificar que el id no exista
	if scope.Exist(p.Valor) {
		//Existe el identificar y retornar el valor
		simbolo := scope.GetSimbolo(p.Valor)
		return simbolo.Valor.(Ast.TipoRetornado)
	} else {
		//No existe el identificador, retornar error semantico
		return Ast.TipoRetornado{Valor: nil, Tipo: Ast.NULL}
	}
}

func NewIdentificador(val string, tipo Ast.TipoDato, fila, columna int) Identificador {
	return Identificador{Tipo: tipo, Valor: val, Fila: fila, Columna: columna}
}

func (op Identificador) GetFila() int {
	return op.Fila
}
func (op Identificador) GetColumna() int {
	return op.Columna
}
