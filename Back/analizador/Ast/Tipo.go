package Ast

type TipoDato int

var ValorTipoDato = [100]string{
	"I64",
	"F64",
	"STRING_OWNED",
	"STRING",
	"&STR",
	"BOOLEAN",
	"USIZE",
	"CHAR",
	"ARRAY",
	"Vec<>",
	"STRUCT",
	"IDENTIFICADOR",
	"LOOP_EXPRESION",
	"IF_EXPRESION",
	"ELSE_EXPRESION",
	"ELSEIF_EXPRESION",
	"MATCH_EXPRESION",
	"ACCESO_VECTOR",
	"ACCESO_MODULO",
	"LLAMADA",
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
	"RETORNO",
	"EJECUTADO",
	"CASE",
	"CASE_EXPRESION",
	"DEFAULT",
	"BREAK_EXPRESION",
	"RETURN_EXPRESION",
	"PRINTF",
	"PRINT_PRIMITIVOS",
	"PRINT_ARRAY",
	"ERROR_SEMANTICO_NO",
	"LLAMADA_FUNCION",
	"PARAMETRO",
	"CAST",
	"SIMBOLO",
	"VALOR",
	"ERROR_NO_EXISTE",
	"ERROR_ACCESO_PRIVADO",
	"MUTABLE",
	"NOT_MUTABLE",
	"OCUPADO",
	"LIBRE",
	"VEC_NEW",
	"VEC_PUSH",
	"VEC_LEN",
	"VEC_CONTAINS",
	"VEC_CAPACITY",
	"VEC_INSERT",
	"VEC_REMOVE",
	"VEC_ACCESO",
	"VEC_FAC",
	"VEC_ELEMENTOS",
	"VEC_WITH_CAPACITY",
	"POW",
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
	USIZE
	CHAR
	ARRAY
	VECTOR
	STRUCT
	IDENTIFICADOR
	LOOP_EXPRESION
	IF_EXPRESION
	ELSE_EXPRESION
	ELSEIF_EXPRESION
	MATCH_EXPRESION
	ACCESO_VECTOR
	ACCESO_MODULO
	LLAMADA
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
	CASE
	CASE_EXPRESION
	DEFAULT
	BREAK_EXPRESION
	RETURN_EXPRESION
	PRINTF
	PRINT_PRIMITIVOS
	PRINT_ARRAY
	ERROR_SEMANTICO_NO
	LLAMADA_FUNCION
	PARAMETRO
	CAST
	SIMBOLO
	VALOR
	ERROR_NO_EXISTE
	ERROR_ACCESO_PRIVADO
	MUTABLE
	NOT_MUTABLE
	OCUPADO
	LIBRE
	VEC_NEW
	VEC_PUSH
	VEC_LEN
	VEC_CONTAINS
	VEC_CAPACITY
	VEC_INSERT
	VEC_REMOVE
	VEC_ACCESO
	VEC_FAC
	VEC_ELEMENTOS
	VEC_WITH_CAPACITY
	POW
	ARRAY_ELEMENTOS
	ARRAY_FAC
	DIMENSION_ARRAY
	DECLARACION_ARRAY
	ACCESO_ARRAY
	ATRIBUTO
	STRUCT_TEMPLATE
	TIPO
	ACCESO_STRUCT
	ASIGNACION_STRUCT
	FUNCION_MAIN
	TO_CHARS
	RANGE_EXPRESION
	RANGE_RANGO
)

type TipoRetornado struct {
	Tipo  TipoDato
	Valor interface{}
}

func EsTransferencia(tipo TipoDato) bool {
	if tipo == BREAK ||
		tipo == BREAK_EXPRESION ||
		tipo == RETURN ||
		tipo == RETURN_EXPRESION ||
		tipo == CONTINUE {
		return true
	} else {
		return false
	}
}

func EsPrimitivo(tipo TipoDato) bool {
	if tipo <= 6 {
		return true
	} else {
		return false
	}
}

func EsFuncion(tipo interface{}) bool {
	validador := false

	switch tipo {
	case FUNCION, VEC_NEW, VEC_ACCESO,
		VEC_LEN, VEC_CONTAINS,
		VEC_CAPACITY, VEC_REMOVE, ARRAY_FAC, ARRAY_ELEMENTOS, ARRAY,
		VEC_ELEMENTOS, VEC_FAC, VEC_WITH_CAPACITY, DIMENSION_ARRAY, LLAMADA_FUNCION:
		validador = true
	default:
		validador = false
	}

	return validador
}
