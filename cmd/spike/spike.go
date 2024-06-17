package spike

import (
	"fmt"
	"log"

	"cs/internal/libs/bootstrap"
	"cs/internal/spike"

	"github.com/spf13/cobra"
)

var (
	port                   string
	traceCollectorEndpoint string
	SpikeCmd               = &cobra.Command{
		Use: "spike",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("running spike service", "port", port, "trace endpoint", traceCollectorEndpoint)

			httpServer := &bootstrap.HTTPServer{
				Name:                   "spike",
				Address:                fmt.Sprintf("0.0.0.0:%s", port),
				TraceCollectorEndpoint: traceCollectorEndpoint,
			}

			if err := httpServer.Start(spike.ExportEndpoints()); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	SpikeCmd.Flags().StringVarP(&port, "servicePort", "p", "", "service port")
	SpikeCmd.Flags().StringVarP(&traceCollectorEndpoint, "collector", "c", "", "metric collector address")
}
