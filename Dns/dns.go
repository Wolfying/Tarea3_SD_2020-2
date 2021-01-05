package main

import (
	"Tarea3/Dns/dns"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

// ServerDns ...
type ServerDns struct { //nolint
}

type Domain struct {
	domain string
	ip     string
}

func (dom *Domain) getDomain() string {
	return dom.domain
}
func (dom *Domain) getIp() string {
	return dom.ip
}

func (dom *Domain) setDomain(cosa string) {
	dom.domain = cosa
}
func (dom *Domain) setIp(cosa string) {
	dom.ip = cosa
}

// GetDomain ...
func (*ServerDns) GetDomain(incomestream dns.DnsHandler_GetDomainServer) error {

	for {
		in, err := incomestream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("imprimo algo antes de return err: %v", err)
			return err
		}
		// _ = in
		domcompare := in.Domname
		archivo := "ZF.txt"
		file, err := os.OpenFile(archivo, os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("Error abriendo el archivo: %s", err)
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var line string
		// var returnip string
		var doms []Domain
		returnip := " "
		for scanner.Scan() {
			line = strings.TrimSuffix(scanner.Text(), "\n")
			line = strings.TrimSuffix(line, "\r")
			splitline := strings.Split(line, " ")
			dom := Domain{
				domain: splitline[0],
				ip:     splitline[3],
			}
			if domcompare == splitline[0] {
				returnip = splitline[3]
				break
			}
			doms = append(doms, dom)
		}
		mensaje := &dns.Response{
			Ip: returnip}
		if err := incomestream.Send(mensaje); err != nil {
			log.Printf("Error en propuesta  %s", err)
			return err
		}
	}
	return nil
}

func main() {
	server := ServerDns{}
	puerto := ":8181"
	lis, err := net.Listen("tcp", puerto)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", puerto, err)
	}
	grpcServer := grpc.NewServer()

	dns.RegisterDnsHandlerServer(grpcServer, &server)
	fmt.Printf("Server Iniciado en el puerto %s \n", puerto)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", puerto, err)
	}
}
