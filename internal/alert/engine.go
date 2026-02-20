package alert

import (
	"database/sql"
	pb "metrics-app/proto"
)

var MetricChannel chan *pb.Metric

func StartAlertEngine(db *sql.DB, workerCount int) {
	MetricChannel = make(chan *pb.Metric, 1000)

	for i := 0; i < workerCount; i++ {
		go worker(db)
	}
}
