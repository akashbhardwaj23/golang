package main

import (
	"goassignment/src/cron"
	"goassignment/src/server"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	reportServer := server.NewReportServer()

	// Start listening on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server starting on port 50051...")

	// Start the cron scheduler
	scheduler := cron.NewScheduler(reportServer)
	scheduler.Start()
	defer scheduler.Stop()

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	log.Println("Server is running. Press Ctrl+C to stop.")

	// Wait for interrupt signal to gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Shutting down server...")
	s.GracefulStop()
	log.Println("Server stopped.")
}
