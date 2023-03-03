// api.go

package mymodule

import (
	"encoding/json"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// APIResponse represents the JSON response from the external API
type APIResponse struct {
	Data string `json:"data"`
}

// GetExternalData sends an HTTP GET request to the external API
// and returns the response
func GetExternalData() (APIResponse, error) {
	resp, err := http.Get("https://api.openweathermap.org/data/3.0/onecall?lat=33.44&lon=-94.04&exclude=hourly,daily&appid=5a11f355013746e7127f8e963a9aad52")
	if err != nil {
		return APIResponse{}, err
	}
	defer resp.Body.Close()

	var data APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return APIResponse{}, err
	}

	return data, nil
}

// HandleExternalDataRequest is a handler function for the module's external data endpoint
func (m Module) HandleExternalDataRequest(ctx sdk.Context, w http.ResponseWriter, r *http.Request) {
	// Make the API request to retrieve external data
	data, err := GetExternalData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response to the client
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
