package Ast

type TipoDato int

var ValorTipoDato = [31]string{
	"i64",
	"f64",
	"String",
	"&str",
	"bool",
	"char",
	"id",
	"usize",
	"arreglo",
	"vector",
	"struct",
	"modulo",
	"funcion",
	"return",
	"continue",
	"break",
	"loop",
	"while",
	"for",
	"match",
	"declaración",
	"asignación",
	"llamada",
	"if",
	"else",
	"elseif",
	"print",
	"expresión",
	"instrucción",
	"primitivo",
	"null",
}

const (
	I64 TipoDato = iota
	F64
	STRING_OWNED
	STRING
	STR
	BOOLEAN
	CHAR
	IDENTIFICADOR
	LOOP_EXPRESION
	IF_EXPRESION
	ELSE_EXPRESION
	ELSEIF_EXPRESION
	MATCH_EXPRESION
	ACCESO_VECTOR
	ACCESO_MODULO
	LLAMADA
	USIZE
	ARRAY
	VECTOR
	STRUCT
	MODULO
	FUNCION
	RETURN
	CONTINUE
	BREAK
	LOOP
	WHILE
	FOR
	MATCH
	DECLARACION
	ASIGNACION
	IF
	ELSE
	ELSEIF
	PRINT
	EXPRESION
	INSTRUCCION
	PRIMITIVO
	NULL
	ERROR
	ERROR_SEMANTICO
	ERROR_LEXICO
	ERROR_SINTACTICO
)

type TipoRetornado struct {
	Tipo  TipoDato
	Valor interface{}
}
