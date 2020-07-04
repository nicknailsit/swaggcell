package core

import "net"

type IServer interface {

}

type Server struct {
	Addr net.Addr
	clients map[string]interface{}
	routes map[string]interface{}
	zones map[string]interface{}

}