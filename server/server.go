package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Server struct {
	Conn     net.Conn
	Listener net.Listener
    DB		MyDB
	QuitChannel  chan os.Signal
}


func NewServer()*Server{
	Listener, err := net.Listen("tcp", "localhost:8000")
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}

	var q= make(chan os.Signal)

	srv:= &Server{Listener:Listener,DB:NewDB(),QuitChannel:q}
	go srv.Serv()
	fmt.Println("->adi db started on port 8000")
	return srv
}


func (srv *Server) Serv(){

	tcpListener := srv.Listener.(*net.TCPListener)

	//err:=tcpListener.SetDeadline(time.Now().Add(10*time.Second))
	//if err!=nil{
	//	fmt.Println("Failed to set deadline ",err.Error())
	//}

	conn,err:= tcpListener.Accept()
	if err!=nil{
		fmt.Println("failed to accept connection ",err)
	}

	srv.Write(conn,"Welcome to adi DB")

	srv.handleConn(conn)
}

func (srv *Server)handleConn(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		l:=strings.ToLower(strings.TrimSpace(scanner.Text()))
		values :=strings.Split(l," ")

		switch  {
		case len(values)==3 && values[0]=="set":
			srv.DB.Set(values[1],values[2])
			srv.Write(conn,"OK")
		case len(values)==2 && values[0]=="get":
			val,ok :=srv.DB.Get(values[1])
			if ok{
				srv.Write(conn,val)
			}else{
				srv.Write(conn,"Value not found.")
			}
		case len(values)==2 && values[0]=="del":
			_,ok := srv.DB.Get(values[1])
			if ok{
				srv.DB.Del(values[1])
				srv.Write(conn,"value deleted")
			}else{
				srv.Write(conn,"Value not found.")
			}
		case len(values)==1 && values[0]=="exit":
			srv.DB.Save()
			srv.QuitChannel<-os.Interrupt
			srv.Write(conn,"Exiting from database.")

		default:
			srv.Write(conn,"Wrong command!")
		}
	}
}

func (srv *Server)Write(conn net.Conn, s string) {
	_, err := fmt.Fprintf(conn, "->%s\n->", s)
	if err!=nil{
		fmt.Println(err)
	}
}