// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "github.com/colegno/arraylist"
import "Back/analizador/Ast"
import "Back/analizador/expresiones"
import "Back/analizador/instrucciones"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 91, 162,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 21, 10, 3, 12, 3, 14, 3, 24,
	11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 39, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 78, 10, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 5, 7, 109, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7,
	141, 10, 7, 12, 7, 14, 7, 144, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 160, 10, 8, 3,
	8, 2, 3, 12, 9, 2, 4, 6, 8, 10, 12, 14, 2, 7, 4, 2, 76, 76, 78, 78, 3,
	2, 72, 74, 3, 2, 76, 77, 4, 2, 64, 68, 70, 70, 4, 2, 68, 68, 70, 70, 2,
	181, 2, 16, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 38, 3, 2, 2, 2, 8, 77, 3,
	2, 2, 2, 10, 79, 3, 2, 2, 2, 12, 108, 3, 2, 2, 2, 14, 159, 3, 2, 2, 2,
	16, 17, 5, 4, 3, 2, 17, 18, 8, 2, 1, 2, 18, 3, 3, 2, 2, 2, 19, 21, 5, 6,
	4, 2, 20, 19, 3, 2, 2, 2, 21, 24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23,
	3, 2, 2, 2, 23, 25, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 25, 26, 8, 3, 1, 2,
	26, 5, 3, 2, 2, 2, 27, 28, 5, 12, 7, 2, 28, 29, 8, 4, 1, 2, 29, 39, 3,
	2, 2, 2, 30, 31, 5, 8, 5, 2, 31, 32, 7, 63, 2, 2, 32, 33, 8, 4, 1, 2, 33,
	39, 3, 2, 2, 2, 34, 35, 5, 10, 6, 2, 35, 36, 7, 63, 2, 2, 36, 37, 8, 4,
	1, 2, 37, 39, 3, 2, 2, 2, 38, 27, 3, 2, 2, 2, 38, 30, 3, 2, 2, 2, 38, 34,
	3, 2, 2, 2, 39, 7, 3, 2, 2, 2, 40, 41, 7, 17, 2, 2, 41, 42, 7, 51, 2, 2,
	42, 43, 7, 71, 2, 2, 43, 44, 5, 12, 7, 2, 44, 45, 8, 5, 1, 2, 45, 78, 3,
	2, 2, 2, 46, 47, 7, 17, 2, 2, 47, 48, 7, 51, 2, 2, 48, 49, 7, 59, 2, 2,
	49, 50, 5, 14, 8, 2, 50, 51, 7, 71, 2, 2, 51, 52, 5, 12, 7, 2, 52, 53,
	8, 5, 1, 2, 53, 78, 3, 2, 2, 2, 54, 55, 7, 17, 2, 2, 55, 56, 7, 16, 2,
	2, 56, 57, 7, 51, 2, 2, 57, 58, 7, 71, 2, 2, 58, 59, 5, 12, 7, 2, 59, 60,
	8, 5, 1, 2, 60, 78, 3, 2, 2, 2, 61, 62, 7, 17, 2, 2, 62, 63, 7, 16, 2,
	2, 63, 64, 7, 51, 2, 2, 64, 65, 7, 59, 2, 2, 65, 66, 5, 14, 8, 2, 66, 67,
	8, 5, 1, 2, 67, 78, 3, 2, 2, 2, 68, 69, 7, 17, 2, 2, 69, 70, 7, 16, 2,
	2, 70, 71, 7, 51, 2, 2, 71, 72, 7, 59, 2, 2, 72, 73, 5, 14, 8, 2, 73, 74,
	7, 71, 2, 2, 74, 75, 5, 12, 7, 2, 75, 76, 8, 5, 1, 2, 76, 78, 3, 2, 2,
	2, 77, 40, 3, 2, 2, 2, 77, 46, 3, 2, 2, 2, 77, 54, 3, 2, 2, 2, 77, 61,
	3, 2, 2, 2, 77, 68, 3, 2, 2, 2, 78, 9, 3, 2, 2, 2, 79, 80, 7, 51, 2, 2,
	80, 81, 7, 71, 2, 2, 81, 82, 5, 12, 7, 2, 82, 83, 8, 6, 1, 2, 83, 11, 3,
	2, 2, 2, 84, 85, 8, 7, 1, 2, 85, 86, 9, 2, 2, 2, 86, 87, 5, 12, 7, 17,
	87, 88, 8, 7, 1, 2, 88, 109, 3, 2, 2, 2, 89, 90, 7, 80, 2, 2, 90, 91, 5,
	12, 7, 2, 91, 92, 7, 81, 2, 2, 92, 93, 8, 7, 1, 2, 93, 109, 3, 2, 2, 2,
	94, 95, 7, 51, 2, 2, 95, 109, 8, 7, 1, 2, 96, 97, 7, 21, 2, 2, 97, 109,
	8, 7, 1, 2, 98, 99, 7, 22, 2, 2, 99, 109, 8, 7, 1, 2, 100, 101, 7, 88,
	2, 2, 101, 109, 8, 7, 1, 2, 102, 103, 7, 49, 2, 2, 103, 109, 8, 7, 1, 2,
	104, 105, 7, 48, 2, 2, 105, 109, 8, 7, 1, 2, 106, 107, 7, 86, 2, 2, 107,
	109, 8, 7, 1, 2, 108, 84, 3, 2, 2, 2, 108, 89, 3, 2, 2, 2, 108, 94, 3,
	2, 2, 2, 108, 96, 3, 2, 2, 2, 108, 98, 3, 2, 2, 2, 108, 100, 3, 2, 2, 2,
	108, 102, 3, 2, 2, 2, 108, 104, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109,
	142, 3, 2, 2, 2, 110, 111, 12, 16, 2, 2, 111, 112, 9, 3, 2, 2, 112, 113,
	5, 12, 7, 17, 113, 114, 8, 7, 1, 2, 114, 141, 3, 2, 2, 2, 115, 116, 12,
	15, 2, 2, 116, 117, 9, 4, 2, 2, 117, 118, 5, 12, 7, 16, 118, 119, 8, 7,
	1, 2, 119, 141, 3, 2, 2, 2, 120, 121, 12, 14, 2, 2, 121, 122, 9, 5, 2,
	2, 122, 123, 5, 12, 7, 15, 123, 124, 8, 7, 1, 2, 124, 141, 3, 2, 2, 2,
	125, 126, 12, 13, 2, 2, 126, 127, 9, 6, 2, 2, 127, 128, 5, 12, 7, 14, 128,
	129, 8, 7, 1, 2, 129, 141, 3, 2, 2, 2, 130, 131, 12, 12, 2, 2, 131, 132,
	7, 56, 2, 2, 132, 133, 5, 12, 7, 13, 133, 134, 8, 7, 1, 2, 134, 141, 3,
	2, 2, 2, 135, 136, 12, 11, 2, 2, 136, 137, 7, 54, 2, 2, 137, 138, 5, 12,
	7, 12, 138, 139, 8, 7, 1, 2, 139, 141, 3, 2, 2, 2, 140, 110, 3, 2, 2, 2,
	140, 115, 3, 2, 2, 2, 140, 120, 3, 2, 2, 2, 140, 125, 3, 2, 2, 2, 140,
	130, 3, 2, 2, 2, 140, 135, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 140,
	3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 13, 3, 2, 2, 2, 144, 142, 3, 2,
	2, 2, 145, 146, 7, 3, 2, 2, 146, 160, 8, 8, 1, 2, 147, 148, 7, 4, 2, 2,
	148, 160, 8, 8, 1, 2, 149, 150, 7, 6, 2, 2, 150, 160, 8, 8, 1, 2, 151,
	152, 7, 5, 2, 2, 152, 160, 8, 8, 1, 2, 153, 154, 7, 7, 2, 2, 154, 160,
	8, 8, 1, 2, 155, 156, 7, 8, 2, 2, 156, 160, 8, 8, 1, 2, 157, 158, 7, 9,
	2, 2, 158, 160, 8, 8, 1, 2, 159, 145, 3, 2, 2, 2, 159, 147, 3, 2, 2, 2,
	159, 149, 3, 2, 2, 2, 159, 151, 3, 2, 2, 2, 159, 153, 3, 2, 2, 2, 159,
	155, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 160, 15, 3, 2, 2, 2, 9, 22, 38,
	77, 108, 140, 142, 159,
}
var literalNames = []string{
	"", "'bool'", "'char'", "'f64'", "'i64'", "'&str'", "'String'", "'usize'",
	"'main'", "'powf'", "'pow'", "'as'", "'Vec'", "'vec'", "'mut'", "'let'",
	"'struct'", "'to_string'", "'to_owned'", "'true'", "'false'", "'println'",
	"'fn'", "'abs'", "'sqrt'", "'clone'", "'new'", "'len'", "'push'", "'remove'",
	"'contains'", "'insert'", "'capacity'", "'with_capacity'", "'if'", "'else'",
	"'match'", "'loop'", "'while'", "'for'", "'in'", "'return'", "'break'",
	"'continue'", "'mod'", "'pub'", "", "", "", "", "'_'", "'|'", "'||'", "'&'",
	"'&&'", "':?'", "'::'", "':'", "'..'", "'.'", "','", "';'", "'>='", "'>'",
	"'<='", "'<'", "'=='", "'=>'", "'!='", "'='", "'%'", "'*'", "'/'", "'->'",
	"'-'", "'+'", "'!'", "'?'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "BOOL", "CHAR", "F64", "I64", "STR", "STRING", "USIZE", "MAIN", "POWF",
	"POW", "AS", "VEC", "VEC_M", "MUT", "LET", "STRUCT", "TO_STRING", "TO_OWNED",
	"TRUE", "FALSE", "PRINT", "FN", "ABS", "SQRT", "CLONE", "NEW", "LEN", "PUSH",
	"REMOVE", "CONTAINS", "INSERT", "CAPACITY", "WITH_CAPACITY", "IF", "ELSE",
	"MATCH", "LOOP", "WHILE", "FOR", "IN", "RETURN", "BREAK", "CONTINUE", "MOD",
	"PUB", "NUMERO", "DECIMAL", "ID_CAMEL", "ID", "DEFAULT", "O", "OR", "AMPERSAND",
	"AND", "PRINT_OP_DEBUG", "DOBLE_DOSPUNTOS", "DOSPUNTOS", "RANGO", "PUNTO",
	"COMA", "PUNTOCOMA", "MAYOR_I", "MAYOR", "MENOR_I", "MENOR", "IGUALDAD",
	"CASE", "DISTINTO", "IGUAL", "MODULO", "MULTIPLICACION", "DIVISION", "FN_TIPO_RETORNO",
	"RESTA", "SUMA", "NOT", "PREGUNTA", "PAR_IZQ", "PAR_DER", "LLAVE_IZQ",
	"LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", "CADENA", "ASCII", "CARACTER",
	"WHITESPACE", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"inicio", "instrucciones", "instruccion", "declaracion", "asignacion",
	"expresion", "tipo_dato",
}

type Nparser struct {
	*antlr.BaseParser
}

// NewNparser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Nparser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNparser(input antlr.TokenStream) *Nparser {
	this := new(Nparser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Nparser.g4"

	return this
}

// Nparser tokens.
const (
	NparserEOF             = antlr.TokenEOF
	NparserBOOL            = 1
	NparserCHAR            = 2
	NparserF64             = 3
	NparserI64             = 4
	NparserSTR             = 5
	NparserSTRING          = 6
	NparserUSIZE           = 7
	NparserMAIN            = 8
	NparserPOWF            = 9
	NparserPOW             = 10
	NparserAS              = 11
	NparserVEC             = 12
	NparserVEC_M           = 13
	NparserMUT             = 14
	NparserLET             = 15
	NparserSTRUCT          = 16
	NparserTO_STRING       = 17
	NparserTO_OWNED        = 18
	NparserTRUE            = 19
	NparserFALSE           = 20
	NparserPRINT           = 21
	NparserFN              = 22
	NparserABS             = 23
	NparserSQRT            = 24
	NparserCLONE           = 25
	NparserNEW             = 26
	NparserLEN             = 27
	NparserPUSH            = 28
	NparserREMOVE          = 29
	NparserCONTAINS        = 30
	NparserINSERT          = 31
	NparserCAPACITY        = 32
	NparserWITH_CAPACITY   = 33
	NparserIF              = 34
	NparserELSE            = 35
	NparserMATCH           = 36
	NparserLOOP            = 37
	NparserWHILE           = 38
	NparserFOR             = 39
	NparserIN              = 40
	NparserRETURN          = 41
	NparserBREAK           = 42
	NparserCONTINUE        = 43
	NparserMOD             = 44
	NparserPUB             = 45
	NparserNUMERO          = 46
	NparserDECIMAL         = 47
	NparserID_CAMEL        = 48
	NparserID              = 49
	NparserDEFAULT         = 50
	NparserO               = 51
	NparserOR              = 52
	NparserAMPERSAND       = 53
	NparserAND             = 54
	NparserPRINT_OP_DEBUG  = 55
	NparserDOBLE_DOSPUNTOS = 56
	NparserDOSPUNTOS       = 57
	NparserRANGO           = 58
	NparserPUNTO           = 59
	NparserCOMA            = 60
	NparserPUNTOCOMA       = 61
	NparserMAYOR_I         = 62
	NparserMAYOR           = 63
	NparserMENOR_I         = 64
	NparserMENOR           = 65
	NparserIGUALDAD        = 66
	NparserCASE            = 67
	NparserDISTINTO        = 68
	NparserIGUAL           = 69
	NparserMODULO          = 70
	NparserMULTIPLICACION  = 71
	NparserDIVISION        = 72
	NparserFN_TIPO_RETORNO = 73
	NparserRESTA           = 74
	NparserSUMA            = 75
	NparserNOT             = 76
	NparserPREGUNTA        = 77
	NparserPAR_IZQ         = 78
	NparserPAR_DER         = 79
	NparserLLAVE_IZQ       = 80
	NparserLLAVE_DER       = 81
	NparserCORCHETE_IZQ    = 82
	NparserCORCHETE_DER    = 83
	NparserCADENA          = 84
	NparserASCII           = 85
	NparserCARACTER        = 86
	NparserWHITESPACE      = 87
	NparserCOMMENT         = 88
	NparserLINE_COMMENT    = 89
)

// Nparser rules.
const (
	NparserRULE_inicio        = 0
	NparserRULE_instrucciones = 1
	NparserRULE_instruccion   = 2
	NparserRULE_declaracion   = 3
	NparserRULE_asignacion    = 4
	NparserRULE_expresion     = 5
	NparserRULE_tipo_dato     = 6
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arraylist.List

	// SetLista sets the lista attribute.
	SetLista(*arraylist.List)

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arraylist.List
	_instrucciones IInstruccionesContext
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_inicio
	return p
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *InicioContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *InicioContext) GetLista() *arraylist.List { return s.lista }

func (s *InicioContext) SetLista(v *arraylist.List) { s.lista = v }

func (s *InicioContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterInicio(s)
	}
}

func (s *InicioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitInicio(s)
	}
}

func (p *Nparser) Inicio() (localctx IInicioContext) {
	this := p
	_ = this

	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NparserRULE_inicio)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)

		var _x = p.Instrucciones()

		localctx.(*InicioContext)._instrucciones = _x
	}

	localctx.(*InicioContext).lista = localctx.(*InicioContext).Get_instrucciones().GetList()

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	list         *arraylist.List
	_instruccion IInstruccionContext
	e            []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) GetList() *arraylist.List { return s.list }

func (s *InstruccionesContext) SetList(v *arraylist.List) { s.list = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Nparser) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NparserRULE_instrucciones)

	localctx.(*InstruccionesContext).list = arraylist.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NparserLET)|(1<<NparserTRUE)|(1<<NparserFALSE))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(NparserNUMERO-46))|(1<<(NparserDECIMAL-46))|(1<<(NparserID-46))|(1<<(NparserRESTA-46))|(1<<(NparserNOT-46)))) != 0) || (((_la-78)&-(0x1f+1)) == 0 && ((1<<uint((_la-78)))&((1<<(NparserPAR_IZQ-78))|(1<<(NparserCADENA-78))|(1<<(NparserCARACTER-78)))) != 0) {
		{
			p.SetState(17)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).list.Add(e.GetEx())
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// GetEx returns the ex attribute.
	GetEx() interface{}

	// SetEx sets the ex attribute.
	SetEx(interface{})

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	ex           interface{}
	_expresion   IExpresionContext
	_declaracion IDeclaracionContext
	_asignacion  IAsignacionContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) GetEx() interface{} { return s.ex }

func (s *InstruccionContext) SetEx(v interface{}) { s.ex = v }

func (s *InstruccionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(NparserPUNTOCOMA, 0)
}

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Nparser) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NparserRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)

			var _x = p.expresion(0)

			localctx.(*InstruccionContext)._expresion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_expresion().GetEx()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		{
			p.SetState(29)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_declaracion().GetEx()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		{
			p.SetState(33)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_asignacion().GetEx()

	}

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_tipo_dato returns the _tipo_dato rule contexts.
	Get_tipo_dato() ITipo_datoContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_tipo_dato sets the _tipo_dato rule contexts.
	Set_tipo_dato(ITipo_datoContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_LET       antlr.Token
	_ID        antlr.Token
	_expresion IExpresionContext
	_tipo_dato ITipo_datoContext
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Get_LET() antlr.Token { return s._LET }

func (s *DeclaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclaracionContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *DeclaracionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclaracionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *DeclaracionContext) Get_tipo_dato() ITipo_datoContext { return s._tipo_dato }

func (s *DeclaracionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *DeclaracionContext) Set_tipo_dato(v ITipo_datoContext) { s._tipo_dato = v }

func (s *DeclaracionContext) GetEx() Ast.Instruccion { return s.ex }

func (s *DeclaracionContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *DeclaracionContext) LET() antlr.TerminalNode {
	return s.GetToken(NparserLET, 0)
}

func (s *DeclaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *DeclaracionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(NparserIGUAL, 0)
}

func (s *DeclaracionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclaracionContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(NparserDOSPUNTOS, 0)
}

func (s *DeclaracionContext) Tipo_dato() ITipo_datoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_datoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_datoContext)
}

func (s *DeclaracionContext) MUT() antlr.TerminalNode {
	return s.GetToken(NparserMUT, 0)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *Nparser) Declaracion() (localctx IDeclaracionContext) {
	this := p
	_ = this

	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NparserRULE_declaracion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(39)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(40)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(41)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), Ast.INDEFINIDO,
			false, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(45)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(46)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(47)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(48)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(49)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(),
			false, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(53)
			p.Match(NparserMUT)
		}
		{
			p.SetState(54)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(55)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(56)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), Ast.INDEFINIDO,
			true, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(59)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(60)
			p.Match(NparserMUT)
		}
		{
			p.SetState(61)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(62)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(63)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}

		valor := expresiones.NewPrimitivo(nil, Ast.NULL)
		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(),
			true, false, Ast.VOID, valor, fila, columna)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(66)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(67)
			p.Match(NparserMUT)
		}
		{
			p.SetState(68)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(69)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(70)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(71)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(72)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(),
			true, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	}

	return localctx
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_ID        antlr.Token
	_expresion IExpresionContext
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_asignacion
	return p
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *AsignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AsignacionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *AsignacionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *AsignacionContext) GetEx() Ast.Instruccion { return s.ex }

func (s *AsignacionContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *AsignacionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(NparserIGUAL, 0)
}

func (s *AsignacionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *Nparser) Asignacion() (localctx IAsignacionContext) {
	this := p
	_ = this

	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NparserRULE_asignacion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)

		var _m = p.Match(NparserID)

		localctx.(*AsignacionContext)._ID = _m
	}
	{
		p.SetState(78)
		p.Match(NparserIGUAL)
	}
	{
		p.SetState(79)

		var _x = p.expresion(0)

		localctx.(*AsignacionContext)._expresion = _x
	}

	fila := (func() int {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetColumn()
		}
	}())
	localctx.(*AsignacionContext).ex = instrucciones.NewAsignacion((func() string {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetText()
		}
	}()), localctx.(*AsignacionContext).Get_expresion().GetEx(), fila, columna)

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_CARACTER returns the _CARACTER token.
	Get_CARACTER() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_AND returns the _AND token.
	Get_AND() antlr.Token

	// Get_OR returns the _OR token.
	Get_OR() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_CARACTER sets the _CARACTER token.
	Set_CARACTER(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_AND sets the _AND token.
	Set_AND(antlr.Token)

	// Set_OR sets the _OR token.
	Set_OR(antlr.Token)

	// GetOp_izq returns the op_izq rule contexts.
	GetOp_izq() IExpresionContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// GetOp_der returns the op_der rule contexts.
	GetOp_der() IExpresionContext

	// SetOp_izq sets the op_izq rule contexts.
	SetOp_izq(IExpresionContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// SetOp_der sets the op_der rule contexts.
	SetOp_der(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Expresion

	// SetEx sets the ex attribute.
	SetEx(Ast.Expresion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Expresion
	op_izq     IExpresionContext
	op         antlr.Token
	_expresion IExpresionContext
	_ID        antlr.Token
	_CARACTER  antlr.Token
	_DECIMAL   antlr.Token
	_NUMERO    antlr.Token
	_CADENA    antlr.Token
	op_der     IExpresionContext
	_AND       antlr.Token
	_OR        antlr.Token
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) GetOp() antlr.Token { return s.op }

func (s *ExpresionContext) Get_ID() antlr.Token { return s._ID }

func (s *ExpresionContext) Get_CARACTER() antlr.Token { return s._CARACTER }

func (s *ExpresionContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *ExpresionContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *ExpresionContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ExpresionContext) Get_AND() antlr.Token { return s._AND }

func (s *ExpresionContext) Get_OR() antlr.Token { return s._OR }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExpresionContext) Set_CARACTER(v antlr.Token) { s._CARACTER = v }

func (s *ExpresionContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *ExpresionContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *ExpresionContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ExpresionContext) Set_AND(v antlr.Token) { s._AND = v }

func (s *ExpresionContext) Set_OR(v antlr.Token) { s._OR = v }

func (s *ExpresionContext) GetOp_izq() IExpresionContext { return s.op_izq }

func (s *ExpresionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ExpresionContext) GetOp_der() IExpresionContext { return s.op_der }

func (s *ExpresionContext) SetOp_izq(v IExpresionContext) { s.op_izq = v }

func (s *ExpresionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ExpresionContext) SetOp_der(v IExpresionContext) { s.op_der = v }

func (s *ExpresionContext) GetEx() Ast.Expresion { return s.ex }

func (s *ExpresionContext) SetEx(v Ast.Expresion) { s.ex = v }

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) RESTA() antlr.TerminalNode {
	return s.GetToken(NparserRESTA, 0)
}

func (s *ExpresionContext) NOT() antlr.TerminalNode {
	return s.GetToken(NparserNOT, 0)
}

func (s *ExpresionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *ExpresionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *ExpresionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *ExpresionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NparserTRUE, 0)
}

func (s *ExpresionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NparserFALSE, 0)
}

func (s *ExpresionContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(NparserCARACTER, 0)
}

func (s *ExpresionContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(NparserDECIMAL, 0)
}

func (s *ExpresionContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(NparserNUMERO, 0)
}

func (s *ExpresionContext) CADENA() antlr.TerminalNode {
	return s.GetToken(NparserCADENA, 0)
}

func (s *ExpresionContext) MULTIPLICACION() antlr.TerminalNode {
	return s.GetToken(NparserMULTIPLICACION, 0)
}

func (s *ExpresionContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(NparserDIVISION, 0)
}

func (s *ExpresionContext) MODULO() antlr.TerminalNode {
	return s.GetToken(NparserMODULO, 0)
}

func (s *ExpresionContext) SUMA() antlr.TerminalNode {
	return s.GetToken(NparserSUMA, 0)
}

func (s *ExpresionContext) MAYOR_I() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR_I, 0)
}

func (s *ExpresionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR, 0)
}

func (s *ExpresionContext) MENOR_I() antlr.TerminalNode {
	return s.GetToken(NparserMENOR_I, 0)
}

func (s *ExpresionContext) MENOR() antlr.TerminalNode {
	return s.GetToken(NparserMENOR, 0)
}

func (s *ExpresionContext) IGUALDAD() antlr.TerminalNode {
	return s.GetToken(NparserIGUALDAD, 0)
}

func (s *ExpresionContext) DISTINTO() antlr.TerminalNode {
	return s.GetToken(NparserDISTINTO, 0)
}

func (s *ExpresionContext) AND() antlr.TerminalNode {
	return s.GetToken(NparserAND, 0)
}

func (s *ExpresionContext) OR() antlr.TerminalNode {
	return s.GetToken(NparserOR, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *Nparser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *Nparser) expresion(_p int) (localctx IExpresionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, NparserRULE_expresion, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserRESTA, NparserNOT:
		{
			p.SetState(83)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpresionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == NparserRESTA || _la == NparserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpresionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(84)

			var _x = p.expresion(15)

			localctx.(*ExpresionContext).op_izq = _x
			localctx.(*ExpresionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).GetOp().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).GetOp().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), nil, true, fila, columna)

	case NparserPAR_IZQ:
		{
			p.SetState(87)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(88)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext)._expresion = _x
		}
		{
			p.SetState(89)
			p.Match(NparserPAR_DER)
		}

		localctx.(*ExpresionContext).ex = localctx.(*ExpresionContext).Get_expresion().GetEx()

	case NparserID:
		{
			p.SetState(92)

			var _m = p.Match(NparserID)

			localctx.(*ExpresionContext)._ID = _m
		}

		id := (func() string {
			if localctx.(*ExpresionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_ID().GetText()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR)

	case NparserTRUE:
		{
			p.SetState(94)
			p.Match(NparserTRUE)
		}

		valor := true
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN)

	case NparserFALSE:
		{
			p.SetState(96)
			p.Match(NparserFALSE)
		}

		valor := false
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN)

	case NparserCARACTER:
		{
			p.SetState(98)

			var _m = p.Match(NparserCARACTER)

			localctx.(*ExpresionContext)._CARACTER = _m
		}

		valor := (func() string {
			if localctx.(*ExpresionContext).Get_CARACTER() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_CARACTER().GetText()
			}
		}())
		valor = valor[1 : len(valor)-1]
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.CHAR)

	case NparserDECIMAL:
		{
			p.SetState(100)

			var _m = p.Match(NparserDECIMAL)

			localctx.(*ExpresionContext)._DECIMAL = _m
		}

		valor, err := strconv.ParseFloat((func() string {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println("Hay un error en el get número")
			fmt.Println(err)
		}
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.F64)

	case NparserNUMERO:
		{
			p.SetState(102)

			var _m = p.Match(NparserNUMERO)

			localctx.(*ExpresionContext)._NUMERO = _m
		}

		valor, err := strconv.Atoi((func() string {
			if localctx.(*ExpresionContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_NUMERO().GetText()
			}
		}()))
		if err != nil {
			fmt.Println("Hay un error en el get número")
			fmt.Println(err)
		}
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.I64)

	case NparserCADENA:
		{
			p.SetState(104)

			var _m = p.Match(NparserCADENA)

			localctx.(*ExpresionContext)._CADENA = _m
		}

		valor := (func() string {
			if localctx.(*ExpresionContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_CADENA().GetText()
			}
		}())
		valor = valor[1 : len(valor)-1]
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.STR)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(109)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-70)&-(0x1f+1)) == 0 && ((1<<uint((_la-70)))&((1<<(NparserMODULO-70))|(1<<(NparserMULTIPLICACION-70))|(1<<(NparserDIVISION-70)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(110)

					var _x = p.expresion(15)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(114)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NparserRESTA || _la == NparserSUMA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(115)

					var _x = p.expresion(14)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(118)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(119)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-62)&-(0x1f+1)) == 0 && ((1<<uint((_la-62)))&((1<<(NparserMAYOR_I-62))|(1<<(NparserMAYOR-62))|(1<<(NparserMENOR_I-62))|(1<<(NparserMENOR-62))|(1<<(NparserIGUALDAD-62))|(1<<(NparserDISTINTO-62)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(120)

					var _x = p.expresion(13)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(123)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(124)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NparserIGUALDAD || _la == NparserDISTINTO) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(125)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(129)

					var _m = p.Match(NparserAND)

					localctx.(*ExpresionContext)._AND = _m
				}
				{
					p.SetState(130)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).Get_AND() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_AND().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).Get_AND() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_AND().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).Get_AND() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).Get_AND().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(133)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(134)

					var _m = p.Match(NparserOR)

					localctx.(*ExpresionContext)._OR = _m
				}
				{
					p.SetState(135)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).Get_OR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_OR().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).Get_OR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_OR().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).Get_OR() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).Get_OR().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// ITipo_datoContext is an interface to support dynamic dispatch.
type ITipo_datoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEx returns the ex attribute.
	GetEx() Ast.TipoDato

	// SetEx sets the ex attribute.
	SetEx(Ast.TipoDato)

	// IsTipo_datoContext differentiates from other interfaces.
	IsTipo_datoContext()
}

type Tipo_datoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ex     Ast.TipoDato
}

func NewEmptyTipo_datoContext() *Tipo_datoContext {
	var p = new(Tipo_datoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_tipo_dato
	return p
}

func (*Tipo_datoContext) IsTipo_datoContext() {}

func NewTipo_datoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_datoContext {
	var p = new(Tipo_datoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_tipo_dato

	return p
}

func (s *Tipo_datoContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_datoContext) GetEx() Ast.TipoDato { return s.ex }

func (s *Tipo_datoContext) SetEx(v Ast.TipoDato) { s.ex = v }

func (s *Tipo_datoContext) BOOL() antlr.TerminalNode {
	return s.GetToken(NparserBOOL, 0)
}

func (s *Tipo_datoContext) CHAR() antlr.TerminalNode {
	return s.GetToken(NparserCHAR, 0)
}

func (s *Tipo_datoContext) I64() antlr.TerminalNode {
	return s.GetToken(NparserI64, 0)
}

func (s *Tipo_datoContext) F64() antlr.TerminalNode {
	return s.GetToken(NparserF64, 0)
}

func (s *Tipo_datoContext) STR() antlr.TerminalNode {
	return s.GetToken(NparserSTR, 0)
}

func (s *Tipo_datoContext) STRING() antlr.TerminalNode {
	return s.GetToken(NparserSTRING, 0)
}

func (s *Tipo_datoContext) USIZE() antlr.TerminalNode {
	return s.GetToken(NparserUSIZE, 0)
}

func (s *Tipo_datoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_datoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_datoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterTipo_dato(s)
	}
}

func (s *Tipo_datoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitTipo_dato(s)
	}
}

func (p *Nparser) Tipo_dato() (localctx ITipo_datoContext) {
	this := p
	_ = this

	localctx = NewTipo_datoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NparserRULE_tipo_dato)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Match(NparserBOOL)
		}
		localctx.(*Tipo_datoContext).ex = Ast.BOOLEAN

	case NparserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(NparserCHAR)
		}
		localctx.(*Tipo_datoContext).ex = Ast.CHAR

	case NparserI64:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(147)
			p.Match(NparserI64)
		}
		localctx.(*Tipo_datoContext).ex = Ast.I64

	case NparserF64:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(149)
			p.Match(NparserF64)
		}
		localctx.(*Tipo_datoContext).ex = Ast.F64

	case NparserSTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(151)
			p.Match(NparserSTR)
		}
		localctx.(*Tipo_datoContext).ex = Ast.STR

	case NparserSTRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(153)
			p.Match(NparserSTRING)
		}
		localctx.(*Tipo_datoContext).ex = Ast.STRING

	case NparserUSIZE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(155)
			p.Match(NparserUSIZE)
		}
		localctx.(*Tipo_datoContext).ex = Ast.USIZE

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Nparser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Nparser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
