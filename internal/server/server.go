package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
	store  db.Store
}

func NewServer(store db.Store) *Server {
	s := &Server{
		router: chi.NewRouter(),
		store:  store,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.Get("/health", s.handleHealth)
	s.router.Post("/students", s.CreateStudent)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func(s *Server) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var params.db.CreateStudentParams
	err:= json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Error while decoding json", http.StatusBadRequest)
		return
	}

	if params.FullName == ""{
		http.Error(w, "Full name is required", http.StatusBadRequest)
		return
	}

	if params.Age <=0 {
		http.Error(w, "Age must be positive", http.StatusBadRequest)
		return
	}

	if params.GroupsName == ""{
		http.Error(w, "Full name is required", http.StatusBadRequest)
		return
	}

	student, err := s.store.CreateStudent(r.Context(), params)
	if err != nil {
		http.Error(w, "Error while creating students", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewDecoder(w).Encode(student)
}



func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.router))
}
