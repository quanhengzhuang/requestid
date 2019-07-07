package requestid

import (
	"sync"

	"github.com/petermattis/goid"
)

var (
	requestIDs = map[int64]interface{}{}
	rwm        sync.RWMutex
)

// Set 设置一个 RequestID
func Set(ID interface{}) {
	goID := getGoID()
	rwm.Lock()
	defer rwm.Unlock()

	requestIDs[goID] = ID
}

// Get 返回设置的 RequestID
func Get() interface{} {
	goID := getGoID()
	rwm.RLock()
	defer rwm.RUnlock()

	return requestIDs[goID]
}

// Delete 删除设置的 RequestID
func Delete() {
	goID := getGoID()
	rwm.Lock()
	defer rwm.Unlock()

	delete(requestIDs, goID)
}

func getGoID() int64 {
	return goid.Get()
}
