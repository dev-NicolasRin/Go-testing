package transport

import (
	"API/internal/model"
	"API/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct{
	service *service.Service
}

func New(s *service.Service) *BookHandler {
	return &BookHandler{
		service: s,
	}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodGet:
			libros,err := h.service.ObtenTodosLosLibros()
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} 
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(libros)
		case http.MethodPost:
			var libro model.Libro
			if err := json.NewDecoder(r.Body).Decode(&libro); err != nil{
				http.Error(w,err.Error(), http.StatusBadRequest)
				return
			}

			created,err := h.service.CrearLibro(libro)
			if err != nil{
				http.Error(w,err.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(created)
		default:
			http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}

func (h *BookHandler) HandleBOokByID(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w,"no lo encontre", http.StatusBadRequest)
	}

	switch r.Method{
	case http.MethodGet:
		libro, err := h.service.ObtenLibroPorID(id)
		if err != nil{
			http.Error(w,"no lo encontramos jeje lol xd", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "applications/json")
		json.NewEncoder(w).Encode(libro)

	case http.MethodPut:
		var libro model.Libro
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil{
			http.Error(w, "input invalido", http.StatusBadRequest)
			return
		}
		updated, err := h.service.UpdateAllLibro(id,libro)
		if err != nil{
			http.Error(w,err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)
	case http.MethodDelete:
		err := h.service.BorrarLibro(id)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "metodo no disponible", http.StatusMethodNotAllowed)
	}
	
}