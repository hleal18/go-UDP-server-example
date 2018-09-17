/* UDPDaytimeServer
 */
 // Principal guía: https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/udp_datagrams.html
 // Paquete main, el punto de entrada principal
 // de ejecución.
package main

// Se importan los paquetes requeridos
import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Inicialización de un String con información del puerto.
	service := ":1200"
	// Asigna una dirección IP a la variable udpAddr
	// con puerto de acuerdo al servicio UDP.
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	// Se comprueba si hay algun error con la instrucción anterior.
	checkError(err)

	// Establece el mecanismo de escucha del servicio UDP.
	conn, err := net.ListenUDP("udp", udpAddr)
	// Se comprueba si la instrucción anterior provocó algún error.
	checkError(err)

	// Ciclo infinito
	for {
		// Para cada nuevo cliente (conexión entrante: conn)
		// se llama la funcion handleClient.
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {

	// Buffer de 512 bytes que lee la entrada.
	var buf [512]byte

	// Realiza el proceso de lectura de paquetes UDP
	// recibidos.
	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	decoded_instruction := string(buf[:n])

	// Imprime por consola la instruccion recibida.
	fmt.Println(decoded_instruction)

	if decoded_instruction == "encender_foco" {
		// Codigo a servicios cloud que gestionen
		// el encendido de un dispositivo.
	} else if decoded_instruction == "abrir_persiana" {
		// Codigo a servicios cloud que gestionen
		// el dispositivo encargado de abrir persianas.
	} else {
		// Ninguna instrucción es válida.
		conn.WriteToUDP([]byte("Instrucciones recibidas no son válidas\n"), addr)
	}

	// Obtiene la fecha actual y la pasa a un String.
	result := "Finalizacion de tareas\n"

	// Envía a la conexión actual (socket UDP) y
	// le envía un datagrama con la información de la petición realizada.
	conn.WriteToUDP([]byte(result), addr)

	// Se finaliza el proceso de gestión del cliente.
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
