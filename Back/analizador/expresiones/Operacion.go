package expresiones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"fmt"
	"math"
	"strconv"
)

var suma_dominante = [6][6]Ast.TipoDato{
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.F64, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.STRING_OWNED},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.STRING},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
}

var resta_dominante = [6][6]Ast.TipoDato{
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.F64, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
}

var mul_div_dominante = [6][6]Ast.TipoDato{
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.F64, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
}

type Operacion struct {
	operando_der Ast.Expresion
	operador     string
	operando_izq Ast.Expresion
	unario       bool
	Fila         int
	Columna      int
}

func NewOperation(op_izq Ast.Expresion, operador string,
	op_der Ast.Expresion, unario bool, fila, columna int) Operacion {
	nuevo := Operacion{
		operando_der: op_der,
		operando_izq: op_izq,
		operador:     operador,
		unario:       unario,
		Fila:         fila,
		Columna:      columna,
	}
	return nuevo
}

func (op Operacion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, Ast.PRIMITIVO
}

func (op Operacion) GetValue(entorno Ast.Scope) Ast.TipoRetornado {
	var tipo_izq Ast.TipoRetornado
	var tipo_der Ast.TipoRetornado
	var result_dominante Ast.TipoDato

	if op.unario {
		tipo_izq = op.operando_izq.GetValue(entorno)
	} else {
		tipo_izq = op.operando_izq.GetValue(entorno)
		tipo_der = op.operando_der.GetValue(entorno)
	}
	if tipo_izq.Tipo > 5 || tipo_der.Tipo > 5 {
		//Error, no se pueden operar porque no es ningún valor operable
		msg := "Semantic error, can't operate " + Ast.ValorTipoDato[tipo_izq.Tipo] +
			" type with " + Ast.ValorTipoDato[tipo_der.Tipo] +
			" type. -- Line: " + strconv.Itoa(op.Fila) +
			" Column: " + strconv.Itoa(op.Columna)
		nError := errores.NewError(op.Fila, op.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	switch op.operador {
	case "+":
		result_dominante = suma_dominante[tipo_izq.Tipo][tipo_der.Tipo]

		if result_dominante == Ast.I64 {
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) + tipo_der.Valor.(int),
			}
		} else if result_dominante == Ast.F64 {

			if tipo_izq.Tipo == Ast.I64 {
				tipo_izq = Ast.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			if tipo_der.Tipo == Ast.I64 {
				tipo_der = Ast.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}

			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) + tipo_der.Valor.(float64),
			}
		} else if result_dominante == Ast.STRING || result_dominante == Ast.STRING_OWNED {
			cadena_izq := fmt.Sprintf("%v", tipo_izq.Valor)
			cadena_der := fmt.Sprintf("%v", tipo_der.Valor)
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: cadena_izq + cadena_der,
			}
		} else if result_dominante == Ast.NULL {
			/*
				return Ast.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't add " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type with " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

	case "-", "!":
		if op.unario {
			//Es una operación unaria
			if tipo_izq.Tipo == Ast.F64 && op.operador == "-" {
				//Es un F64
				return Ast.TipoRetornado{
					Tipo:  Ast.F64,
					Valor: tipo_izq.Valor.(float64) * -1,
				}
			} else if tipo_izq.Tipo == Ast.I64 {
				//Es un int
				if op.operador == "-" {
					return Ast.TipoRetornado{
						Tipo:  Ast.I64,
						Valor: tipo_izq.Valor.(int) * -1,
					}
				} else {
					var valorFinal = tipo_izq.Valor.(int)
					//Verificar la regla del bitwise
					if valorFinal >= 0 {
						valorFinal++
					} else {
						valorFinal++
					}
					return Ast.TipoRetornado{
						Tipo:  Ast.I64,
						Valor: valorFinal * -1,
					}
				}

			} else if tipo_izq.Tipo == Ast.BOOLEAN && op.operador == "!" {
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: !tipo_izq.Valor.(bool),
				}

			} else {
				//Tipo no operable
				msg := "Semantic error, can't operate (!) with a " + Ast.ValorTipoDato[tipo_izq.Tipo] +
					" type. -- Line: " + strconv.Itoa(op.Fila) +
					" Column: " + strconv.Itoa(op.Columna)
				nError := errores.NewError(op.Fila, op.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}

		}

		result_dominante = resta_dominante[tipo_izq.Tipo][tipo_der.Tipo]

		if result_dominante == Ast.I64 {
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) - tipo_der.Valor.(int),
			}
		} else if result_dominante == Ast.F64 {

			if tipo_izq.Tipo == Ast.I64 {
				tipo_izq = Ast.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			if tipo_der.Tipo == Ast.I64 {
				tipo_der = Ast.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}

			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) - tipo_der.Valor.(float64),
			}
		} else if result_dominante == Ast.NULL {
			/*
				return Ast.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't subtract " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type with " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

	case "*":
		result_dominante = mul_div_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == Ast.I64 {
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) * tipo_der.Valor.(int),
			}

		} else if result_dominante == Ast.F64 {
			if tipo_izq.Tipo == Ast.I64 {
				tipo_izq = Ast.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			if tipo_der.Tipo == Ast.I64 {
				tipo_der = Ast.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) * tipo_der.Valor.(float64),
			}

		} else if result_dominante == Ast.NULL {
			/*
				return Ast.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't multiply " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type with " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}
	case "/":
		result_dominante = mul_div_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == Ast.I64 {
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) / tipo_der.Valor.(int),
			}

		} else if result_dominante == Ast.F64 {
			if tipo_izq.Tipo == Ast.I64 {
				tipo_izq = Ast.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			if tipo_der.Tipo == Ast.I64 {
				tipo_der = Ast.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) / tipo_der.Valor.(float64),
			}

		} else if result_dominante == Ast.NULL {
			/*
				return Ast.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't divide " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type with " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}
	case "%":
		result_dominante = mul_div_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == Ast.I64 {
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) % tipo_der.Valor.(int),
			}

		} else if result_dominante == Ast.F64 {
			if tipo_izq.Tipo == Ast.I64 {
				tipo_izq = Ast.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			if tipo_der.Tipo == Ast.I64 {
				tipo_der = Ast.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: math.Mod(tipo_izq.Valor.(float64), tipo_der.Valor.(float64)),
			}

		} else if result_dominante == Ast.NULL {
			/*
				return Ast.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't divide " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type with " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}
	case "&&", "||":
		if tipo_der.Tipo != Ast.BOOLEAN || tipo_izq.Tipo != Ast.BOOLEAN {
			msg := "Semantic error, can't logically operate " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type with " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}
		if op.operador == "&&" {
			return Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: tipo_izq.Valor.(bool) && tipo_der.Valor.(bool),
			}
		} else {
			return Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: tipo_izq.Valor.(bool) || tipo_der.Valor.(bool),
			}
		}

	case ">", "<", ">=", "<=", "==", "!=":
		var val_der interface{}
		var val_izq interface{}
		result_dominante = suma_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == Ast.I64 || result_dominante == Ast.F64 {

			if tipo_izq.Tipo == Ast.F64 || tipo_der.Tipo == Ast.F64 {

				if tipo_izq.Tipo == Ast.F64 {
					val_izq = tipo_izq.Valor.(float64)
				} else {
					val_izq = (float64)(tipo_izq.Valor.(int))
				}

				if tipo_izq.Tipo == Ast.F64 {
					val_der = tipo_der.Valor.(float64)
				} else {
					val_der = (float64)(tipo_der.Valor.(int))
				}
				switch op.operador {
				case ">":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) > val_der.(float64),
					}
				case "<":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) < val_der.(float64),
					}
				case ">=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) >= val_der.(float64),
					}
				case "<=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) <= val_der.(float64),
					}
				case "==":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) == val_der.(float64),
					}
				case "!=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) != val_der.(float64),
					}
				}
			} else {
				val_izq = tipo_izq.Valor.(int)
				val_der = tipo_der.Valor.(int)
				switch op.operador {
				case ">":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) > val_der.(int),
					}
				case "<":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) < val_der.(int),
					}
				case ">=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) >= val_der.(int),
					}
				case "<=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) <= val_der.(int),
					}
				case "==":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) == val_der.(int),
					}
				case "!=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) != val_der.(int),
					}
				}
			}
		} else if result_dominante == Ast.STR {
			//Es una comparación entre STR
			val_izq = tipo_izq.Valor.(string)
			val_der = tipo_der.Valor.(string)
			switch op.operador {
			case ">":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) > len(val_der.(string)),
				}
			case "<":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) < len(val_der.(string)),
				}
			case ">=":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) >= len(val_der.(string)),
				}
			case "<=":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) <= len(val_der.(string)),
				}
			case "==":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) == len(val_der.(string)),
				}
			case "!=":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) != len(val_der.(string)),
				}
			}
			return Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: len(val_izq.(string)) > len(val_der.(string)),
			}
		}
		msg := "Semantic error, can't compare " + Ast.ValorTipoDato[tipo_izq.Tipo] +
			" with " + Ast.ValorTipoDato[tipo_der.Tipo] +
			" type. -- Line: " + strconv.Itoa(op.Fila) +
			" Column: " + strconv.Itoa(op.Columna)
		nError := errores.NewError(op.Fila, op.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	return Ast.TipoRetornado{Tipo: Ast.NULL, Valor: nil}
}
