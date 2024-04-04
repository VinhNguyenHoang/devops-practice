package tom

import (
	"log"

	core "cs/internal/tom"

	"github.com/spf13/cobra"
)

var (
	TomCmd = &cobra.Command{
		Use: "tom",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Print("running tom service")

			sv, err := core.NewServer()
			if err != nil {
				return err
			}

			return sv.Start()
		},
	}
)

func init() {

}
