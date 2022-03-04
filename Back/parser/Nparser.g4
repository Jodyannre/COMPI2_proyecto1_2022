parser grammar Nparser;

options{
    tokenVocab = Nlexer;
    language = Go;
}


@header{
    import "github.com/colegno/arraylist"
    import "Back/analizador/Ast"
    import "Back/analizador/expresiones"
    import "Back/analizador/instrucciones"
    import "Back/analizador/exp_ins"
}

/* 
@members{
}
*/


inicio returns[*arraylist.List lista] 
            : instrucciones	
            {
                $lista = $instrucciones.list
            }
;


bloque returns[*arraylist.List list] 
            : LLAVE_IZQ instrucciones LLAVE_DER	
            {
                $list = $instrucciones.list
            }
;

instrucciones returns [*arraylist.List list]
			@init{
				$list =  arraylist.New()
			}
			: e += instruccion*  {
				listInt := localctx.(*InstruccionesContext).GetE()
					for _, e := range listInt {
						$list.Add(e.GetEx())
					}
			}
;

instruccion returns[interface{} ex] 
			:expresion 		    	{$ex = $expresion.ex}	
            |declaracion PUNTOCOMA  {$ex = $declaracion.ex}	
            |asignacion PUNTOCOMA   {$ex = $asignacion.ex}
            |control_if             {$ex = $control_if.ex}	    		     
;

/* 
instrucciones returns[*arraylist.List list]  : instruccion;


instruccion returns[*arraylist.List list] : expresion;
*/


declaracion returns[Ast.Instruccion ex]
    : LET ID IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,Ast.INDEFINIDO,
            false,false,Ast.VOID,$expresion.ex,fila,columna)
        }
    | LET ID IGUAL control_expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,Ast.INDEFINIDO,
            false,false,Ast.VOID,$control_expresion.ex,fila,columna)
        }
    | LET ID DOSPUNTOS tipo_dato IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            false,false,Ast.VOID,$expresion.ex,fila,columna)            
        }
    | LET ID DOSPUNTOS tipo_dato IGUAL control_expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            false,false,Ast.VOID,$control_expresion.ex,fila,columna)            
        }
    | LET MUT ID IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,Ast.INDEFINIDO,
                true,false,Ast.VOID,$expresion.ex,fila,columna)                 
        }
    | LET MUT ID IGUAL control_expresion
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = instrucciones.NewDeclaracion($ID.text,Ast.INDEFINIDO,
            true,false,Ast.VOID,$control_expresion.ex,fila,columna)  
    }
    | LET MUT ID DOSPUNTOS tipo_dato
        {
            valor := expresiones.NewPrimitivo(nil, Ast.NULL)
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            true,false,Ast.VOID,valor,fila,columna)                
        }
    | LET MUT ID DOSPUNTOS tipo_dato IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            true,false,Ast.VOID,$expresion.ex,fila,columna)               
        }
    | LET MUT ID DOSPUNTOS tipo_dato IGUAL control_expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            true,false,Ast.VOID,$control_expresion.ex,fila,columna)               
        }
    //| STRUCT ID_CAMEL LLAVE_IZQ atributos LLAVE_DER
;


asignacion returns[Ast.Instruccion ex]
    : ID IGUAL expresion
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = instrucciones.NewAsignacion($ID.text,$expresion.ex,fila,columna)
    }
    | ID IGUAL control_if_exp
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = instrucciones.NewAsignacion($ID.text,$control_if_exp.ex,fila,columna)
    }
;


/* 
atributos returns[*arraylist.List list]
@init{ list := arrayList.New()}
    : atributos atributo
    | atributo
;

atributo returns[simbolos.Atributo ex]
    : ID DOSPUNTOS tipo_dato
;
*/

expresion returns[Ast.Expresion ex] 
    :   op=(RESTA|NOT) op_izq= expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,nil,true,fila,columna)
        }
    |   op_izq=expresion op=(MULTIPLICACION|DIVISION|MODULO) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }    
    
    |   op_izq=expresion op=(SUMA|RESTA) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion op=(MAYOR_I|MAYOR|MENOR_I|MENOR|IGUALDAD|DISTINTO) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion op=(IGUALDAD|DISTINTO) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion AND op_der=expresion
        {
            fila := $AND.line
            columna := $AND.pos
            $ex = expresiones.NewOperation($op_izq.ex,$AND.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion OR op_der=expresion
        {
            fila := $OR.line
            columna := $OR.pos
            $ex = expresiones.NewOperation($op_izq.ex,$OR.text,$op_der.ex,false,fila,columna)
        }
    |   PAR_IZQ expresion PAR_DER
        {
            $ex = $expresion.ex
        }
    |   ID		
        {
            id := $ID.text
            $ex = expresiones.NewIdentificador (id,Ast.IDENTIFICADOR)
        }
    |   TRUE        
        {
            valor := true
            $ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN)
        }         
    |   FALSE
        {
            valor := false
            $ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN)
        }     
    |   CARACTER
        {
            valor := $CARACTER.text
            valor = valor[1:len(valor)-1]
            $ex = expresiones.NewPrimitivo(valor, Ast.CHAR)
        }   
    |   DECIMAL
        {
            valor,err := strconv.ParseFloat($DECIMAL.text,64)
            if err!= nil{
                fmt.Println("Hay un error en el get número")
                fmt.Println(err)
            }
            $ex = expresiones.NewPrimitivo(valor, Ast.F64)
        }   
    |   NUMERO
        {
            valor,err := strconv.Atoi($NUMERO.text)
            if err!= nil{
                fmt.Println("Hay un error en el get número")
                fmt.Println(err)
            }
            $ex = expresiones.NewPrimitivo(valor, Ast.I64)
        }   
    |   CADENA
        {
            valor := $CADENA.text
            valor = valor[1:len(valor)-1]
            $ex = expresiones.NewPrimitivo(valor, Ast.STR)
        }             
;

tipo_dato returns[Ast.TipoDato ex]
    :   BOOL    {$ex = Ast.BOOLEAN}
    |   CHAR    {$ex = Ast.CHAR}
    |   I64     {$ex = Ast.I64}
    |   F64     {$ex = Ast.F64}
    |   STR     {$ex = Ast.STR}
    |   STRING  {$ex = Ast.STRING}
    |   USIZE   {$ex = Ast.USIZE}
;





control_if returns[Ast.Instruccion ex]
	:IF PAR_IZQ expresion PAR_DER bloqueIf = bloque
	{
		fila:= $IF.line
		columna:= $PAR_IZQ.pos
		columna++
		lista_null := arraylist.New()
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_null,Ast.IF,fila,columna,false)
	}	
	|IF PAR_IZQ expresion PAR_DER bloqueIf = bloque ELSE bloqueElse = bloque
	  
	{
		fila:= $IF.line
		columna:= $PAR_IZQ.pos
		columna++
        lista_entonces := arraylist.New()
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion.ex,$bloqueElse.list,lista_null,Ast.ELSE,fila,columna,false)
		lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF,fila,columna,false)	
	}
	|IF PAR_IZQ expresion PAR_DER bloqueIf = bloque bloque_else_if
	{
		fila:= $IF.line
		columna:= $PAR_IZQ.pos
		columna++
		lista_entonces := $bloque_else_if.list
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF,fila,columna,false)		
	}
	|IF PAR_IZQ expresion PAR_DER bloqueIf = bloque bloque_else_if ELSE bloqueElse = bloque
	{
		fila:= $IF.line
		columna:= $PAR_IZQ.pos
		columna++
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion.ex,$bloqueElse.list,lista_null,Ast.ELSE,fila,columna,false)
		lista_entonces := $bloque_else_if.list
        lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF,fila,columna,false)		
	}
;


bloque_else_if returns [*arraylist.List list]
@init{ $list = arraylist.New()}
: lista += else_if+ {
                    listas := localctx.(*Bloque_else_ifContext).GetLista()
                    for _, e := range listas {
                        $list.Add(e.GetEx())
                    }
    }
;


else_if returns [Ast.Instruccion ex]
    : ELSE IF PAR_IZQ expresion PAR_DER bloque
    {
        fila:= $ELSE.line
		columna:= $PAR_IZQ.pos
		columna++
        lista_null := arraylist.New()
        $ex = exp_ins.NewIF($expresion.ex,$bloque.list,lista_null,Ast.ELSEIF,fila,columna,false)	
    }
;





control_if_exp returns[Ast.Instruccion ex]
	:IF PAR_IZQ expresion PAR_DER bloqueIf = bloque
	{
		fila:= $IF.line
		columna:= $PAR_IZQ.pos
		columna++
		lista_null := arraylist.New()
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_null,Ast.IF_EXPRESION,fila,columna,true)
	}	
	|IF PAR_IZQ expresion PAR_DER bloqueIf = bloque ELSE bloqueElse = bloque
	  
	{
		fila:= $IF.line
		columna:= $PAR_IZQ.pos
		columna++
        lista_entonces := arraylist.New()
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion.ex,$bloqueElse.list,lista_null,Ast.ELSE_EXPRESION,fila,columna,true)
		lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF_EXPRESION,fila,columna,true)	
	}
	|IF PAR_IZQ expresion PAR_DER bloqueIf = bloque bloque_else_if_exp
	{
		fila:= $IF.line
		columna:= $PAR_IZQ.pos
		columna++
		lista_entonces := $bloque_else_if_exp.list
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF_EXPRESION,fila,columna,true)		
	}
	|IF PAR_IZQ expresion PAR_DER bloqueIf = bloque bloque_else_if_exp ELSE bloqueElse = bloque
	{
		fila:= $IF.line
		columna:= $PAR_IZQ.pos
		columna++
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion.ex,$bloqueElse.list,lista_null,Ast.ELSE_EXPRESION,fila,columna,true)
		lista_entonces := $bloque_else_if_exp.list
        lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF_EXPRESION,fila,columna,true)		
	}
;


bloque_else_if_exp returns [*arraylist.List list]
@init{ $list = arraylist.New()}
: lista += else_if_exp+ {
                    listas := localctx.(*Bloque_else_if_expContext).GetLista()
                    for _, e := range listas {
                        $list.Add(e.GetEx())
                    }
    }
;


else_if_exp returns [Ast.Instruccion ex]
    : ELSE IF PAR_IZQ expresion PAR_DER bloque
    {
        fila:= $ELSE.line
		columna:= $PAR_IZQ.pos
		columna++
        lista_null := arraylist.New()
        $ex = exp_ins.NewIF($expresion.ex,$bloque.list,lista_null,Ast.ELSEIF_EXPRESION,fila,columna,true)	
    }
;


control_expresion returns [Ast.Instruccion ex]
    : control_if_exp {$ex = $control_if_exp.ex}
;