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

// EnterDeclaracion_funcion is called when production declaracion_funcion is entered.
func (s *BaseNparserListener) EnterDeclaracion_funcion(ctx *Declaracion_funcionContext) {}

// ExitDeclaracion_funcion is called when production declaracion_funcion is exited.
func (s *BaseNparserListener) ExitDeclaracion_funcion(ctx *Declaracion_funcionContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseNparserListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseNparserListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterAccesos_vector_array_asignacion is called when production accesos_vector_array_asignacion is entered.
func (s *BaseNparserListener) EnterAccesos_vector_array_asignacion(ctx *Accesos_vector_array_asignacionContext) {
}

// ExitAccesos_vector_array_asignacion is called when production accesos_vector_array_asignacion is exited.
func (s *BaseNparserListener) ExitAccesos_vector_array_asignacion(ctx *Accesos_vector_array_asignacionContext) {
}

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

// EnterControl_match is called when production control_match is entered.
func (s *BaseNparserListener) EnterControl_match(ctx *Control_matchContext) {}

// ExitControl_match is called when production control_match is exited.
func (s *BaseNparserListener) ExitControl_match(ctx *Control_matchContext) {}

// EnterControl_case is called when production control_case is entered.
func (s *BaseNparserListener) EnterControl_case(ctx *Control_caseContext) {}

// ExitControl_case is called when production control_case is exited.
func (s *BaseNparserListener) ExitControl_case(ctx *Control_caseContext) {}

// EnterCases is called when production cases is entered.
func (s *BaseNparserListener) EnterCases(ctx *CasesContext) {}

// ExitCases is called when production cases is exited.
func (s *BaseNparserListener) ExitCases(ctx *CasesContext) {}

// EnterCase_match is called when production case_match is entered.
func (s *BaseNparserListener) EnterCase_match(ctx *Case_matchContext) {}

// ExitCase_match is called when production case_match is exited.
func (s *BaseNparserListener) ExitCase_match(ctx *Case_matchContext) {}

// EnterControl_match_exp is called when production control_match_exp is entered.
func (s *BaseNparserListener) EnterControl_match_exp(ctx *Control_match_expContext) {}

// ExitControl_match_exp is called when production control_match_exp is exited.
func (s *BaseNparserListener) ExitControl_match_exp(ctx *Control_match_expContext) {}

// EnterControl_case_exp is called when production control_case_exp is entered.
func (s *BaseNparserListener) EnterControl_case_exp(ctx *Control_case_expContext) {}

// ExitControl_case_exp is called when production control_case_exp is exited.
func (s *BaseNparserListener) ExitControl_case_exp(ctx *Control_case_expContext) {}

// EnterCases_exp is called when production cases_exp is entered.
func (s *BaseNparserListener) EnterCases_exp(ctx *Cases_expContext) {}

// ExitCases_exp is called when production cases_exp is exited.
func (s *BaseNparserListener) ExitCases_exp(ctx *Cases_expContext) {}

// EnterCase_match_exp is called when production case_match_exp is entered.
func (s *BaseNparserListener) EnterCase_match_exp(ctx *Case_match_expContext) {}

// ExitCase_match_exp is called when production case_match_exp is exited.
func (s *BaseNparserListener) ExitCase_match_exp(ctx *Case_match_expContext) {}

// EnterIreturn is called when production ireturn is entered.
func (s *BaseNparserListener) EnterIreturn(ctx *IreturnContext) {}

// ExitIreturn is called when production ireturn is exited.
func (s *BaseNparserListener) ExitIreturn(ctx *IreturnContext) {}

// EnterIbreak is called when production ibreak is entered.
func (s *BaseNparserListener) EnterIbreak(ctx *IbreakContext) {}

// ExitIbreak is called when production ibreak is exited.
func (s *BaseNparserListener) ExitIbreak(ctx *IbreakContext) {}

// EnterIcontinue is called when production icontinue is entered.
func (s *BaseNparserListener) EnterIcontinue(ctx *IcontinueContext) {}

// ExitIcontinue is called when production icontinue is exited.
func (s *BaseNparserListener) ExitIcontinue(ctx *IcontinueContext) {}

// EnterControl_loop is called when production control_loop is entered.
func (s *BaseNparserListener) EnterControl_loop(ctx *Control_loopContext) {}

// ExitControl_loop is called when production control_loop is exited.
func (s *BaseNparserListener) ExitControl_loop(ctx *Control_loopContext) {}

// EnterControl_loop_exp is called when production control_loop_exp is entered.
func (s *BaseNparserListener) EnterControl_loop_exp(ctx *Control_loop_expContext) {}

// ExitControl_loop_exp is called when production control_loop_exp is exited.
func (s *BaseNparserListener) ExitControl_loop_exp(ctx *Control_loop_expContext) {}

// EnterPrintNormal is called when production printNormal is entered.
func (s *BaseNparserListener) EnterPrintNormal(ctx *PrintNormalContext) {}

// ExitPrintNormal is called when production printNormal is exited.
func (s *BaseNparserListener) ExitPrintNormal(ctx *PrintNormalContext) {}

// EnterPrintFormato is called when production printFormato is entered.
func (s *BaseNparserListener) EnterPrintFormato(ctx *PrintFormatoContext) {}

// ExitPrintFormato is called when production printFormato is exited.
func (s *BaseNparserListener) ExitPrintFormato(ctx *PrintFormatoContext) {}

// EnterElementosPrint is called when production elementosPrint is entered.
func (s *BaseNparserListener) EnterElementosPrint(ctx *ElementosPrintContext) {}

// ExitElementosPrint is called when production elementosPrint is exited.
func (s *BaseNparserListener) ExitElementosPrint(ctx *ElementosPrintContext) {}

// EnterControl_while is called when production control_while is entered.
func (s *BaseNparserListener) EnterControl_while(ctx *Control_whileContext) {}

// ExitControl_while is called when production control_while is exited.
func (s *BaseNparserListener) ExitControl_while(ctx *Control_whileContext) {}

// EnterParametros_funcion is called when production parametros_funcion is entered.
func (s *BaseNparserListener) EnterParametros_funcion(ctx *Parametros_funcionContext) {}

// ExitParametros_funcion is called when production parametros_funcion is exited.
func (s *BaseNparserListener) ExitParametros_funcion(ctx *Parametros_funcionContext) {}

// EnterParametro is called when production parametro is entered.
func (s *BaseNparserListener) EnterParametro(ctx *ParametroContext) {}

// ExitParametro is called when production parametro is exited.
func (s *BaseNparserListener) ExitParametro(ctx *ParametroContext) {}

// EnterLlamada_funcion is called when production llamada_funcion is entered.
func (s *BaseNparserListener) EnterLlamada_funcion(ctx *Llamada_funcionContext) {}

// ExitLlamada_funcion is called when production llamada_funcion is exited.
func (s *BaseNparserListener) ExitLlamada_funcion(ctx *Llamada_funcionContext) {}

// EnterParametros_llamada is called when production parametros_llamada is entered.
func (s *BaseNparserListener) EnterParametros_llamada(ctx *Parametros_llamadaContext) {}

// ExitParametros_llamada is called when production parametros_llamada is exited.
func (s *BaseNparserListener) ExitParametros_llamada(ctx *Parametros_llamadaContext) {}

// EnterParametro_llamada_referencia is called when production parametro_llamada_referencia is entered.
func (s *BaseNparserListener) EnterParametro_llamada_referencia(ctx *Parametro_llamada_referenciaContext) {
}

// ExitParametro_llamada_referencia is called when production parametro_llamada_referencia is exited.
func (s *BaseNparserListener) ExitParametro_llamada_referencia(ctx *Parametro_llamada_referenciaContext) {
}

// EnterElementos_vector is called when production elementos_vector is entered.
func (s *BaseNparserListener) EnterElementos_vector(ctx *Elementos_vectorContext) {}

// ExitElementos_vector is called when production elementos_vector is exited.
func (s *BaseNparserListener) ExitElementos_vector(ctx *Elementos_vectorContext) {}

// EnterMetodos_iniciar_vector is called when production metodos_iniciar_vector is entered.
func (s *BaseNparserListener) EnterMetodos_iniciar_vector(ctx *Metodos_iniciar_vectorContext) {}

// ExitMetodos_iniciar_vector is called when production metodos_iniciar_vector is exited.
func (s *BaseNparserListener) ExitMetodos_iniciar_vector(ctx *Metodos_iniciar_vectorContext) {}

// EnterMetodos_vector is called when production metodos_vector is entered.
func (s *BaseNparserListener) EnterMetodos_vector(ctx *Metodos_vectorContext) {}

// ExitMetodos_vector is called when production metodos_vector is exited.
func (s *BaseNparserListener) ExitMetodos_vector(ctx *Metodos_vectorContext) {}

// EnterPotencia is called when production potencia is entered.
func (s *BaseNparserListener) EnterPotencia(ctx *PotenciaContext) {}

// ExitPotencia is called when production potencia is exited.
func (s *BaseNparserListener) ExitPotencia(ctx *PotenciaContext) {}

// EnterArray is called when production array is entered.
func (s *BaseNparserListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseNparserListener) ExitArray(ctx *ArrayContext) {}

// EnterDimension_array is called when production dimension_array is entered.
func (s *BaseNparserListener) EnterDimension_array(ctx *Dimension_arrayContext) {}

// ExitDimension_array is called when production dimension_array is exited.
func (s *BaseNparserListener) ExitDimension_array(ctx *Dimension_arrayContext) {}

// EnterDimension_acceso_array is called when production dimension_acceso_array is entered.
func (s *BaseNparserListener) EnterDimension_acceso_array(ctx *Dimension_acceso_arrayContext) {}

// ExitDimension_acceso_array is called when production dimension_acceso_array is exited.
func (s *BaseNparserListener) ExitDimension_acceso_array(ctx *Dimension_acceso_arrayContext) {}
