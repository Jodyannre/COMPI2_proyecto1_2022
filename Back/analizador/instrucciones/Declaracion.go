package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"fmt"
	"strconv"
)

type Declaracion struct {
	Id          string
	Tipo        Ast.TipoDato
	TipoRetorno Ast.TipoDato
	Mutable     bool
	Publico     bool
	Valor       interface{}
	Fila        int
	Columna     int
}

func NewDeclaracion(id string, tipo Ast.TipoDato, mutable, publico bool, tipoRetorno Ast.TipoDato,
	valor interface{}, fila int, columna int) Declaracion {
	nd := Declaracion{
		Id:          id,
		Tipo:        tipo,
		TipoRetorno: tipoRetorno,
		Mutable:     mutable,
		Publico:     publico,
		Valor:       valor,
		Fila:        fila,
		Columna:     columna,
	}
	return nd
}

func (d Declaracion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d Declaracion) Run(scope *Ast.Scope) interface{} {
	//Verificar que el id no exista

	existe := scope.Exist_actual(d.Id)

	//Verificar que sea un modulo, función o vector
	if d.Tipo == Ast.MODULO {
		fmt.Println("Declaración de módulo")
	}
	if d.Tipo == Ast.FUNCION {
		fmt.Println("Declaración de Función")
	}
	if d.Tipo == Ast.VECTOR {
		fmt.Println("Declaración de vector")
	}
	if d.Tipo == Ast.STRUCT {
		fmt.Println("Declaración de struct")
	}

	//Verificar que los tipos sean correctos si

	//Primero verificar que no es un if expresion
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()

	//Verificar que sea un primitivo i64 y la declaración sea usize

	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = d.Valor.(Ast.Instruccion).Run(scope)
	} else if tipoIn == Ast.FUNCION {
		preValor = Ast.TipoRetornado{
			Valor: d.Valor,
			Tipo:  Ast.FUNCION,
		}
	} else {
		preValor = d.Valor.(Ast.Expresion).GetValue(scope)
	}
	valor := preValor.(Ast.TipoRetornado)
	//Cambiar valor de i64 a usize si la declaración es usize
	if d.Tipo == Ast.USIZE && tipoIn == Ast.I64 {
		valor.Tipo = Ast.USIZE
	}

	//Revisar si se declara un vec o array y la expresion es diferente
	if d.Tipo == Ast.VECTOR && valor.Tipo != Ast.VECTOR {
		if valor.Tipo == Ast.ERROR {
			return valor
		}
		//Error, no se puede inicializar un vector con un valor
		msg := "Semantic error, can't initialize a Vector with " + Ast.ValorTipoDato[valor.Tipo] + " type" +
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

	//Revisar si el retorno es un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	if valor.Tipo == d.Tipo && !existe {
		//El tipo es correcto y no existe en el entorno actual
		//Crear símbolo y agregarlo a la tabla del entorno actual
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          d.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
			Entorno:       scope,
		}
		//Si es vector o array verificar si es referencia o no
		if valor.Tipo == Ast.VECTOR {
			nSimbolo.Referencia = valor.Valor.(expresiones.Vector).Referencia
			nValor := valor.Valor.(expresiones.Vector)
			if nValor.TipoVector == Ast.INDEFINIDO {
				nValor.TipoVector = d.Tipo
			}
			nSimbolo.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: nValor}
		}
		//Si es función, módulo o struct, agregarlos a las listas globales
		scope.Add(nSimbolo)
		if valor.Tipo == Ast.FUNCION ||
			valor.Tipo == Ast.MODULO ||
			valor.Tipo == Ast.STRUCT {
			scope.Addfms(nSimbolo)
		}

	} else if d.Tipo == Ast.INDEFINIDO && !existe {
		//Es una declaración sin valor asignado
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          valor.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
		}
		if valor.Tipo == Ast.VECTOR {
			nSimbolo.Referencia = valor.Valor.(expresiones.Vector).Referencia
			nValor := valor.Valor.(expresiones.Vector)
			if nValor.TipoVector == Ast.INDEFINIDO {
				nValor.TipoVector = d.Tipo
			}
			nSimbolo.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: nValor}
		}
		scope.Add(nSimbolo)
	} else if d.Tipo != Ast.INDEFINIDO && !existe && valor.Tipo == Ast.NULL {
		//Es una declaración sin valor asignado
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          valor.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
		}
		scope.Add(nSimbolo)
	} else if existe {
		//Ya existe y generar error semántico
		//fmt.Println("Error, ese elemento ya existe en el ámbito local")
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
	} else {
		//Error de tipos, generar error semántico
		//fmt.Println("Error, los tipos no coinciden en la declaración")
		msg := "Semantic error, can't assign " + Ast.ValorTipoDato[int(valor.Tipo)] +
			" type to " + Ast.ValorTipoDato[int(d.Tipo)] +
			" type. -- Line: " + strconv.Itoa(d.Fila) +
			" Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
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

func (op Declaracion) GetFila() int {
	return op.Fila
}
func (op Declaracion) GetColumna() int {
	return op.Columna
}
