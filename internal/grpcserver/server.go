package grpcserver

import (
	"database/sql"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"metrics-app/internal/alert"
	pb "metrics-app/proto"
)

type Server struct {
	pb.UnimplementedMetricsServiceServer
	DB *sql.DB
}

func StartGRPCServer(db *sql.DB) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMetricsServiceServer(grpcServer, &Server{DB: db})
	reflection.Register(grpcServer)

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

func (s *Server) StreamMetrics(stream pb.MetricsService_StreamMetricsServer) error {

	for {
		metric, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Ack{Message: "Stream completed"})
		}
		if err != nil {
			return err
		}
		log.Println("gRPC received:", metric.ServiceName, metric.CpuUsage, metric.MemoryUsage)

		_, err = s.DB.Exec(
			"INSERT INTO metrics(service_name, cpu_usage, memory_usage) VALUES (?, ?, ?)",
			metric.ServiceName,
			metric.CpuUsage,
			metric.MemoryUsage,
		)
		if err != nil {
			log.Println("Metric insert error:", err)
		}

		alert.MetricChannel <- metric
	}
}
