// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNparserListener is a complete listener for a parse tree produced by Nparser.
type BaseNparserListener struct{}

var _ NparserListener = &BaseNparserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNparserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNparserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNparserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNparserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInicio is called when production inicio is entered.
func (s *BaseNparserListener) EnterInicio(ctx *InicioContext) {}

// ExitInicio is called when production inicio is exited.
func (s *BaseNparserListener) ExitInicio(ctx *InicioContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseNparserListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseNparserListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseNparserListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseNparserListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *BaseNparserListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *BaseNparserListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseNparserListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseNparserListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaseNparserListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseNparserListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterTipo_dato is called when production tipo_dato is entered.
func (s *BaseNparserListener) EnterTipo_dato(ctx *Tipo_datoContext) {}

// ExitTipo_dato is called when production tipo_dato is exited.
func (s *BaseNparserListener) ExitTipo_dato(ctx *Tipo_datoContext) {}
