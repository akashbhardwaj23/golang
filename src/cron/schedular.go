package cron

import (
	"context"
	"goassignment/src/server"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron            *cron.Cron
	reportServer    *server.ReportServer
	predefinedUsers []string
}

func NewScheduler(reportServer *server.ReportServer) *Scheduler {
	return &Scheduler{
		cron:            cron.New(cron.WithSeconds()),
		reportServer:    reportServer,
		predefinedUsers: []string{"user1", "user2", "user3", "user4", "user5"},
	}
}

func (s *Scheduler) Start() {
	log.Println("Starting cron Scheduler")

	_, err := s.cron.AddFunc("*/10 * * * * *", s.generateReportsJob)

	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	s.cron.Start()
	log.Println("Cron Scheduler Started - will generate reports every 10 seconds")

}

func (s *Scheduler) Stop() {
	log.Printf("Stopping Cron Scheduler")
	s.cron.Stop()
	log.Printf("Cron Scheduler Stopped")
}

func (s *Scheduler) generateReportsJob() {
	log.Printf("[%s] Cron Job Trigerred - Generating Reports For PredifinedUser", time.Now().Format(time.RFC3339))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, userId := range s.predefinedUsers {
		req := &server.GeneratedReportRequest{
			UserId: userId,
		}

		resp, err := s.reportServer.GenerateReport(ctx, req)
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

	reportCount := s.reportServer.GetReportCount()
	log.Printf("[%s] Cron Job Completed - Total Report In Memory : %d", time.Now().Format(time.RFC3339), reportCount)

}
