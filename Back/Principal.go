package main

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/visitantes"
	"Back/parser"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Key struct {
	Nombre string
	Padre  string
}

func main() {
	/*
		match, _ := regexp.MatchString("[{ *}|{:?}]", "peach {   } hola mundo {:?}")
		fmt.Println(match)
		formato1, _ := regexp.Compile("{ *}|{:[\x3F]}")
		formato2, _ := regexp.Compile("lala")
		cadena := "El resultado de {:?} se obtiene por sumar {   }{}, pero el {} me gusta más."

		fmt.Println(formato1.MatchString(cadena))
		fmt.Println(formato2.MatchString(cadena))

		nuevo := formato1.Split(cadena, -1)
		array := arraylist.New()
		array2 := formato1.FindAllStringIndex(cadena, -1)
		subString := cadena[42:47]
		subString = strings.Replace(subString, " ", "", -1)
		fmt.Println(subString)
		fmt.Println(array2)
		var salida string
		array.Add("5")
		array.Add("2")
		array.Add("3")
		array.Add("6")
		for i := range nuevo {
			if nuevo[i] == "" && i < 1 {
				//En el primero agrego el primer elemento
				salida += array.GetValue(i).(string)
			} else if nuevo[i] == "" && i == len(nuevo)-1 {
				//En el último no hago nada
			} else {
				if i >= array.Len() {
					salida += nuevo[i]
				} else {
					salida += nuevo[i] + array.GetValue(i).(string)
				}
			}
		}

		fmt.Println(salida)
	*/
	var input string = `
	[
	[ [ 1, 3, 5, 7], [ 5;4 ] ],
	[ [ 2, 4, 6, 8], [ 10;4 ] ],
	[ [ 2; 4 ], [ 0; 4 ] ]
	];

	println!("{:?}",[1,5,6,7,8]);
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
			strconv.Itoa(errores_lexicos.Errores[i].Fila) + " Column: " +
			strconv.Itoa(errores_lexicos.Errores[i].Columna))
	}

	//Recorrer el array de errores sintácticos
	for i := 0; i < len(errores_sintacticos.Errores); i++ {
		//fmt.Printf("Error #"+strconv.Itoa(i)+" %v \n", errores_sintacticos.errores[i])
		errores_sintacticos.Errores[i].Tipo = Ast.ERROR_SINTACTICO
		nvisitor.Errores.Add(errores_sintacticos.Errores[i])
		//Actualizar la consola
		nvisitor.UpdateConsola("Syntax error , " + errores_sintacticos.Errores[i].Msg + " -- Line: " +
			strconv.Itoa(errores_sintacticos.Errores[i].Fila) + " Column: " +
			strconv.Itoa(errores_sintacticos.Errores[i].Columna))
	}

	antlr.ParseTreeWalkerDefault.Walk(nvisitor, tree)
	fmt.Println(nvisitor.GetConsola())

}
