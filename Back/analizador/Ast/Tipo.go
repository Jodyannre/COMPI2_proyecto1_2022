package Ast

type TipoDato int

var ValorTipoDato = [50]string{
	"I64",
	"F64",
	"STRING_OWNED",
	"STRING",
	"STR",
	"BOOLEAN",
	"CHAR",
	"IDENTIFICADOR",
	"LOOP_EXPRESION",
	"IF_EXPRESION",
	"ELSE_EXPRESION",
	"ELSEIF_EXPRESION",
	"MATCH_EXPRESION",
	"ACCESO_VECTOR",
	"ACCESO_MODULO",
	"LLAMADA",
	"USIZE",
	"ARRAY",
	"VECTOR",
	"STRUCT",
	"MODULO",
	"FUNCION",
	"RETURN",
	"CONTINUE",
	"BREAK",
	"LOOP",
	"WHILE",
	"FOR",
	"MATCH",
	"DECLARACION_VARIABLE",
	"DECLARACION_FUNCION",
	"DECLARACION_MODULO",
	"DECLARACION_VECTOR",
	"DECLARACION",
	"ASIGNACION",
	"IF",
	"ELSE",
	"ELSEIF",
	"PRINT",
	"EXPRESION",
	"INSTRUCCION",
	"PRIMITIVO",
	"NULL",
	"ERROR",
	"ERROR_SEMANTICO",
	"ERROR_LEXICO",
	"ERROR_SINTACTICO",
	"VARIABLE",
	"INDEFINIDO",
	"VOID",
}

type Key struct {
	Padre  string
	Nombre string
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
	DECLARACION_VARIABLE
	DECLARACION_FUNCION
	DECLARACION_MODULO
	DECLARACION_VECTOR
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
	VARIABLE
	INDEFINIDO
	VOID
	RETORNO
	EJECUTADO
)

type TipoRetornado struct {
	Tipo  TipoDato
	Valor interface{}
}
