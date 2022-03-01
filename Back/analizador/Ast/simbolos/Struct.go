package simbolos

import (
	"Back/analizador/Ast"
	"strconv"

	"github.com/colegno/arraylist"
)

type Struct struct {
	Nombre    string
	Atributos map[string]interface{}
	Fila      int
	Columna   int
}

type Atributo struct {
	Nombre string
	Tipo   Ast.TipoDato
	Valor  interface{}
}

func NewStructTemplate(nombre string, atributos *arraylist.List, fila, columna int) Struct {
	att := make(map[string]interface{})
	var att_val Atributo
	//Agregar los atributos al struct template
	for i := 0; i < atributos.Len(); i++ {
		att_val = atributos.GetValue(i).(Atributo)
		att[att_val.Nombre] = att_val
	}

	nuevo := Struct{
		Nombre:    nombre,
		Atributos: att,
		Fila:      fila,
		Columna:   columna,
	}
	return nuevo
}

func NewAtributoTemplate(nombre string, tipo Ast.TipoDato) Atributo {
	nuevo := Atributo{
		Nombre: nombre,
		Tipo:   tipo,
	}
	return nuevo
}

func NewStruct(plantilla Struct, atributos *arraylist.List, scope *Ast.Scope) interface{} {
	nuevo := plantilla
	var att Atributo
	if len(nuevo.Atributos) != atributos.Len() {
		//Error la cantidad de atributos no concuerda
	}

	for i := 0; i < atributos.Len(); i++ {
		att = atributos.GetValue(i).(Atributo)
		if val, ok := nuevo.Atributos[att.Nombre]; ok {
			//Verificar que los tipos esten bien
			resultado := ValidarTipo(val.(Atributo).Tipo, att.Valor.(Ast.Expresion), *scope, nuevo.Fila, nuevo.Columna)
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

func ValidarTipo(tipo Ast.TipoDato, expresion Ast.Expresion, scope Ast.Scope, fila, columna int) Ast.TipoRetornado {
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
