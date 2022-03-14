package simbolos

import (
	"Back/analizador/Ast"

	"github.com/colegno/arraylist"
)

type StructTemplate struct {
	Tipo        Ast.TipoDato
	Nombre      string
	Atributos   map[string]interface{}
	AtributosIn *arraylist.List
	Publico     bool
	Fila        int
	Columna     int
}

func NewStructTemplate(nombre string, atributos *arraylist.List, publico bool, fila, columna int) StructTemplate {
	att := make(map[string]interface{})
	//Agregar los elementos al nuevo struct template
	nuevo := StructTemplate{
		Tipo:        Ast.STRUCT,
		Nombre:      nombre,
		Atributos:   att,
		Publico:     publico,
		Fila:        fila,
		Columna:     columna,
		AtributosIn: atributos,
	}
	return nuevo
}

func (s StructTemplate) GetValue(scope *Ast.Scope) interface{} {
	sinAtributos := false
	if s.AtributosIn.Len() == 0 {
		//No tiene atributos, no sé si es error
		sinAtributos = true
	}
	if !sinAtributos {
		for i := 0; i < s.AtributosIn.Len(); i++ {
			att_val := s.AtributosIn.GetValue(i).(Atributo)
			for key, _ := range s.Atributos {
				if key == att_val.Nombre {
					return Ast.TipoRetornado{Valor: "Error, elemento repetido", Tipo: Ast.ERROR}
				}
			}
			s.Atributos[att_val.Nombre] = att_val.Tipo
		}
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.STRUCT,
		Valor: s,
	}
}
