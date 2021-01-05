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

func main() {

	var conn *grpc.ClientConn
	var reader *bufio.Reader
	conn, err := grpc.Dial("8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Conexion fallida: %s", err)
	}

	defer conn.Close()

	link := broker.NewBrokerHandlerClient(conn)

	ctx, err1 := context.WithTimeout(context.Background(), 10*time.Second)
	var domain string
	exito1 := 0
	for exito1 == 0 {
		fmt.Println("Escriba un dominio")
		fmt.Println("2) Descargar Libro")
		reader = bufio.NewReader(os.Stdin)
		fmt.Printf("Ingrese opcion: ")
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSuffix(opcion, "\n")
		domain = strings.TrimSuffix(opcion, "\r")
	}
	_ = domain
	// preguntar centralizado o distribuido
	var response broker.BrokerHandler_ConnectDomainClient
	response, err = link.ConnectDomain(ctx)
	_ = err1
	if err != nil {
		log.Printf("Conexion fallida al conectarse al datanode para mandar archivo: %s", err)
	}

	waitc := make(chan []broker.Response)

	go func() {
		for {
			in, err := response.Recv()
			if err == io.EOF {
				// waitc <- libros
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error al recibir un mensaje: %v", err)
			}
			log.Printf("El server retorna el siguiente mensaje: %v", in.Ip)
			// libros = append(libros, *in)
		}
	}()

}
