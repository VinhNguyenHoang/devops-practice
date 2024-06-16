package jerry

import (
	"fmt"
	"log"

	"cs/internal/jerry"
	"cs/internal/libs/bootstrap"

	"github.com/spf13/cobra"
)

var (
	port                   string
	traceCollectorEndpoint string
	JerryCmd               = &cobra.Command{
		Use: "jerry",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("running jerry service", "port", port, "trace endpoint", traceCollectorEndpoint)

			httpServer := &bootstrap.HTTPServer{
				Name:                   "jerry",
				Address:                fmt.Sprintf("0.0.0.0:%s", port),
				TraceCollectorEndpoint: traceCollectorEndpoint,
			}

			if err := httpServer.Start(jerry.ExportEndpoints()); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	JerryCmd.Flags().StringVarP(&port, "servicePort", "p", "", "service port")
	JerryCmd.Flags().StringVarP(&traceCollectorEndpoint, "collector", "c", "", "metric collector address")
}
