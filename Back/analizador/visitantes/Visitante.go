package visitantes

import (
	"Back/analizador/Ast"
	//"Back/analizador/errores"
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
	//fmt.Println(ctx.GetEx())
	instrucciones := ctx.GetEx()
	fmt.Println(instrucciones)
	//Creando el scope global
	EntornoGlobal := Ast.NewScope("global", nil)
	EntornoGlobal.Global = true
	/*
		var resultado Ast.TipoRetornado = instrucciones.(Ast.Expresion).GetValue(EntornoGlobal)
		fmt.Println(resultado)
	*/
	resultado := instrucciones.(Ast.Instruccion).Run(&EntornoGlobal)
	fmt.Println(resultado)
}

func (v *Visitador) GetConsola() string {
	//return v.Consola
	return v.Consola
}

func (v *Visitador) UpdateConsola(entrada string) {
	v.Consola += entrada + "\n"
}
