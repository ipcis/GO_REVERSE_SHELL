package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Goroutine zum Lesen von Eingaben vom Server starten
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			// Eingabe vom Server lesen
			message, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Fehler beim Lesen der Eingabe:", err)
				return
			}

			// Eingabe an den Client senden
			_, err = conn.Write([]byte(message))
			if err != nil {
				log.Println("Fehler beim Senden der Daten an den Client:", err)
				return
			}
		}
	}()

	// Eingehende Daten vom Client lesen und ausgeben
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		receivedData := scanner.Text()
		fmt.Println(">", receivedData)

		// Hier kannst du die empfangenen Daten verarbeiten oder auf bestimmte Befehle reagieren

		// Beispiel: Wenn "quit" empfangen wird, die Verbindung schließen
		if receivedData == "quit" {
			fmt.Println("Verbindung geschlossen.")
			return
		}
	}

	if scanner.Err() != nil {
		log.Println("Fehler beim Lesen der Daten:", scanner.Err())
	}
}

func main() {
	// Server starten
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Fehler beim Starten des Servers:", err)
	}
	defer listener.Close()

	fmt.Println("Server läuft. Warte auf Verbindungen...")

	// Auf eingehende Verbindungen warten
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Fehler beim Akzeptieren der Verbindung:", err)
			continue
		}

		fmt.Println("Neue Verbindung hergestellt:", conn.RemoteAddr())

		// Verbindung in einem separaten Goroutine behandeln
		go handleConnection(conn)
	}
}
