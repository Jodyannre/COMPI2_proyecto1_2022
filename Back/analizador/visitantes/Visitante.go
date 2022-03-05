package visitantes

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/parser"
	"fmt"

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
				//Hay error y agregarlo a la lista de errores
				v.Errores.Add(respuesta)
				v.Consola += respuesta.(Ast.TipoRetornado).Valor.(errores.CustomSyntaxError).Msg + "\n"
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
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR_SEMANTICO {
				//Hay error y agregarlo a la lista de errores
				//v.Errores.Add(respuesta)
				EntornoGlobal.Errores.Add(respuesta)
				EntornoGlobal.Consola += respuesta.(Ast.TipoRetornado).Valor.(errores.CustomSyntaxError).Msg + "\n"
			} else
			//Verificar que sea un string para agregarlo a la consola
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.STRING {
				EntornoGlobal.Consola += respuesta.(Ast.TipoRetornado).Valor.(string) + "\n"
			}
		}
	}
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
