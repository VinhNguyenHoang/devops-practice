package tom

import (
	"context"
	"time"

	"cs/internal/libs/database/mongodb"
	"cs/internal/libs/util"
	"cs/internal/tom/handler"
)

type Server struct {
	mdb *mongodb.MongoDBImpl
}

// TODO: passin the db uri when init
func NewServer() (*Server, error) {
	dbimlp := &mongodb.MongoDBImpl{
		URI: "mongodb://root@localhost:27017",
	}

	err := dbimlp.Init(context.Background())
	if err != nil {
		return nil, err
	}
	s := &Server{
		mdb: dbimlp,
	}

	return s, nil
}

func (s *Server) Start() error {
	err := util.DoWithInterval(time.Second*5, handler.StartCollectors)

	e := <-err
	return e
}
