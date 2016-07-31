package main

import (
"fmt"                                                           // import fmt
"os/exec"                                                       // import os/exec
"os"                                                            // import os
)

func Records(dom string) string {

chkWho := exec.Command("whois", dom)                                                            // initial whois command
chkDig := exec.Command("grep", "Name Server\\|Date\\|Status\\|Registrar:\\|Admin E")            // grep arguments

chkDig.Stdin, _ = chkWho.StdoutPipe()                                                           // pipe
chkDig.Stdout=os.Stdout                                                                         // output pipe
_ = chkDig.Start()                                                                              // start dig
_ = chkWho.Run()                                                                                // run initial check
_ = chkDig.Wait()                                                                               // dig wait



return "\n"                                                                                     // return new line when done
}


func main(){

fmt.Println("\n\n")                                                                             // line spacing
fmt.Println(Records(os.Args[1]))                                        // call function with first argument included

}
