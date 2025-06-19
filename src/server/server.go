package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	pb "goassignment/proto"
)

type ReportServer struct {
	pb.UnimplementedReportServiceServer
	reports map[string]string // InMemory reports Map
	mutx    sync.RWMutex
}

func NewReportServer() *ReportServer {
	return &ReportServer{
		reports: make(map[string]string),
	}
}

func (s *ReportServer) GenerateReport(ctx context.Context, req *pb.GenerateReportRequest) (*pb.GeneratedReportResponse, error) {
	userId := req.GetUserId()
	log.Printf("[%s] GenerateReport called for user: %s", time.Now().Format(time.RFC3339), userId)

	if userId == "" {
		log.Panicf("[%s] Error: Empty User Id", time.Now().Format(time.RFC3339))
		return &pb.GeneratedReportResponse{
			ReportId: "",
			Error:    "user_id cannot be empty",
		}, nil
	}

	reportId := fmt.Sprintf("report_%s_%d", req.GetUserId(), time.Now().Unix())

	// Simulate the report generation
	time.Sleep(100 * time.Millisecond)

	s.mutx.Lock()
	s.reports[reportId] = fmt.Sprintf("Report content for user %s generated at %s",
		userId, time.Now().Format(time.RFC3339))
	s.mutx.Unlock()

	log.Printf("[%s] Report generated successfully: %s", time.Now().Format(time.RFC3339), reportId)

	return &pb.GeneratedReportResponse{
		ReportId: reportId,
		Error:    "",
	}, nil

}

func (s *ReportServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Printf("[%s] HealthCheck called", time.Now().Format(time.RFC3339))

	s.mutx.RLock()
	reportCount := len(s.reports)
	s.mutx.RUnlock()

	status := fmt.Sprintf("healthy - %d reports in memory", reportCount)

	return &pb.HealthCheckResponse{
		Status: status,
	}, nil
}

func (s *ReportServer) GetReportCount() int {
	s.mutx.RLock()
	defer s.mutx.RUnlock()
	return len(s.reports)
}

func (s *ReportServer) GetReport(reportId string) (string, bool) {
	s.mutx.RLock()
	defer s.mutx.RUnlock()

	content, found := s.reports[reportId]

	return content, found
}
func (s *ReportServer) GetAllReports() map[string]string {
	s.mutx.RLock()
	defer s.mutx.RUnlock()

	reports := make(map[string]string)

	for k, v := range s.reports {
		reports[k] = v
	}

	return reports
}
