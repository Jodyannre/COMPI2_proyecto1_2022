package simbolos

import "Back/analizador/Ast"

type Atributo struct {
	Nombre  string
	Tipo    Ast.TipoDato
	Valor   interface{}
	Fila    int
	Columna int
	Publico bool
}

func NewAtributoTemplate(nombre string, tipo Ast.TipoDato, publico bool, fila, columna int) Atributo {
	nuevo := Atributo{
		Nombre:  nombre,
		Tipo:    tipo,
		Fila:    fila,
		Columna: columna,
		Publico: publico,
	}
	return nuevo
}

func NewAtributo(nombre string, valor interface{}, fila, columna int) Atributo {
	nuevo := Atributo{
		Nombre:  nombre,
		Fila:    fila,
		Valor:   valor,
		Columna: columna,
	}
	return nuevo
}

func (p Atributo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.SIMBOLO, p.Tipo
}

func (p Atributo) GetFila() int {
	return p.Fila
}
func (p Atributo) GetColumna() int {
	return p.Columna
}
