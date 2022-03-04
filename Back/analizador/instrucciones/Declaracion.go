package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"fmt"
	"strconv"
)

type Declaracion struct {
	Id          string
	Tipo        Ast.TipoDato
	TipoRetorno Ast.TipoDato
	Mutable     bool
	Publico     bool
	Valor       interface{}
	Fila        int
	Columna     int
}

func NewDeclaracion(id string, tipo Ast.TipoDato, mutable, publico bool, tipoRetorno Ast.TipoDato,
	valor interface{}, fila int, columna int) Declaracion {
	nd := Declaracion{
		Id:          id,
		Tipo:        tipo,
		TipoRetorno: tipoRetorno,
		Mutable:     mutable,
		Publico:     publico,
		Valor:       valor,
		Fila:        fila,
		Columna:     columna,
	}
	return nd
}

func (d Declaracion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d Declaracion) Run(scope *Ast.Scope) interface{} {
	//Verificar que el id no exista

	existe := scope.Exist_actual(d.Id)

	//Verificar que sea un modulo, función o vector
	if d.Tipo == Ast.MODULO {
		fmt.Println("Declaración de módulo")
	}
	if d.Tipo == Ast.FUNCION {
		fmt.Println("Declaración de Función")
	}
	if d.Tipo == Ast.VECTOR {
		fmt.Println("Declaración de vector")
	}
	if d.Tipo == Ast.STRUCT {
		fmt.Println("Declaración de struct")
	}

	//Verificar que los tipos sean correctos si

	//Primero verificar que no es un if expresion
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()
	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION {
		preValor = d.Valor.(Ast.Instruccion).Run(scope)
	} else {
		preValor = d.Valor.(Ast.Expresion).GetValue(scope)
	}
	valor := preValor.(Ast.TipoRetornado)

	//Revisar si el retorno es un error
	if valor.Tipo == Ast.ERROR_SEMANTICO {
		return valor
	}

	if valor.Tipo == d.Tipo && !existe {
		//El tipo es correcto y no existe en el entorno actual
		//Crear símbolo y agregarlo a la tabla del entorno actual
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          d.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
		}
		scope.Add(nSimbolo)
	} else if d.Tipo == Ast.INDEFINIDO && !existe {
		//Es una declaración sin valor asignado
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          valor.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
		}
		scope.Add(nSimbolo)
	} else if d.Tipo != Ast.INDEFINIDO && !existe && valor.Tipo == Ast.NULL {
		//Es una declaración sin valor asignado
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          valor.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
		}
		scope.Add(nSimbolo)
	} else if existe {
		//Ya existe y generar error semántico
		//fmt.Println("Error, ese elemento ya existe en el ámbito local")
		msg := "Semantic error, the element \"" + d.Id + "\" already exist in this scope." +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	} else {
		//Error de tipos, generar error semántico
		//fmt.Println("Error, los tipos no coinciden en la declaración")
		msg := "Semantic error, can't assign " + Ast.ValorTipoDato[int(valor.Tipo)] +
			" type to " + Ast.ValorTipoDato[int(d.Tipo)] +
			" type. -- Line: " + strconv.Itoa(d.Fila) +
			" Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

/*
func (d Declaracion) DeclararModulo(scope *Ast.Scope) interface{}{

}
*/
