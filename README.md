# Tarea3_SD_2020-2
Tarea 3 Sistemas Distribuidos 2020-2

Integrante:
-   Andrés Herera
    ROL: 201473593-6

Broker en dist141
Dns's en dist142, dist143, dist144

Ejecutar siguiente comando en todas las maquinas antes de ejecutar programas:
export PATH="$PATH:$(go env GOPATH)/bin" >> ~/.bashrc

make runBroker
make runDns
make runClient
make runAdmin

Cliente no cierra solo la conexion, cerrar con ctrl+c
No hay merge entre los servidores DNS, no se pudo :c
No hay reloj de vector.
Servidores reciben comandos, pero solo se puede crear un dominio.
Cliente si puede solicitar la ip perteneciente a un dominio.