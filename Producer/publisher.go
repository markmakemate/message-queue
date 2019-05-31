package Producer

import (
	"net"
	"io/ioutil"
	"fmt"
	"log"
	"MQ"
)
func chkError(err error){
	if err != nil{
		log.Fatal(err);
	}
}
func listen(addr string) net.Listener{
	listener, err := net.Listen("tcp", addr)
	chkError(err)
	return listener
}
func Handler(conn net.Conn){

}
func connect(listener net.Listener){
	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal(err)
			continue
		}

	}


}


