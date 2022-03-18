package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionTotal struct {
	Id      string
	Mutable bool
	Publico bool
	Tipo    Ast.TipoRetornado
	Valor   interface{}
	Fila    int
	Columna int
}

func NewDeclaracionTotal(id string, valor interface{}, tipo Ast.TipoRetornado, mutable, publico bool,
	fila int, columna int) DeclaracionTotal {
	nd := DeclaracionTotal{
		Id:      id,
		Mutable: mutable,
		Publico: publico,
		Valor:   valor,
		Tipo:    tipo,
		Fila:    fila,
		Columna: columna,
	}
	return nd
}

func (d DeclaracionTotal) Run(scope *Ast.Scope) interface{} {
	//Verificar si es un tipo especial
	var esEspecial bool = false
	//Verificar que el id no exista

	existe := scope.Exist_actual(d.Id)

	if existe {
		//Ya existe y generar error semántico
		msg := "Semantic error, the element \"" + d.Id + "\" already exist in this scope." +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar que no es un if expresion
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()

	//Verificar que sea un primitivo i64 y la declaración sea usize

	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = d.Valor.(Ast.Instruccion).Run(scope)
	} else {
		preValor = d.Valor.(Ast.Expresion).GetValue(scope)
	}
	valor := preValor.(Ast.TipoRetornado)

	//Cambiar valor de i64 a usize si la declaración es usize y el valor que viene es un i64
	if d.Tipo.Tipo == Ast.USIZE && tipoIn == Ast.I64 {
		valor.Tipo = Ast.USIZE
	}

	//Revisar si el retorno es un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//comparar los tipos
	if !EsTipoEspecial(valor.Tipo) {
		//No es struct,vector,array, entonces comparar los tipos normalmente
		if d.Tipo.Tipo != valor.Tipo {
			//Error de tipos
			msg := "Semantic error, can't initialize a" + expresiones.Tipo_String(d.Tipo) +
				"with " + Ast.ValorTipoDato[valor.Tipo] + " value." +
				" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
			nError := errores.NewError(d.Fila, d.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}

	} else {
		//Primero verificar si el tipo de la declaración no es un acceso a modulo

		if d.Tipo.Tipo == Ast.ACCESO_MODULO {
			//Ejecutar el acceso y cambiar el tipo de la declaración
			nTipo := GetTipoEstructura(d.Tipo, scope)
			errors := ErrorEnTipo(nTipo)
			if errors.Tipo == Ast.ERROR {
				msg := "Semantic error, type error." +
					" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
				nError := errores.NewError(d.Fila, d.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			//De lo contrario actualizar el tipo de la declaracion
			d.Tipo = nTipo
		}

		//Si es un tipo especial, entonces comparar los tipos en profundidad
		tipoEspecial := GetTipoEspecial(valor.Tipo, valor.Valor, scope)
		if tipoEspecial.Tipo == Ast.ERROR {
			//Erro de tipos
			msg := "Semantic error, type error." +
				" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
			nError := errores.NewError(d.Fila, d.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

		if !expresiones.CompararTipos(d.Tipo, tipoEspecial) {
			//Error, los tipos no son correctos
			msg := "Semantic error, type error." +
				" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
			nError := errores.NewError(d.Fila, d.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
		esEspecial = true
	}

	//Todo bien crear y agregar el símbolo

	nSimbolo := Ast.Simbolo{
		Identificador: d.Id,
		Valor:         valor,
		Fila:          d.Fila,
		Columna:       d.Columna,
		Tipo:          valor.Tipo,
		Mutable:       d.Mutable,
		Publico:       d.Publico,
		Entorno:       scope,
	}
	//Agregar el tipo especial
	if esEspecial {
		nSimbolo.TipoEspecial = d.Tipo
	}
	scope.Add(nSimbolo)

	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.EJECUTADO,
	}

}

func (op DeclaracionTotal) GetFila() int {
	return op.Fila
}
func (op DeclaracionTotal) GetColumna() int {
	return op.Columna
}

func (d DeclaracionTotal) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func EsTipoEspecial(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.VECTOR, Ast.ARRAY, Ast.STRUCT:
		return true
	default:
		return false
	}
}

func GetTipoEspecial(tipo Ast.TipoDato, elemento interface{}, scope *Ast.Scope) Ast.TipoRetornado {
	switch tipo {
	case Ast.VECTOR:
		//Get el tipo en estructura
		tipoN := GetTipoEstructura(elemento.(expresiones.Vector).TipoVector, scope)
		errores := ErrorEnTipo(tipoN)
		if errores.Tipo == Ast.ERROR {
			return errores
		} else {
			return tipoN
		}
	case Ast.ARRAY:
		tipoN := GetTipoEstructura(elemento.(expresiones.Array).TipoDelArray, scope)
		errores := ErrorEnTipo(tipoN)
		if errores.Tipo == Ast.ERROR {
			return errores
		} else {
			return tipoN
		}
	case Ast.STRUCT:
		plantilla := elemento.(Ast.Structs).GetPlantilla(scope)
		return Ast.TipoRetornado{
			Tipo:  Ast.STRUCT,
			Valor: plantilla,
		}
	default:
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: true,
		}

	}
}

func GetTipoEstructura(tipo Ast.TipoRetornado, scope *Ast.Scope) Ast.TipoRetornado {

	if tipo.Tipo == Ast.VECTOR {
		return Ast.TipoRetornado{
			Tipo:  Ast.VECTOR,
			Valor: GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope),
		}
	}
	if tipo.Tipo == Ast.ARRAY {
		return Ast.TipoRetornado{
			Tipo:  Ast.ARRAY,
			Valor: GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope),
		}
	}
	if tipo.Tipo == Ast.ACCESO_MODULO {
		//Verificar que se pueda acceder o que exista
		simboloStruct := tipo.Valor.(Ast.AccesosM).GetTipoFromAccesoModulo(tipo, scope)
		if simboloStruct.Tipo == Ast.ERROR {
			return simboloStruct
		}
		if simboloStruct.Tipo != Ast.STRUCT_TEMPLATE {
			//No es un struct
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: true,
			}
		}
		//De lo contrario devolvio el simbolo
		tipoNuevo := simboloStruct.Valor
		return Ast.TipoRetornado{
			Valor: tipoNuevo,
			Tipo:  Ast.STRUCT,
		}
	}
	if tipo.Tipo == Ast.STRUCT {
		//Verificar que el struct exista
		nombreStruct := tipo.Valor.(string)
		if scope.Exist(nombreStruct) {
			return tipo
		} else {
			//No existe el struct
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: true,
			}
		}

	}
	return Ast.TipoRetornado{
		Tipo:  tipo.Tipo,
		Valor: true,
	}
}

func ErrorEnTipo(tipo Ast.TipoRetornado) Ast.TipoRetornado {
	if tipo.Tipo == Ast.ERROR {
		return tipo
	}
	if expresiones.EsTipoFinal(tipo.Tipo) {
		return Ast.TipoRetornado{
			Tipo:  Ast.BOOLEAN,
			Valor: true,
		}
	}
	return ErrorEnTipo(tipo.Valor.(Ast.TipoRetornado))
}
