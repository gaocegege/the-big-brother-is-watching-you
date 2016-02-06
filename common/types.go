package common

import "time"

const (
	// DefaultDBName is the name of the default database
	DefaultDBName string = "dev"
	// VendorCollectionName is the name of the vendor collection
	VendorCollectionName string = "VendorCollection"
	// RecordCollectionName is the name of the record collection
	RecordCollectionName string = "RecordCollection"

	// GithubOrigin is the host name of github
	GithubOrigin = "github.com"
	// MockOrigin is the host name of mock
	MockOrigin = "mock.me"
)

// Vendor is the object represent a vendor
type Vendor struct {
	VendorID string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Host     string    `bson:"host,omitempty" json:"host,omitempty"`
	LastTime time.Time `bson:"last_time,omitempty" json:"last_time,omitempty"`
	Records  []Record  `bson:"records,omitempty" json:"records,omitempty"`
}

// Record is a object, which means that ttb has found a update
type Record struct {
	RecordID   string    `bson:"_id,omitempty" json:"_id,omitempty"`
	VendorID   string    `bson:"vendor_id,omitempty" json:"vendor_id,omitempty"`
	Content    string    `bson:content,omitempty" json:"content,omitempty"`
	CreateTime time.Time `bson:create_time,omitempty" json:"create_time,omitempty"`
	URL        string    `bson:"url,omitempty" json:"url,omitempty"`
}
