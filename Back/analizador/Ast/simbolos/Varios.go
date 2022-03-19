package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

func (a AccesoModulo) GetTipoFromAccesoModulo(tipo Ast.TipoRetornado, scope *Ast.Scope) Ast.TipoRetornado {
	acceso := tipo.Valor.(AccesoModulo)
	resultadoAcceso := acceso.GetValue(scope)

	if resultadoAcceso.Tipo == Ast.ERROR {
		return resultadoAcceso
	}

	if resultadoAcceso.Tipo != Ast.STRUCT_TEMPLATE {
		fila := acceso.Fila
		columna := acceso.Columna
		msg := "Semantic error, an STRUCT was expected." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	simboloStruct := resultadoAcceso.Valor.(Ast.Simbolo)
	nombrePlantilla := simboloStruct.Identificador
	return Ast.TipoRetornado{
		Tipo:  Ast.STRUCT_TEMPLATE,
		Valor: nombrePlantilla,
	}

}

func (a AccesoModulo) GetTipoEstructura(tipo Ast.TipoRetornado, scope *Ast.Scope) Ast.TipoRetornado {

	if tipo.Tipo == Ast.VECTOR {
		return Ast.TipoRetornado{
			Tipo:  Ast.VECTOR,
			Valor: a.GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope),
		}
	}
	if tipo.Tipo == Ast.ARRAY {
		return Ast.TipoRetornado{
			Tipo:  Ast.ARRAY,
			Valor: a.GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope),
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
		//tipoNuevo := simboloStruct.Valor.(Ast.Simbolo).Identificador
		return Ast.TipoRetornado{
			Valor: simboloStruct.Valor,
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

func GetTipoEstructura(tipo Ast.TipoRetornado, scope *Ast.Scope, elemento interface{}) Ast.TipoRetornado {

	if tipo.Tipo == Ast.VECTOR {
		return Ast.TipoRetornado{
			Tipo:  Ast.VECTOR,
			Valor: GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope, elemento),
		}
	}
	if tipo.Tipo == Ast.ARRAY {
		return Ast.TipoRetornado{
			Tipo:  Ast.ARRAY,
			Valor: GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope, elemento),
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
			fila := elemento.(Ast.Abstracto).GetFila()
			columna := elemento.(Ast.Abstracto).GetColumna()
			_, tipoParticular := elemento.(Ast.Abstracto).GetTipo()
			msg := ""
			msg = "Semantic error,  STRUCT expected, found " + Ast.ValorTipoDato[tipoParticular] +

				" -- Line: " + strconv.Itoa(fila) +
				" Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
		//De lo contrario devolvio el simbolo
		//tipoNuevo := simboloStruct.Valor.(Ast.Simbolo).Identificador
		return Ast.TipoRetornado{
			Valor: simboloStruct.Valor,
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
			fila := elemento.(Ast.Abstracto).GetFila()
			columna := elemento.(Ast.Abstracto).GetColumna()
			msg := ""
			msg = "Semantic error, element \"" + nombreStruct + "\", doesn't exist" +

				" -- Line: " + strconv.Itoa(fila) +
				" Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
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

func EsPosibleReferencia(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.VECTOR, Ast.ARRAY, Ast.STRUCT, Ast.ACCESO_MODULO:
		return true
	default:
		return false
	}
}
