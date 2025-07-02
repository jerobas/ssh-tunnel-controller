package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/jerobas/territo/repos"
	"github.com/jerobas/territo/types"
	"github.com/jerobas/territo/utils"
)

var (
	tunnelMu sync.RWMutex
	validate = validator.New()
)

func MainRoute(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handleGet(w)
	case http.MethodPost:
		handlePost(w, r.Body)
	case http.MethodDelete:
		handleDelete(w, r.Body)
	default:
		utils.Error(w, http.StatusMethodNotAllowed, "Unsupported method")
	}
}

func handlePost(w http.ResponseWriter, body io.ReadCloser) {
	defer body.Close()
	var data types.CreateTunnelDTO

	if err := json.NewDecoder(body).Decode(&data); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if err := validate.Struct(data); err != nil {
		utils.Error(w, http.StatusBadRequest, fmt.Sprintf("Validation failed: %v", err))
		return
	}

	tunnelMu.Lock()
	defer tunnelMu.Unlock()

	if !repos.CreateTunnel(data) {
		utils.Error(w, http.StatusConflict, "Tunnel already exists")
		return
	}

	utils.Success(w, http.StatusCreated, "Tunnel created successfully")
}

func handleGet(w http.ResponseWriter) {
	tunnelMu.RLock()
	defer tunnelMu.RUnlock()

	tunnels := repos.GetTunnels()

	utils.JSON(w, http.StatusOK, tunnels)
}

func handleDelete(w http.ResponseWriter, body io.ReadCloser) {
	defer body.Close()
	var data types.KillTunnelDTO

	if err := json.NewDecoder(body).Decode(&data); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if err := validate.Struct(data); err != nil {
		utils.Error(w, http.StatusBadRequest, fmt.Sprintf("Validation failed: %v", err))
		return
	}

	tunnelMu.Lock()
	defer tunnelMu.Unlock()

	if !repos.KillTunnel(data) {
		utils.Error(w, http.StatusConflict, "Tunnel don't exist")
		return
	}

	utils.Success(w, http.StatusOK, "Tunnel deleted successfully")
}
