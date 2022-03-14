package expresiones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type DimensionArray struct {
	Tipo         Ast.TipoDato
	TipoArray    Ast.TipoDato
	TipoEspecial string
	Elementos    *arraylist.List
	Fila         int
	Columna      int
}

func NewDimensionArray(elementos *arraylist.List, TipoArray Ast.TipoDato, fila, columna int) DimensionArray {
	nA := DimensionArray{
		Tipo:         Ast.DIMENSION_ARRAY,
		Elementos:    elementos,
		Fila:         fila,
		Columna:      columna,
		TipoArray:    TipoArray,
		TipoEspecial: "",
	}
	return nA
}

func (d DimensionArray) GetValue(scope *Ast.Scope) Ast.TipoRetornado {

	//Reordenar la lista y verificar que son usize
	listaDimensiones := arraylist.New()
	var validarUsize Ast.TipoRetornado
	var elemento interface{}
	var valor Ast.TipoRetornado
	for i := d.Elementos.Len() - 1; i >= 0; i-- {
		elemento = d.Elementos.GetValue(i)
		valor = elemento.(Ast.Expresion).GetValue(scope)
		_, tipoParticular := elemento.(Ast.Abstracto).GetTipo()
		//Validar que sea usize
		validarUsize = EsUsize(valor, tipoParticular, elemento, scope)
		if validarUsize.Tipo == Ast.ERROR {
			return validarUsize
		}
		listaDimensiones.Add(valor)

	}
	return Ast.TipoRetornado{
		Tipo:  Ast.DIMENSION_ARRAY,
		Valor: listaDimensiones,
	}

}

func (a DimensionArray) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, a.Tipo
}

func (a DimensionArray) GetFila() int {
	return a.Fila
}
func (a DimensionArray) GetColumna() int {
	return a.Columna
}

func EsUsize(valor Ast.TipoRetornado, tipoParticular Ast.TipoDato, elemento interface{}, scope *Ast.Scope) Ast.TipoRetornado {
	if (valor.Tipo != Ast.USIZE && valor.Tipo != Ast.I64) ||
		tipoParticular == Ast.IDENTIFICADOR && valor.Tipo == Ast.I64 {
		//Error, se espera un usize
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected USIZE, found. " + Ast.ValorTipoDato[valor.Tipo] +
			". -- Line: " + strconv.Itoa(fila) +
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
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.BOOLEAN,
	}
}
