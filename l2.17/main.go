package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// создаем флаг таймаут
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	// парсим аргументы командной строки
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("Usage: <host> <port>")
	}

	host := args[0]
	port := args[1]
	address := net.JoinHostPort(host, port)

	// соединение с сервером с таймаутом
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Printf("Failed to connect %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	doneChan := make(chan struct{})

	// чтение из сокета -> stdout
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil && err != io.EOF {
			log.Println("Error reading from server:", err)
		}
		doneChan <- struct{}{}
	}()

	// чтение из stdin -> сокет
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil && err != io.EOF {
			log.Println("Error writing to server:", err)
		}
		doneChan <- struct{}{}
	}()

	// обработка сигналов
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signalChan:
		fmt.Println("\nClosing connection (signal received)...")
	case <-doneChan:
		fmt.Println("\nConnection closed.")
	}
}
