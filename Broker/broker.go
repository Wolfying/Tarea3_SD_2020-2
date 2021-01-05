package main

import (
	"Tarea3/Broker/broker"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

// ServerBroker ...
type ServerBroker struct {
	name string
}

// ConnectDomain ...
func (*ServerBroker) ConnectDomain(incomestream broker.BrokerHandler_ConnectDomainServer) error {
	for {
		in, err := incomestream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("imprimo algo antes de return err: %v", err)
			return err
		}
		_ = in
	}

	return nil
}

func main() {
	server := ServerBroker{}
	puerto := ":8180"
	lis, err := net.Listen("tcp", puerto)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", puerto, err)
	}
	grpcServer := grpc.NewServer()

	broker.RegisterBrokerHandlerServer(grpcServer, &server)
	fmt.Printf("Server Iniciado en el puerto %s \n", puerto)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", puerto, err)
	}

}
