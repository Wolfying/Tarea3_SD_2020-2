package main

import (
	"Tarea3/Broker/broker"
	"Tarea3/Dns/dns"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

// ServerBroker ...
type ServerBroker struct {
	name string
}

var ips []string = []string{"dist142:8181", "dist143:8181", "dist144:8181"}

// ConnectDomain ...
func (*ServerBroker) ConnectDomain(incomestream broker.BrokerHandler_ConnectDomainServer) error {
	for {
		in, err := incomestream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// log.Printf("imprimo algo antes de return err: %v", err)
			return err
		}
		_ = in
		log.Printf("Se recibe solicitud para el dominio: %s", in.Domname)
		ip := askDomainDns(in.Domname)
		mensaje := &broker.Response{
			Ip: ip,
		}

		if err := incomestream.Send(mensaje); err != nil {
			log.Printf("Error en propuesta  %s", err)
			return err
		}
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
		var mensaje *broker.Response
		if DnsCreateDomain(in.Domname) {
			mensaje = &broker.Response{
				Ip: "Dominio creado con exito",
			}

		} else {
			mensaje = &broker.Response{
				Ip: "Dominio no ha podido ser creado",
			}
		}
		if err := incomestream.Send(mensaje); err != nil {
			log.Printf("Error en solicitud  %s", err)
			return err
		}
	}
	return nil
}
func DnsCreateDomain(dom string) bool {
	var conn *grpc.ClientConn
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	ip := ips[rand.Intn(len(ips))]

	conn, err := grpc.Dial(ip, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := dns.NewDnsHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := link.CreateDomain(ctx)
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
				log.Fatalf("Error al crear dominio")
			}

			if in.Ip == " " {
				log.Printf("El dominio no existe")
			}
			waitc <- in.Ip
		}
	}()

	mensaje := dns.Domcreate{
		Domname: dom,
	}

	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	response.CloseSend()
	<-waitc

	return true
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

		var mensaje *broker.Response
		if DnsDeleteDomain(in.Domname) {
			mensaje = &broker.Response{
				Ip: "Dominio eliminado con exito",
			}

		} else {
			mensaje = &broker.Response{
				Ip: "Dominio no ha podido ser eliminado",
			}
		}
		if err := incomestream.Send(mensaje); err != nil {
			log.Printf("Error en solicitud  %s", err)
			return err
		}
	}
	return nil
}

func DnsDeleteDomain(dom string) bool {
	var conn *grpc.ClientConn
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	ip := ips[rand.Intn(len(ips))]

	conn, err := grpc.Dial(ip, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := dns.NewDnsHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := link.DeleteDomain(ctx)
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
				log.Fatalf("Error al crear dominio")
			}

			if in.Ip == " " {
				log.Printf("El dominio no existe")
			}
			waitc <- in.Ip
		}
	}()

	mensaje := dns.Dom{
		Domname: dom,
	}

	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	response.CloseSend()
	<-waitc

	return true
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
		var mensaje *broker.Response
		if DnsUpdateDomain(in.OldDom, in.NewDom) {
			mensaje = &broker.Response{
				Ip: "Dominio actualizado con exito",
			}

		} else {
			mensaje = &broker.Response{
				Ip: "Dominio no ha podido ser actualizado",
			}
		}
		if err := incomestream.Send(mensaje); err != nil {
			log.Printf("Error en solicitud  %s", err)
			return err
		}
	}
	return nil
}

func DnsUpdateDomain(oldDom string, newDom string) bool {
	var conn *grpc.ClientConn

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	ip := ips[rand.Intn(len(ips))]

	conn, err := grpc.Dial(ip, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := dns.NewDnsHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := link.UpdateDomain(ctx)
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
				log.Fatalf("Error al crear dominio")
			}

			if in.Ip == " " {
				log.Printf("El dominio no existe")
			}
			waitc <- in.Ip
		}
	}()

	mensaje := dns.DomUpdate{
		OldDom: oldDom,
		NewDom: newDom,
	}

	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	response.CloseSend()
	<-waitc

	return true
}

func askDomainDns(domain string) string {
	var conn *grpc.ClientConn
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	ip := ips[rand.Intn(len(ips))]

	conn, err := grpc.Dial(ip, grpc.WithInsecure())

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
