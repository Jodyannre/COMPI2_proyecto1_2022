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

// EnterBloque is called when production bloque is entered.
func (s *BaseNparserListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BaseNparserListener) ExitBloque(ctx *BloqueContext) {}

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

// EnterControl_if is called when production control_if is entered.
func (s *BaseNparserListener) EnterControl_if(ctx *Control_ifContext) {}

// ExitControl_if is called when production control_if is exited.
func (s *BaseNparserListener) ExitControl_if(ctx *Control_ifContext) {}

// EnterBloque_else_if is called when production bloque_else_if is entered.
func (s *BaseNparserListener) EnterBloque_else_if(ctx *Bloque_else_ifContext) {}

// ExitBloque_else_if is called when production bloque_else_if is exited.
func (s *BaseNparserListener) ExitBloque_else_if(ctx *Bloque_else_ifContext) {}

// EnterElse_if is called when production else_if is entered.
func (s *BaseNparserListener) EnterElse_if(ctx *Else_ifContext) {}

// ExitElse_if is called when production else_if is exited.
func (s *BaseNparserListener) ExitElse_if(ctx *Else_ifContext) {}

// EnterControl_if_exp is called when production control_if_exp is entered.
func (s *BaseNparserListener) EnterControl_if_exp(ctx *Control_if_expContext) {}

// ExitControl_if_exp is called when production control_if_exp is exited.
func (s *BaseNparserListener) ExitControl_if_exp(ctx *Control_if_expContext) {}

// EnterBloque_else_if_exp is called when production bloque_else_if_exp is entered.
func (s *BaseNparserListener) EnterBloque_else_if_exp(ctx *Bloque_else_if_expContext) {}

// ExitBloque_else_if_exp is called when production bloque_else_if_exp is exited.
func (s *BaseNparserListener) ExitBloque_else_if_exp(ctx *Bloque_else_if_expContext) {}

// EnterElse_if_exp is called when production else_if_exp is entered.
func (s *BaseNparserListener) EnterElse_if_exp(ctx *Else_if_expContext) {}

// ExitElse_if_exp is called when production else_if_exp is exited.
func (s *BaseNparserListener) ExitElse_if_exp(ctx *Else_if_expContext) {}

// EnterControl_expresion is called when production control_expresion is entered.
func (s *BaseNparserListener) EnterControl_expresion(ctx *Control_expresionContext) {}

// ExitControl_expresion is called when production control_expresion is exited.
func (s *BaseNparserListener) ExitControl_expresion(ctx *Control_expresionContext) {}
