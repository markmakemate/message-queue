package Consumer

import (
	"net"
	"log"
	"fmt"
	"io/ioutil"
)

func chkError(err error){
	if err != nil{
		log.Fatal(err);
	}
}
func connect(addr string){
	
}