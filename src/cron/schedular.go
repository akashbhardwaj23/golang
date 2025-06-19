package cron

import (
	"context"
	"log"
	"time"

	pb "goassignment/proto"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron            *cron.Cron
	client          pb.ReportServiceClient
	predefinedUsers []string
}

func NewScheduler(client pb.ReportServiceClient) *Scheduler {
	return &Scheduler{
		cron:            cron.New(cron.WithSeconds()),
		client:          client,
		predefinedUsers: []string{"user1", "user2", "user3", "user4", "user5"},
	}
}

func (s *Scheduler) Start() {
	log.Printf("[%s] Starting cron Scheduler", time.Now().Format(time.RFC3339))

	// Runs every 10 seconds and generate a Report
	_, err := s.cron.AddFunc("*/10 * * * * *", s.generateReportsJob)

	if err != nil {
		log.Fatalf("[%s]Failed to add cron job: %v", time.Now().Format(time.RFC3339), err)
	}

	_, err = s.cron.AddFunc("*/5 * * * * *", s.generateHealthReport)
	if err != nil {
		log.Fatalf("[%s] Failed to add health check cron job: %v", time.Now().Format(time.RFC3339), err)
	}

	s.cron.Start()
	log.Printf("[%s] Cron Scheduler Started - will generate reports every 10 seconds", time.Now().Format(time.RFC3339))

}

func (s *Scheduler) Stop() {
	log.Printf("[%s] Stopping Cron Scheduler", time.Now().Format(time.RFC3339))
	s.cron.Stop()
	log.Printf("[%s] Cron Scheduler Stopped", time.Now().Format(time.RFC3339))
}

func (s *Scheduler) generateReportsJob() {
	log.Printf("[%s] Cron Job Trigerred - Generating Reports For PredifinedUser", time.Now().Format(time.RFC3339))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, userId := range s.predefinedUsers {
		req := &pb.GenerateReportRequest{
			UserId: userId,
		}

		// Made a gRpc call using the client
		resp, err := s.client.GenerateReport(ctx, req)
		if err != nil {
			log.Printf("[%s] Error Generating Reports For Users %s: %v", time.Now().Format(time.RFC3339), userId, err)
			continue
		}

		if resp.Error != "" {
			log.Printf("[%s] Report Generation Failed for User %s : %s", time.Now().Format(time.RFC3339), userId, resp.Error)
			continue
		}

		log.Printf("[%s] Cron Job - Reports Generated for User %s : %s", time.Now().Format(time.RFC3339), userId, resp.Error)
	}

	// reportCount := s.reportServer.GetReportCount()
	// log.Printf("[%s] Cron Job Completed - Total Report In Memory : %d", time.Now().Format(time.RFC3339), reportCount)

}

func (s *Scheduler) generateHealthReport() {
	log.Printf("[%s] Generating Health Report", time.Now().Format(time.RFC3339))

	// Function should be running till this time
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	req := &pb.HealthCheckRequest{}
	defer cancel()

	healthCheckResponse, err := s.client.HealthCheck(ctx, req)

	if err != nil {
		log.Printf("[%s] Cron Job Execution Errored", time.Now().Format(time.RFC3339))
		return
	}

	log.Printf("[%s] Cron Job for Health Check Generated Report [%s]", time.Now().Format(time.RFC3339), healthCheckResponse.Status)
}
