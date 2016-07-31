package main

import (
"fmt"                                                           // import fmt
"os/exec"                                                       // import os/exec
"os"
)

func Records(dom string, searchBy string, ns string) string {

args := []string{"+noall", "+answer"}                           // limit results of dig with noall and answer
out := exec.Command("dig", args[0], args[1], dom, searchBy)     // check records based on what is passed to the function
stdout, err := out.Output()

if err != nil {                                                 // bit of error checking
        println(err.Error())
        return err.Error()
}                                                               // end if
return (string(stdout))                                         // output the results of the command
}




func main(){

argOne := [4]string{"A", "MX", "NS", "TXT"}                     // create array with records to dig for
argTwo := [3]string{" ", "@NS2", "@8.8.4.4"}                    // create array with NS to dig at

fmt.Println("\n\n")
fmt.Println("-----------------------------------------------------------------")
fmt.Println("|\t\t\t\tDig\t\t\t\t|")
fmt.Println("-----------------------------------------------------------------")


for servers := 0; servers < 3; servers++ {                                              // increment which nameserver is being used
        for recordType := 0; recordType < 4; recordType++ {                             // increment which record is being dug at
                fmt.Println(Records(os.Args[1], argOne[recordType], argTwo[servers]))   // dig at specified nameservers for specified records
        }



        if servers == 0 {                                       // if digging at NS2 add header
                fmt.Println("-----------------------------------------------------------------")
                fmt.Println("|\t\t\t\tAt NS2\t\t\t\t|")
                fmt.Println("-----------------------------------------------------------------\n")
        } else if servers == 1 {
                fmt.Println("-----------------------------------------------------------------")
                fmt.Println("|\t\t\tAt Google Nameservers\t\t\t|")      // if digging at google NS add header
                fmt.Println("-----------------------------------------------------------------\n")
        }

}       // end server for loop

}
