package colorize

import (
    "fmt"
    "log"
    "testing"
    //"os"
)


var (
    DEBUG *log.Logger   
    WARN *log.Logger
)

func TestStatefulColor(t *testing.T){
    var red, blue Colorize
    
    //set some state
    red.SetColor(COLOR_RED)
    blue.SetColor(COLOR_BLUE)
    
    out_string := fmt.Sprintf("This %s a %s", red.Fmt("is"), blue.Fmt("test"))
    basis_string := "This \033[0;31mis\033[0m a \033[0;34mtest\033[0m"
    if out_string != basis_string{
        t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
    }else{
        fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
    }



    //Output:This \033[0;31mis a test
    //
}
