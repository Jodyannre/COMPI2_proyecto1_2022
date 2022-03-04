package main

import (

	//"primeraGramatica/Analizador/Ast"

	//"fmt"
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/visitantes"
	"Back/parser"
	"fmt"

	//"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Key struct {
	Nombre string
	Padre  string
}

func main() {

	var input string = `
	let mut variable1:bool = if (true){
		false
	};
	`

	//Obteniendo el input
	cadena_entrada := antlr.NewInputStream(input)

	//Creando los costum errors
	errores_sintacticos := &errores.CustomErrorListener{}
	errores_lexicos := &errores.CustomErrorListener{}

	//Creando el lexer
	lexer := parser.NewNlexer(cadena_entrada)

	//Modificando el listener de los errores
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errores_lexicos)

	//Obteniendo los tokens
	streamTokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	//Creando el parser
	nparser := parser.NewNparser(streamTokens)

	//Modificando el listener de los errores
	nparser.RemoveErrorListeners()
	nparser.AddErrorListener(errores_sintacticos)
	nparser.BuildParseTrees = true

	//Leyendo la producción de inicio
	tree := nparser.Inicio()

	//fmt.Printf("errores léxicos \n %v", errores_lexicos.errores)
	//fmt.Printf("\nerrores sintáticos")
	//fmt.Printf("\n%v", errores_sintacticos.errores)

	//Visitador para recorrer el árbol
	var nvisitor = visitantes.NewVisitor()

	//Agregar errores léxicos y sintácticos al array de errores en el nVisitor

	//Recorrer el array de errores léxicos
	for i := 0; i < len(errores_lexicos.Errores); i++ {
		//fmt.Printf("Error #"+strconv.Itoa(i)+" %v \n", errores_lexicos.errores[i])
		errores_lexicos.Errores[i].Tipo = Ast.ERROR_LEXICO
		nvisitor.Errores.Add(errores_lexicos.Errores[i])
		//Actualizar la consola
		nvisitor.UpdateConsola("Lexical error, " + errores_lexicos.Errores[i].Msg + " -- Line: " +
			strconv.Itoa(errores_lexicos.Errores[i].Linea) + " Column: " +
			strconv.Itoa(errores_lexicos.Errores[i].Columna))
	}

	//Recorrer el array de errores sintácticos
	for i := 0; i < len(errores_sintacticos.Errores); i++ {
		//fmt.Printf("Error #"+strconv.Itoa(i)+" %v \n", errores_sintacticos.errores[i])
		errores_sintacticos.Errores[i].Tipo = Ast.ERROR_SINTACTICO
		nvisitor.Errores.Add(errores_sintacticos.Errores[i])
		//Actualizar la consola
		nvisitor.UpdateConsola("Syntax error , " + errores_sintacticos.Errores[i].Msg + " -- Line: " +
			strconv.Itoa(errores_sintacticos.Errores[i].Linea) + " Column: " +
			strconv.Itoa(errores_sintacticos.Errores[i].Columna))
	}

	antlr.ParseTreeWalkerDefault.Walk(nvisitor, tree)
	fmt.Println(nvisitor.GetConsola())

}
