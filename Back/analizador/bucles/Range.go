package bucles

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type Range struct {
	Tipo     Ast.TipoDato
	ValorInf interface{}
	ValorSup interface{}
	Fila     int
	Columna  int
}

func NewRange(tipo Ast.TipoDato, valInf, valSup interface{}, fila, columna int) Range {
	nR := Range{
		Tipo:     tipo,
		ValorSup: valSup,
		ValorInf: valInf,
		Fila:     fila,
		Columna:  columna,
	}
	return nR
}

func (r Range) GetValue(scope *Ast.Scope) Ast.TipoRetornado {

	//Verificar el tipo de range
	ValorSup := r.ValorSup.(Ast.Expresion).GetValue(scope)
	ValorInf := r.ValorInf.(Ast.Expresion).GetValue(scope)
	var numero int
	listaElemento := arraylist.New()
	var limInf, limSup int
	var vectorSalida expresiones.Vector
	var nElemento Ast.TipoRetornado
	tipoVector := Ast.TipoRetornado{
		Tipo:  Ast.I64,
		Valor: true,
	}

	//Verificar errores
	if ValorSup.Tipo == Ast.ERROR {
		return ValorSup
	}

	if ValorInf.Tipo == Ast.ERROR {
		return ValorInf
	}

	if r.Tipo == Ast.RANGE_EXPRESION {
		//Quiere decir que se itera un vector
		//Solo viene la parte inf

		//Verificar que sea un vector o un array o error

		if ValorInf.Tipo != Ast.VECTOR &&
			ValorInf.Tipo != Ast.ARRAY {
			//ERROR
			fila := r.ValorInf.(Ast.Abstracto).GetFila()
			columna := r.ValorInf.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, (VECTOR|ARRAY) was expected, found " + Ast.ValorTipoDato[ValorInf.Tipo] +
				". -- Line:" + strconv.Itoa(fila) + " Column: " +
				strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			scope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		} else {
			return ValorInf
		}

	} else {
		//QUiere decir que se iteran con un iterador numerico
		//crear la lista y luego el vector
		//Conseguir los límites
		//Verificar que ambos son números

		if ValorInf.Tipo != Ast.I64 && ValorInf.Tipo != Ast.USIZE {
			fila := r.ValorInf.(Ast.Abstracto).GetFila()
			columna := r.ValorInf.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, (i64|USIZE) was expected, found " + Ast.ValorTipoDato[ValorInf.Tipo] +
				". -- Line:" + strconv.Itoa(fila) + " Column: " +
				strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			scope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

		if ValorSup.Tipo != Ast.I64 && ValorSup.Tipo != Ast.USIZE {
			fila := r.ValorInf.(Ast.Abstracto).GetFila()
			columna := r.ValorInf.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, (i64|USIZE) was expected, found " + Ast.ValorTipoDato[ValorSup.Tipo] +
				". -- Line:" + strconv.Itoa(fila) + " Column: " +
				strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			scope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

		//Todo bien, recuperar los int
		limInf = ValorInf.Valor.(int)
		limSup = ValorSup.Valor.(int)

		//Recorrer la lista e ir creando los elementos que se van a recorrer

		for i := limInf; i <= limSup; i++ {
			numero = int(i)
			nElemento = Ast.TipoRetornado{
				Tipo:  Ast.I64,
				Valor: numero,
			}
			listaElemento.Add(nElemento)
		}

		vectorSalida = expresiones.NewVector(listaElemento, tipoVector, listaElemento.Len(),
			listaElemento.Len(), false, r.Fila, r.Columna)
		return Ast.TipoRetornado{
			Tipo:  Ast.VECTOR,
			Valor: vectorSalida,
		}

	}

}

func (op Range) GetFila() int {
	return op.Fila
}
func (op Range) GetColumna() int {
	return op.Columna
}
func (f Range) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, f.Tipo
}
