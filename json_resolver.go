package pastry

import (
	"encoding/json"
	"net/http"
)

// JSONResolver resolves JSON responses
type JSONResolver struct{}

// ResolveJSONError returns a default json error response
func (resolver *JSONResolver) ResolveJSONError(w http.ResponseWriter, code int, message string) {
	resolver.ResolveJSON(w, code, map[string]string{"error": message})
}

// ResolveJSON returns a json encoded response
func (resolver *JSONResolver) ResolveJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
