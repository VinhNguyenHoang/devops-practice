package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"cs/internal/tom/collectors"
)

func StartCollectors() error {
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
