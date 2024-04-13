package tom

import (
	"context"
	"time"

	"cs/internal/libs/database/mongodb"
	"cs/internal/libs/util"
)

var Collection = "blocks"

type Block struct {
	BlockInt  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Server struct {
	mdb *mongodb.MongoDBImpl
}

// TODO: passin the db uri when init
func NewServer() (*Server, error) {
	dbimlp := &mongodb.MongoDBImpl{
		URI:    "mongodb://user1:password1@mongodb-0.mongodb-headless.default.svc.cluster.local:27017,mongodb-1.mongodb-headless.default.svc.cluster.local:27017",
		DBName: "stg",
	}

	err := dbimlp.Init(context.Background(), Collection)
	if err != nil {
		return nil, err
	}
	s := &Server{
		mdb: dbimlp,
	}

	return s, nil
}

func (s *Server) Start() error {
	err := util.DoWithInterval(5*time.Second, func() error {
		block := &Block{
			BlockInt:  int(time.Now().Unix()),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		err := s.mdb.Insert(context.TODO(), Collection, block)
		return err
	})

	e := <-err
	if e != nil {
		return e
	}
	return nil
}
