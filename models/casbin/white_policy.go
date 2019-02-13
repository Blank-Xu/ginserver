package casbin

import (
	"sync"
)

var (
	_whitePolicy     = make(map[string]bool)
	_whitePolicyLock sync.RWMutex
)

func setWhitePolicy(path, method string) {
	path = getWhitePolicyPath(path, method)
	_whitePolicyLock.Lock()
	_whitePolicy[path] = true
	_whitePolicyLock.Unlock()
}

func getWhitePolicy(path, method string) (has bool) {
	path = getWhitePolicyPath(path, method)
	_whitePolicyLock.RLock()
	has = _whitePolicy[path]
	_whitePolicyLock.RUnlock()
	return
}

func getWhitePolicyPath(path, method string) string {
	return path + ":" + method
}
