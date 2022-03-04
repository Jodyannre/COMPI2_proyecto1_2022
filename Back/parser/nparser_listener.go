// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NparserListener is a complete listener for a parse tree produced by Nparser.
type NparserListener interface {
	antlr.ParseTreeListener

	// EnterInicio is called when entering the inicio production.
	EnterInicio(c *InicioContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterTipo_dato is called when entering the tipo_dato production.
	EnterTipo_dato(c *Tipo_datoContext)

	// EnterControl_if is called when entering the control_if production.
	EnterControl_if(c *Control_ifContext)

	// EnterBloque_else_if is called when entering the bloque_else_if production.
	EnterBloque_else_if(c *Bloque_else_ifContext)

	// EnterElse_if is called when entering the else_if production.
	EnterElse_if(c *Else_ifContext)

	// EnterControl_if_exp is called when entering the control_if_exp production.
	EnterControl_if_exp(c *Control_if_expContext)

	// EnterBloque_else_if_exp is called when entering the bloque_else_if_exp production.
	EnterBloque_else_if_exp(c *Bloque_else_if_expContext)

	// EnterElse_if_exp is called when entering the else_if_exp production.
	EnterElse_if_exp(c *Else_if_expContext)

	// EnterControl_expresion is called when entering the control_expresion production.
	EnterControl_expresion(c *Control_expresionContext)

	// ExitInicio is called when exiting the inicio production.
	ExitInicio(c *InicioContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitTipo_dato is called when exiting the tipo_dato production.
	ExitTipo_dato(c *Tipo_datoContext)

	// ExitControl_if is called when exiting the control_if production.
	ExitControl_if(c *Control_ifContext)

	// ExitBloque_else_if is called when exiting the bloque_else_if production.
	ExitBloque_else_if(c *Bloque_else_ifContext)

	// ExitElse_if is called when exiting the else_if production.
	ExitElse_if(c *Else_ifContext)

	// ExitControl_if_exp is called when exiting the control_if_exp production.
	ExitControl_if_exp(c *Control_if_expContext)

	// ExitBloque_else_if_exp is called when exiting the bloque_else_if_exp production.
	ExitBloque_else_if_exp(c *Bloque_else_if_expContext)

	// ExitElse_if_exp is called when exiting the else_if_exp production.
	ExitElse_if_exp(c *Else_if_expContext)

	// ExitControl_expresion is called when exiting the control_expresion production.
	ExitControl_expresion(c *Control_expresionContext)
}
