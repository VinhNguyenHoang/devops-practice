package tom

import (
	"context"
	"log"

	"cs/internal/libs/database/mongodb"
)

type Server struct {
	mdb *mongodb.MongoDBImpl
}

// TODO: passin the db uri when init
func NewServer() (*Server, error) {
	dbimlp := &mongodb.MongoDBImpl{
		URI: "mongodb://user1:password1@mongodb-0.mongodb-headless.default.svc.cluster.local:27017,mongodb-1.mongodb-headless.default.svc.cluster.local:27017/stg",
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
	db := s.mdb.GetDatabase("stg")
	if db == nil {
		log.Printf("database is empty")
	}
	log.Println(db.Name())
	return nil
}
