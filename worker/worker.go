package worker

import (
	"log"
	"time"

	"github.com/gaocegege/the-big-brother-is-watching-you/common"
	"github.com/gaocegege/the-big-brother-is-watching-you/source"
	"github.com/gaocegege/the-big-brother-is-watching-you/storage"
)

const (
	fetchTime = 240
)

// Worker fetch all the records and put them to db
type Worker struct {
	sm *source.Manager
	vm *storage.VendorCollectionManager
	rm *storage.RecordCollectionManager
}

// NewWorker returns a new Worker object
func NewWorker(sm *source.Manager, vm *storage.VendorCollectionManager, rm *storage.RecordCollectionManager) *Worker {
	return &Worker{
		sm: sm,
		vm: vm,
		rm: rm,
	}
}

// Work now
func (w *Worker) Work() {
	for _, source := range w.sm.GetSources() {
		var lastTime time.Time

		vendor, err := w.vm.FindVendorByHost(source.GetHostName())

		// if the vendor doesn't exist, alloc a new document
		if err != nil {
			log.Print(err)

			// trick: add a negative duration
			lastTime = time.Now().Add(-time.Duration(fetchTime) * time.Hour)

			vendorDoc := &common.Vendor{
				Host:     source.GetHostName(),
				LastTime: time.Now(),
			}
			vendorID, err := w.vm.NewVendorDocument(vendorDoc)
			if err != nil {
				log.Fatal(err)
			}

			vendor, err = w.vm.FindVendorByID(vendorID)
			if err != nil {
				log.Fatal(err)
			}
		}

		// else set the last time
		lastTime = vendor.LastTime
		vendor.LastTime = time.Now()
		w.vm.UpdateVendorDocument(vendor.VendorID, *vendor)

		// check out the last date, and fetch origin
		records, err := source.FetchFromOrigin(lastTime)
		if err != nil {
			log.Fatal(err)
		}

		// store
		for _, record := range records {
			w.rm.NewRecordDocument(&record)
			w.vm.AddNewRecord(vendor.VendorID, record.RecordID)
		}

		log.Printf("New Records from %s has been fetched", source.GetHostName())
	}
}
