package worker

import (
	"github.com/gaocegege/the-big-brother-is-watching-you/source"
	"github.com/gaocegege/the-big-brother-is-watching-you/storage"
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

}
