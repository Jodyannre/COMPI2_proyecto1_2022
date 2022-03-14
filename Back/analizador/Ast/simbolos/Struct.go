package simbolos

import (
	"Back/analizador/Ast"
	"strconv"

	"github.com/colegno/arraylist"
)

func NewStruct(plantilla StructTemplate, atributos *arraylist.List, scope *Ast.Scope) interface{} {
	nuevo := plantilla
	var att Atributo
	if len(nuevo.Atributos) != atributos.Len() {
		//Error la cantidad de atributos no concuerda
	}

	for i := 0; i < atributos.Len(); i++ {
		att = atributos.GetValue(i).(Atributo)
		if val, ok := nuevo.Atributos[att.Nombre]; ok {
			//Verificar que los tipos esten bien
			resultado := ValidarTipo(val.(Atributo).Tipo, att.Valor.(Ast.Expresion), scope, nuevo.Fila, nuevo.Columna)
			if resultado.Tipo == Ast.ERROR_SEMANTICO {
				return resultado
			}
		} else {
			return att
			//Error ese atributo no existe
		}
	}
	return nil
}

func ValidarTipo(tipo Ast.TipoDato, expresion Ast.Expresion, scope *Ast.Scope, fila, columna int) Ast.TipoRetornado {
	resultado := expresion.GetValue(scope)

	if tipo == resultado.Tipo {
		//Todo bien retornar el resultado
		return resultado
	}
	//Sino error semántico de tipos
	msg := "Semantic error, can't assign " + Ast.ValorTipoDato[resultado.Tipo] +
		" type to " + Ast.ValorTipoDato[tipo] +
		" type. -- Line: " + strconv.Itoa(fila) +
		" Column: " + strconv.Itoa(columna)
	return Ast.TipoRetornado{Tipo: Ast.ERROR_SEMANTICO, Valor: msg}
}
