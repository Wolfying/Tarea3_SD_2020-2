package main

import (
	"Tarea3/Broker/broker"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func CreateDomain(dom string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist141:8180", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}
	defer conn.Close()
	link := broker.NewBrokerHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10000*time.Second)
	var response broker.BrokerHandler_CreateDomainClient
	response, err = link.CreateDomain(ctx)
	_ = err1
	if err != nil {
		log.Printf("Conexion fallida al conectarse al datanode para mandar archivo: %s", err)
	}
	waitc := make(chan broker.Response)

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				waitc <- *in
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al recibir un mensaje: %v", err)
			}

			if in.Ip == " " {
				log.Printf("No existe una IP asociada al dominio ingresado en el servidor utilizado")
			} else {
				log.Printf("El server retorna el siguiente mensaje: %v", in.Ip)
			}
		}
	}()

	mensaje := broker.Domcreate{
		Domname: dom,
	}
	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	// response.CloseSend()
	<-waitc
}

func UpdateDomain(oldDom string, newDom string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist141:8180", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}
	defer conn.Close()
	link := broker.NewBrokerHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10000*time.Second)
	var response broker.BrokerHandler_UpdateDomainClient
	response, err = link.UpdateDomain(ctx)
	_ = err1
	if err != nil {
		log.Printf("Conexion fallida al conectarse al datanode para mandar archivo: %s", err)
	}
	waitc := make(chan broker.Response)

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				waitc <- *in
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al recibir un mensaje: %v", err)
			}

			if in.Ip == " " {
				log.Printf("No existe una IP asociada al dominio ingresado en el servidor utilizado")
			} else {
				log.Printf("El server retorna el siguiente mensaje: %v", in.Ip)
			}
		}
	}()

	mensaje := broker.DomUpdate{
		OldDom: oldDom,
		NewDom: newDom,
	}
	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	// response.CloseSend()
	<-waitc
}

func DeleteDomain(dom string) {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist141:8180", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}
	defer conn.Close()
	link := broker.NewBrokerHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10000*time.Second)
	var response broker.BrokerHandler_ConnectDomainClient
	response, err = link.DeleteDomain(ctx)
	_ = err1
	if err != nil {
		log.Printf("Conexion fallida al conectarse al datanode para mandar archivo: %s", err)
	}
	waitc := make(chan broker.Response)

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				waitc <- *in
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al recibir un mensaje: %v", err)
			}

			if in.Ip == " " {
				log.Printf("No existe una IP asociada al dominio ingresado en el servidor utilizado")
			} else {
				log.Printf("El server retorna el siguiente mensaje: %v", in.Ip)
			}
		}
	}()

	mensaje := broker.Dom{
		Domname: dom,
	}
	if err := response.Send(&mensaje); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	// response.CloseSend()
	<-waitc
}

func main() {

	var reader *bufio.Reader
	var domain string
	var funcion string
	exito1 := 0
	for exito1 == 0 {
		fmt.Printf("Ingrese comando: ")
		reader = bufio.NewReader(os.Stdin)
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSuffix(opcion, "\n")
		opcion = strings.TrimSuffix(opcion, "\r")
		cosas := strings.Split(opcion, " ")
		funcion = cosas[0]
		if funcion == "Create" {
			domain = cosas[1]
			CreateDomain(domain)
			exito1 = 1
		} else if funcion == "Update" {
			domain = cosas[1]
			if len(cosas) >= 3 {
				newDom := cosas[2]
				UpdateDomain(domain, newDom)
			}
			exito1 = 1
		} else if funcion == "Delete" {
			domain = cosas[1]
			DeleteDomain(domain)
			exito1 = 1
		} else {
			fmt.Println("Ingrese opcion una opción válida: ")
		}

		// exito1 = 1
	}
	_ = domain
	// preguntar centralizado o distribuido

}
