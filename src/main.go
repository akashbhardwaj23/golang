package main

import (
	"goassignment/proto"
	"goassignment/src/cron"
	"goassignment/src/server"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcServer := grpc.NewServer()

	reportServer := server.NewReportServer()
	proto.RegisterReportServiceServer(grpcServer, reportServer)

	// Start listening on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("[%s] gRPC server starting on port 50051...", time.Now().Format(time.RFC3339))

	//Go routine for grpc
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Give some time for the server to start
	time.Sleep(2 * time.Second)
	log.Println("Server is running. Press Ctrl+C to stop.")

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("[%s] Did Not Connect to Grpc Server for Cron Job %v", time.Now().Format(time.RFC3339), err)
	}

	reportClient := proto.NewReportServiceClient(conn)

	// Start the cron scheduler
	scheduler := cron.NewScheduler(reportClient)
	scheduler.Start()
	defer scheduler.Stop()

	// Wait for interrupt signal to gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // Block until a signal is received

	log.Panicf("[%s] Shutting down server...", time.Now().Format(time.RFC3339))
	grpcServer.GracefulStop()
	log.Panicf("[%s] Server stopped.", time.Now().Format(time.RFC3339))
}
