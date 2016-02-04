package storage

import "gopkg.in/mgo.v2"

// CollectionManager is the m
type CollectionManager struct {
	s   *mgo.Session
	col *mgo.Collection
}

func newCollectionManager(s *mgo.Session, DBName string, collectionName string) *CollectionManager {
	return &CollectionManager{
		s:   s,
		col: s.DB(DBName).C(collectionName),
	}
}