package storage

import (
	"github.com/gaocegege/the-big-brother-is-watching-you/common"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RecordCollectionManager is the manager for Record in mongo
type RecordCollectionManager struct {
	manager *CollectionManager
}

// NewRecordCollectionManager returns a new RecordCollectionManager object
func NewRecordCollectionManager(s *mgo.Session) *RecordCollectionManager {
	return &RecordCollectionManager{
		manager: newCollectionManager(s, common.DefaultDBName, common.RecordCollectionName),
	}
}

// NewRecordDocument generates a new Record document
func (m *RecordCollectionManager) NewRecordDocument(r *common.Record) (string, error) {
	r.RecordID = uuid.NewV4().String()
	_, err := m.manager.col.Upsert(bson.M{"_id": r.RecordID}, r)
	return r.RecordID, err
}

// FindRecordByID finds a Record entity by ID.
func (m *RecordCollectionManager) FindRecordByID(rID string) (*common.Record, error) {
	Record := &common.Record{}
	err := m.manager.col.Find(bson.M{"_id": rID}).One(Record)
	return Record, err
}
