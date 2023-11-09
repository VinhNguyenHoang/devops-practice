package tom

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"cs/internal/libs/util"
	"cs/internal/tom/collectors"
)

type Server struct {
	// address with port
	Address string
}

func (s *Server) Start() error {
	err := util.DoWithInterval(time.Second*5, startCollectors)

	e := <-err
	return e
}

func startCollectors() error {
	fmt.Println("starting collector")
	cellPhonesCollector := collectors.NewCellphonesCollector("")

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
