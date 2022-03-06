package visitantes

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/parser"
	"fmt"
	"strconv"

	"github.com/colegno/arraylist"
)

type Visitador struct {
	*parser.BaseNparserListener
	Consola       string
	Errores       arraylist.List
	EntornoGlobal Ast.Scope
}

func NewVisitor() *Visitador {
	return new(Visitador)
}

func (v *Visitador) ExitInicio(ctx *parser.InicioContext) {
	fmt.Println("Ingreso>>>>>>>>>>>>>>>")
	instrucciones := ctx.GetLista()
	//fmt.Println(instrucciones)
	//Desde aquí ya tengo todo el resultado del parser, listo para ejecutar

	/*
		1) Declarar el scope
		2) Primera pasada para declarar todas las variables
		3) Correr todas las instrucciones
	*/

	//Creando el scope global
	EntornoGlobal := Ast.NewScope("global", nil)
	EntornoGlobal.Global = true
	//Variables para las declaraciones
	var actual interface{}
	var tipo, _ interface{}
	var respuesta interface{}

	//Primera pasada para agregar todas las declaraciones de las variables
	for i := 0; i < instrucciones.Len(); i++ {
		actual = instrucciones.GetValue(i)
		if actual != nil {
			_, tipo = actual.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}
		if tipo.(Ast.TipoDato) == Ast.DECLARACION {
			//Declarar variables globales
			respuesta = actual.(Ast.Instruccion).Run(&EntornoGlobal)
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				//Hay error pero ya esta agregado
				continue
			}
		}
	}

	//Ejecutar todas instrucciones siguientes
	for i := 0; i < instrucciones.Len(); i++ {
		actual = instrucciones.GetValue(i)
		if actual != nil {
			_, tipo = actual.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}
		if tipo.(Ast.TipoDato) != Ast.DECLARACION {
			//Declarar variables globales
			respuesta = actual.(Ast.Instruccion).Run(&EntornoGlobal)
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				//Hay error, pero ya esta agregado
				continue
			}

			if respuesta.(Ast.TipoRetornado).Tipo == Ast.BREAK ||
				respuesta.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION {
				//Error de break o return
				fila := respuesta.(Ast.Abstracto).GetFila()
				columna := respuesta.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, break statement outside a loop" +
					" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				EntornoGlobal.Errores.Add(nError)
				EntornoGlobal.Consola += msg + "\n"
			}

			if respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN ||
				respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
				fila := respuesta.(Ast.Abstracto).GetFila()
				columna := respuesta.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, return statement outside a function." +
					" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				EntornoGlobal.Errores.Add(nError)
				EntornoGlobal.Consola += msg + "\n"
			}
		}
	}
	EntornoGlobal.UpdateScopeGlobal()
	v.Consola += EntornoGlobal.Consola
	for i := 0; i < EntornoGlobal.Errores.Len(); i++ {
		v.Errores.Add(EntornoGlobal.Errores.GetValue(i))
	}
}

func (v *Visitador) GetConsola() string {
	//return v.Consola
	return v.Consola
}

func (v *Visitador) UpdateConsola(entrada string) {
	v.Consola += entrada + "\n"
}
