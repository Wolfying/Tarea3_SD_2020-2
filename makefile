#················································Broker·························································
runBroker:
	protoc -I Broker/broker Broker/broker/broker.proto --go_out=plugins=grpc:./
	go run Broker/broker.go

compileBroker:
	protoc -I Broker/broker Broker/broker/broker.proto --go_out=plugins=grpc:./
#··················································································································
#················································Dns·························································
runDns:
	protoc -I Dns/dns Dns/dns/dns.proto --go_out=plugins=grpc:./
	go run Dns/dns.go

compileDns:
 	protoc -I Dns/dns Dns/dns/dns.proto --go_out=plugins=grpc:./
 
#··················································································································
#················································Cliente···························································
runCliente:
	go run Client/cliente.go
#··················································································································
#················································Admin·························································
runAdmin:
	protoc -I Admin/admin Admin/admin/admin.proto --go_out=plugins=grpc:./
	go run Admin/admin.go

compileAdmin:
 	protoc -I Admin/admin Admin/admin/admin.proto --go_out=plugins=grpc:./
 
#··················································································································
