package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "metrics-app/proto"

	"google.golang.org/grpc"
)

func startAgent(agentName string) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewMetricsServiceClient(conn)

	stream, err := client.StreamMetrics(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		cpu := rand.Intn(100)
		mem := rand.Intn(100)

		log.Printf("[%s] Sending CPU=%d MEM=%d\n", agentName, cpu, mem)

		err := stream.Send(&pb.Metric{
			ServiceName: agentName,
			CpuUsage:    int32(cpu),
			MemoryUsage: int32(mem),
		})
		if err != nil {
			log.Println("Send error:", err)
			return
		}

		time.Sleep(3 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 100; i++ {
		go startAgent(fmt.Sprintf("Agent-%d", i))
	}

	select {}
}
