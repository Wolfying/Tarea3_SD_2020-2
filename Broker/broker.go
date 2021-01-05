package main

import (
	"Tarea3/Broker/broker"
	"Tarea3/Dns/dns"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

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

func (*ServerBroker) CreateDomain(incomestream broker.BrokerHandler_CreateDomainServer) error {
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
func (*ServerBroker) DeleteDomain(incomestream broker.BrokerHandler_DeleteDomainServer) error {
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
func (*ServerBroker) UpdateDomain(incomestream broker.BrokerHandler_UpdateDomainServer) error {
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

func askDomainDns(domain string) string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("8181", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := dns.NewDnsHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := link.GetDomain(ctx)
	_ = err1
	if err != nil {
		log.Printf("Conexion fallida: %s", err)
	}

	waitc := make(chan string)

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al realizar pedido del dominio")
			}

			if in.Ip == " " {
				log.Printf("El dominio no existe")
			}
			waitc <- in.Ip
		}
	}()

	mensaje := dns.Inmessage{
		Domname: domain,
	}

	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	response.CloseSend()
	receive := <-waitc

	return receive
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
