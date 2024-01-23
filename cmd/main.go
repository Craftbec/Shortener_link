package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Craftbec/Shortener_link/config"
	"github.com/Craftbec/Shortener_link/internal/grcpserver"
	"github.com/Craftbec/Shortener_link/internal/httpserver"
	"github.com/Craftbec/Shortener_link/internal/storage"
)

func main() {
	storageTypeFlag := flag.String("storage", "in-memory", "Storage selection (postgres or in-memory")
	flag.Parse()
	conf, err := config.NewConfigStruct()
	if err != nil {
		log.Fatalln(err)
	}
	var storeData storage.Storage
	if *storageTypeFlag == "postgres" {
		storeData, err = storage.NewDB(conf)
		if err != nil {
			log.Fatal(err)
		}
	} else if *storageTypeFlag == "in-memory" {
		storeData = storage.NewInMemory()
	} else {
		log.Fatal("Invalid storage type. Available: in-memory or postgres")
	}
	done := make(chan struct{})
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-stop
		cancel()
		close(done)
	}()
	go func() {
		err := grcpserver.GRPCServer(ctx, storeData, conf)
		if err != nil {
			log.Fatalf("GRPC server error: %v\n", err)
		}
	}()
	go func() {
		err := httpserver.HTTPServer(ctx, conf)
		if err != nil {
			log.Fatalf("HTTP server error: %v\n", err)
		}
	}()
	<-done
	storeData.GracefulStopDB()
	log.Println("Shutting down gracefully")
}
