package tom

import (
	"encoding/json"
	"os"

	"cs/internal/tom/collectors"
)

type Server struct {
	// address with port
	Address string
}

func (s *Server) Start() error {

	cellPhonesCollector := collectors.NewCellphonesCollector()

	err := cellPhonesCollector.RunCollect()
	if err != nil {
		return err
	}

	data := cellPhonesCollector.GetCollection()

	b, _ := json.Marshal(data)

	err = os.WriteFile("tmp/data.txt", b, 0777)
	if err != nil {
		panic(err)
	}

	return nil
}
