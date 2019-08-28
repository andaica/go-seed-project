package home

import (
	"encoding/json"
	"net/http"

	log "../../logger"
	"../../middleware"
)

const message = "Hello seed!"

func RegistRouter(mux *http.ServeMux) {
	// NOT use middleware
	// mux.HandleFunc("/", log.TimeLogging(HomePage))
	// DO use middleware, code is much more complex:
	//
	mux.HandleFunc("/", middleware.NewHandler(HomePage).Use(log.TimeLogging).Get())
	mux.HandleFunc("/pages", registHandle(ListPage).Get())
	mux.HandleFunc("/pages/add", registHandle(AddPage).Post())
	mux.HandleFunc("/pages/update", registHandle(UpdatePage).Post())
	mux.HandleFunc("/pages/delete", registHandle(DeletePage).Post())
}

func registHandle(h http.HandlerFunc) (handler middleware.Handler) {
	handler = middleware.NewHandler(h).Use(log.TimeLogging)
	return
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func ListPage(w http.ResponseWriter, r *http.Request) {
	listPages, err := allPages()
	if err != nil {
		response := map[string]interface{}{"status": "NG", "message": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listPages)
}

func AddPage(w http.ResponseWriter, r *http.Request) {
	page, err := getDataFromPost(r)
	if err != nil {
		response := map[string]interface{}{"status": "NG", "message": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	page, err = insertPage(page)
	if err != nil {
		response := map[string]interface{}{"status": "NG", "message": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"status": "OK", "page": page}
	json.NewEncoder(w).Encode(response)
}

func UpdatePage(w http.ResponseWriter, r *http.Request) {
	page, err := getDataFromPost(r)
	if err != nil {
		response := map[string]interface{}{"status": "NG", "message": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	page, err = updatePage(page)
	if err != nil {
		response := map[string]interface{}{"status": "NG", "message": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"status": "OK", "page": page}
	json.NewEncoder(w).Encode(response)
}

func DeletePage(w http.ResponseWriter, r *http.Request) {
	page, err := getDataFromPost(r)
	if err != nil {
		response := map[string]interface{}{"status": "NG", "message": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = deletePage(page)
	if err != nil {
		response := map[string]interface{}{"status": "NG", "message": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"status": "OK", "result": "Delete success"}
	json.NewEncoder(w).Encode(response)
}

func getDataFromPost(r *http.Request) (page PageModel, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&page)
	if err != nil {
		return
	}
	return
}
