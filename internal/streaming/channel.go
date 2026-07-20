package streaming

import "telecom_anomaly_engine/internal/domain"

/*Global channel for metrics*/
var MetricsChannel = make(chan domain.Metric, 1000)
