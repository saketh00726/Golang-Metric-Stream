package api

import (
	"database/sql"
	"log"
	"net/http"
)

func StartRESTServer(db *sql.DB) {

	http.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT service_name, metric, timestamp FROM alerts")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		var result []map[string]interface{}

		for rows.Next() {

			var service string
			var metricMsg string
			var timestamp string

			err := rows.Scan(&service, &metricMsg, &timestamp)
			if err != nil {
				log.Println("Scan error:", err)
				continue
			}

			result = append(result, map[string]interface{}{
				"service": service,
				"metric":  metricMsg,
				"time":    timestamp,
			})
		}

		w.Header().Set("Content-Type", "text/html")

		w.Write([]byte(`
<html>
<head>
	<title>Alerts Dashboard</title>
	<style>
		body {
    font-family: Arial, sans-serif;
    background-color: #f4f6f8;
    margin: 0;
    padding: 20px;
}

.container {
    max-width: 1000px;
    margin: 0 auto;
    text-align: center;
}

table {
    border-collapse: collapse;
    margin: 20px auto;
    background-color: white;
}

th, td {
    padding: 10px 20px;
    border: 1px solid #ccc;
}

th {
    background-color: #007BFF;
    color: white;
}
		.container {
			text-align: center;
		}
		table {
			border-collapse: collapse;
			margin-top: 20px;
			background-color: white;
		}
		th, td {
			padding: 10px 20px;
			border: 1px solid #ccc;
		}
		th {
			background-color: #007BFF;
			color: white;
		}
	</style>
</head>
<body>
	<div class="container">
		<h2>Alerts Table</h2>
		<table>
			<tr>
				<th>Service</th>
				<th>Metric</th>
				<th>Time</th>
			</tr>
`))

		for _, alert := range result {
			row := "<tr><td>" + alert["service"].(string) +
				"</td><td>" + alert["metric"].(string) +
				"</td><td>" + alert["time"].(string) +
				"</td></tr>"
			w.Write([]byte(row))
		}

		w.Write([]byte(`
		</table>
	</div>
</body>
</html>
`))
	})

	log.Println("REST server running on :8080")
	http.ListenAndServe(":8080", nil)
}
