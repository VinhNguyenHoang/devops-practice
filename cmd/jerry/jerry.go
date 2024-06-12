package jerry

import (
	"log"

	"cs/internal/jerry"

	"github.com/spf13/cobra"
)

var (
	JerryCmd = &cobra.Command{
		Use: "jerry",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("running jerry service")

			sv := jerry.Server{"0.0.0.0:30000"}
			if err := sv.Start(); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {

}
