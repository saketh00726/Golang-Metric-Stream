package alert

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func worker(db *sql.DB) {
	for metric := range MetricChannel {

		log.Println("Worker received:", metric.ServiceName, metric.CpuUsage, metric.MemoryUsage)

		if metric.CpuUsage > 80 {
			insertAlert(db, metric.ServiceName, "CPU > 80%")
		}

		if metric.MemoryUsage > 90 {
			insertAlert(db, metric.ServiceName, "Memory > 90%")
		}
	}
}

func insertAlert(db *sql.DB, serviceName, metricMsg string) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := db.ExecContext(
		ctx,
		"INSERT INTO alerts(service_name, metric, timestamp) VALUES (?, ?, ?)",
		serviceName,
		metricMsg,
		time.Now(),
	)

	if err != nil {
		log.Println("Insert failed:", err)
		return
	}

	log.Printf("[ALERT STORED] %s -> %s\n", serviceName, metricMsg)
}
