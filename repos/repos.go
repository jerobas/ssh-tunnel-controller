package repos

import (
	"log"
	"maps"
	"sync"

	"jerobas.com/yepee/types"
	"jerobas.com/yepee/utils"
)

var (
	mu     sync.RWMutex
	routes types.RoutesType
)

func GetRoutesVariable() types.RoutesType {
	localRoutes := make(types.RoutesType)

	mu.RLock()
	maps.Copy(localRoutes, routes)
	mu.RUnlock()

	return localRoutes
}

func LoadRoutesVariable() {
	loadedRoutes, err := utils.ReadJSON("routes.json")
	if err != nil {
		log.Fatalf("Failed to load routes: %v", err)
	}

	mu.Lock()
	routes = loadedRoutes
	mu.Unlock()
}

func AssignRoutesVariable(value types.RoutesType) {
	mu.Lock()
	routes = value
	mu.Unlock()
}
