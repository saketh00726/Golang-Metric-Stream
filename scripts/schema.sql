CREATE TABLE metrics (
    id INT AUTO_INCREMENT PRIMARY KEY,
    service_name VARCHAR(100),
    cpu_usage INT,
    memory_usage INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE alerts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    service_name VARCHAR(100),
    alert_type VARCHAR(50),
    value INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);