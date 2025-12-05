package ovpnstats2influxdb

import (
	"encoding/json"
	"github.com/emibcn/ovpnstats"
)

// Metric represents metric events with all required information to write it into an InfluxDB
type Metric struct {
	Fields map[string]interface{}
	Tags   map[string]string
}

func createMetrics(clients []ovpnstats.ClientInfo, routes []ovpnstats.RoutingInfo) []Metric {
	return []Metric{Metric{map[string]interface{}{"clients": len(clients), "routes": len(routes)}, nil}}
}

func createMetricsClients(clients []ovpnstats.ClientInfo) []Metric {
	var inInterface []map[string]interface{}
	var metrics []Metric

	inrec, _ := json.Marshal(clients)
	json.Unmarshal(inrec, &inInterface)

	for _, obj := range inInterface {
		tags := map[string]string{
			"connection": obj["Name"].(string),
		}

		var met = Metric{obj, tags}
		metrics = append(metrics, met)
	}

	return metrics
}

func createMetricsRoutes(routes []ovpnstats.RoutingInfo) []Metric {
	var inInterface []map[string]interface{}
	var metrics []Metric

	inrec, _ := json.Marshal(routes)
	json.Unmarshal(inrec, &inInterface)

	for _, obj := range inInterface {
		tags := map[string]string{
			"connection": obj["CommonName"].(string),
		}

		var met = Metric{obj, tags}
		metrics = append(metrics, met)
	}

	return metrics
}
