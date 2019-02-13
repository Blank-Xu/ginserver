package casbin

import (
	"sync"
)

var (
	_whitePolicy     = make([]string, 0)
	_whitePolicyLock sync.RWMutex
)

func setWhitePolicy(path, method string) {
	path = getWhitePolicyPath(path, method)
	_whitePolicyLock.Lock()
	_whitePolicy = append(_whitePolicy, path)
	_whitePolicyLock.Unlock()
}

func getWhitePolicy(path, method string) bool {
	path = getWhitePolicyPath(path, method)
	_whitePolicyLock.RLock()
	for _, value := range _whitePolicy {
		if value == path {
			return true
		}
	}
	_whitePolicyLock.RUnlock()
	return false
}

func getWhitePolicyPath(path, method string) string {
	return method + path
}
