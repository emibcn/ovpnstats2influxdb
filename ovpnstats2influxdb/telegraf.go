package ovpnstats2influxdb

import (
	"fmt"

	"github.com/emibcn/ovpnstats"
)

// RunTelegraf prints Telegraf-compatible metric-output to stdout
func RunTelegraf(path string) error {
	clients, routes, err := ovpnstats.ParseStatusFile(path)
	if err != nil {
		return err
	}

	metrics := createMetrics(clients, routes)

	// convert metrics to influxdb line protocol
	points, err := createBatchPoints("openvpn", metrics)
	if err != nil {
		return err
	}

	// output line protocol lines
	for _, point := range points {
		fmt.Println(point.String())
	}

	metrics_clients := createMetricsClients(clients)

	// convert metrics to influxdb line protocol
	points_clients, err := createBatchPoints("openvpn", metrics_clients)
	if err != nil {
		return err
	}

	// output line protocol lines
	for _, point := range points_clients {
		fmt.Println(point.String())
	}

	metrics_routes := createMetricsRoutes(routes)

	// convert metrics to influxdb line protocol
	points_routes, err := createBatchPoints("openvpn", metrics_routes)
	if err != nil {
		return err
	}

	// output line protocol lines
	for _, point := range points_routes {
		fmt.Println(point.String())
	}

	return nil
}
