package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"jerobas.com/yepee/repos"
	"jerobas.com/yepee/types"
	"jerobas.com/yepee/utils"
)

func PostRoute(w http.ResponseWriter, r *http.Request) {
	// if r.Header.Get("Content-Type") != "application/json" {
	// 	http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
	// }

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var newRoute types.RouteTypeDTO
	err := decoder.Decode(&newRoute)
	if err != nil {
		http.Error(w, "Invalid body format", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := utils.ValidateStruct(newRoute); err != nil {
		http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusBadRequest)
		return
	}

	localRoutes := repos.GetRoutesVariable()
	if _, ok := localRoutes[newRoute.Port]; ok {
		http.Error(w, "Route already exist", http.StatusBadRequest)
		return
	}

	newRouteFinal := types.RouteType{PID: 1, Subpath: newRoute.Subpath}
	localRoutes[newRoute.Port] = newRouteFinal

	err = utils.UpdateJSON("routes.json", localRoutes)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	repos.AssignRoutesVariable(localRoutes)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Route created")
}

func GetRoute(w http.ResponseWriter, r *http.Request) {
	localRoutes := repos.GetRoutesVariable()

	marshaledJSON, err := json.Marshal(localRoutes)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledJSON)
}
