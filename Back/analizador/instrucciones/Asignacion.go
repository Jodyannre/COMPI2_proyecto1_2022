package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
)

type Asignacion struct {
	Id      string
	Valor   interface{}
	Fila    int
	Columna int
}

func NewAsignacion(id string, valor interface{}, fila int, columna int) Asignacion {
	na := Asignacion{
		Id:      id,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
	}
	return na
}

func (a Asignacion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.ASIGNACION
}

func (a Asignacion) Run(scope *Ast.Scope) interface{} {
	//Verificar que el id  exista
	existe := scope.Exist(a.Id)
	//Obtener el valor del id
	simbolo_id := scope.GetSimbolo(a.Id)
	//Verificar que los tipos sean correctos
	//Primero verificar que no es un if expresion
	_, tipoIn := a.Valor.(Ast.Abstracto).GetTipo()
	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION {
		preValor = a.Valor.(Ast.Instruccion).Run(scope)
	} else {
		preValor = a.Valor.(Ast.Expresion).GetValue(scope)
	}
	valor := preValor.(Ast.TipoRetornado)

	if existe {
		//Primero verificar si es mutable
		if !simbolo_id.Mutable {
			//No es mutable, error semántico
			msg := "Semantic error, can't modify a non-mutable " + Ast.ValorTipoDato[int(simbolo_id.Tipo)] +
				" type. -- Line: " + strconv.Itoa(a.Fila) +
				" Column: " + strconv.Itoa(a.Columna)
			nError := errores.NewError(a.Fila, a.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

		//Existe, ahora verificar los tipos
		if simbolo_id.Valor.(Ast.TipoRetornado).Tipo == valor.Tipo {
			//Los tipos son correctos, actualizar el símbolo
			simbolo_id.Valor = valor
			scope.UpdateSimbolo(a.Id, simbolo_id)
		} else {
			//Revisar si el retorno es un error
			if valor.Tipo == Ast.ERROR {
				return valor
			}
			//Error de tipos, generar un error semántico
			//fmt.Println("Erro de tipos")
			msg := "Semantic error, can't assign " + Ast.ValorTipoDato[int(valor.Tipo)] +
				" type to " + Ast.ValorTipoDato[int(simbolo_id.Valor.(Ast.TipoRetornado).Tipo)] +
				" type. -- Line: " + strconv.Itoa(a.Fila) +
				" Column: " + strconv.Itoa(a.Columna)
			nError := errores.NewError(a.Fila, a.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
	} else {
		//No existe, generar un error semántico
		msg := "Semantic error, the element \"" + a.Id + "\" doesn't exist in any scope." +
			" -- Line:" + strconv.Itoa(a.Fila) + " Column: " + strconv.Itoa(a.Columna)
		nError := errores.NewError(a.Fila, a.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (op Asignacion) GetFila() int {
	return op.Fila
}
func (op Asignacion) GetColumna() int {
	return op.Columna
}
