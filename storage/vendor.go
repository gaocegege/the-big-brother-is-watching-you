package storage

import (
	"github.com/gaocegege/the-big-brother-is-watching-you/common"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// VendorCollectionManager is the manager for vendor in mongo
type VendorCollectionManager struct {
	manager *CollectionManager
}

// NewVendorCollectionManager returns a new VendorCollectionManager object
func NewVendorCollectionManager(s *mgo.Session) *VendorCollectionManager {
	return &VendorCollectionManager{
		manager: newCollectionManager(s, common.DefaultDBName, common.VendorCollectionName),
	}
}

// NewVendorDocument generates a new vendor document
func (m *VendorCollectionManager) NewVendorDocument(v *common.Vendor) (string, error) {
	v.VendorID = uuid.NewV4().String()
	_, err := m.manager.col.Upsert(bson.M{"_id": v.VendorID}, v)
	return v.VendorID, err
}

// FindVendorByID finds a vendor entity by ID.
func (m *VendorCollectionManager) FindVendorByID(vID string) (*common.Vendor, error) {
	vendor := &common.Vendor{}
	err := m.manager.col.Find(bson.M{"_id": vID}).One(vendor)
	return vendor, err
}

// FindVendorByHost finds a vendor entity by host.
func (m *VendorCollectionManager) FindVendorByHost(host string) (*common.Vendor, error) {
	vendor := &common.Vendor{}
	err := m.manager.col.Find(bson.M{"host": host}).One(vendor)
	return vendor, err
}

// UpdateVendorDocument updates a vendor entirely.
func (m *VendorCollectionManager) UpdateVendorDocument(vID string, vendor common.Vendor) error {
	filter := bson.M{"_id": vID}
	change := mgo.Change{
		Update: bson.M{"$set": vendor},
	}
	_, err := m.manager.col.Find(filter).Apply(change, &vendor)
	return err
}

// AddNewRecord adds a new record (record ID) to a given service.
func (m *VendorCollectionManager) AddNewRecord(vID string, rID string) error {
	change := mgo.Change{
		Update: bson.M{"$push": bson.M{"records": rID}},
	}

	_, err := m.manager.col.Find(bson.M{"_id": vID}).Apply(change, nil)
	return err
}

// AddNewRecords adds new records (record ID) to a given service.
func (m *VendorCollectionManager) AddNewRecords(vID string, rID []string) error {
	change := mgo.Change{
		Update: bson.M{"$push": bson.M{"records": rID}},
	}

	_, err := m.manager.col.Find(bson.M{"_id": vID}).Apply(change, nil)
	return err
}
