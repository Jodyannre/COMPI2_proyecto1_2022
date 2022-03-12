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
	//Verificar que no existan errores sintácticos o semánticos
	if v.Errores.Len() > 0 {
		return
	}

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
	var tipo, tipoGeneral interface{}
	var respuesta interface{}

	//Primera pasada para agregar todas las declaraciones de las variables
	for i := 0; i < instrucciones.Len(); i++ {
		actual = instrucciones.GetValue(i)
		if actual != nil {
			tipoGeneral, tipo = actual.(Ast.Abstracto).GetTipo()
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
			tipoGeneral, tipo = actual.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}
		//Error, no puede haber expresiones sueltas
		if tipoGeneral == Ast.EXPRESION && !EsFuncion(tipo) {
			fila := actual.(Ast.Abstracto).GetFila()
			columna := actual.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, an instruction was expected." +
				" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			EntornoGlobal.Errores.Add(nError)
			EntornoGlobal.Consola += msg + "\n"
			continue
		} else if EsFuncion(tipo) {
			//LLamandas a funciones y métodos de vectores o arrays o nativos
			respuesta = actual.(Ast.Expresion).GetValue(&EntornoGlobal)
		} else if tipo.(Ast.TipoDato) != Ast.DECLARACION {
			//Declarar variables globales
			respuesta = actual.(Ast.Instruccion).Run(&EntornoGlobal)
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				//Hay error, pero ya esta agregado
				continue
			}

			if Ast.EsTransferencia(respuesta.(Ast.TipoRetornado).Tipo) {
				//Error de break o return
				valor := actual.(Ast.Abstracto)
				fila := valor.GetFila()
				columna := valor.GetColumna()
				msg := "Semantic error," + Ast.ValorTipoDato[respuesta.(Ast.TipoRetornado).Tipo] +
					" statement not allowed in main method." +
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

func EsFuncion(tipo interface{}) bool {
	validador := false

	switch tipo {
	case Ast.FUNCION, Ast.VEC_NEW,
		Ast.VEC_LEN, Ast.VEC_CONTAINS,
		Ast.VEC_CAPACITY, Ast.VEC_REMOVE:
		validador = true
	default:
		validador = false
	}

	return validador
}
