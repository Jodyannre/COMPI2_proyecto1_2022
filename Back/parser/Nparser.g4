parser grammar Nparser;

options{
    tokenVocab = Nlexer;
    language = Go;
}


@header{
    import "github.com/colegno/arraylist"
}

/* 
@members{
}
*/


inicio returns[*arraylist.List lista] 
            : expresion
            {

            }
            ;

/* 
instrucciones returns[*arraylist.List list]  : instruccion;


instruccion returns[*arraylist.List list] : expresion;
*/

expresion returns[*arraylist.List list] : 'hola';
