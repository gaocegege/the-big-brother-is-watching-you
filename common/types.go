package common

import "time"

const (
	// DefaultDBName is the name of the default database
	DefaultDBName         	string = "dev"
	// VendorCollectionName is the name of the collection
	VendorCollectionName 	string = "VendorCollection"
)

// Vendor is the object represent a vendor
type Vendor struct {
	VendorID	string		`bson:"_id,omitempty" json:"_id,omitempty"`
	Host 		string 		`bson:"host,omitempty" json:"host,omitempty"`
	LastDate 	time.Time 	`bson:"last_time,omitempty" json:"last_time,omitempty"`
	Records		[]Record	`bson:"records,omitempty" json:"records,omitempty"`
}

// Record is a object, which means that ttb has found a update
type Record struct {
	RecordID	string		`bson:"_id,omitempty" json:"_id,omitempty"`
	VendorID	string		`bson:"vendor_id,omitempty" json:"vendor_id,omitempty"`
	Content		string		`bson:content,omitempty" json:"content,omitempty"`
	CreatDate	time.Time	`bson:create_time,omitempty" json:"create_time,omitempty"`
	URL			string 		`bson:"url,omitempty" json:"url,omitempty"`
}