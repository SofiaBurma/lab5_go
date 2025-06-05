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
	s.router.Get("/students", s.ListStudents)
	s.router.Post("/students", s.CreateStudent)
	s.router.Get("/students/{id}", s.GetStudentByID)
    s.router.Put("/students/{id}", s.UpdateStudent)
    s.router.Delete("/students/{id}", s.DeleteStudent)
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
		http.Error(w, "Groups name is required", http.StatusBadRequest)
		return
	}

	student, err := s.store.CreateStudent(r.Context(), params)
	if err != nil {
		http.Error(w, "Error while creating students", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (s *Server) ListStudents(w http.ResponseWriter, r *http.Request) {
	students, err := s.store.ListStudents(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve students", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]db.User{"users": users})
}

func (s *Server) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	student, err := s.store.GetStudentByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func (s *Server) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var params db.UpdateStudentParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if params.ID <= 0 {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	if params.FullName == "" {
		http.Error(w, "Full name is required", http.StatusBadRequest)
		return
	}

	if params.Age <= 0 {
		http.Error(w, "Age must be positive", http.StatusBadRequest)
		return
	}

	if params.GroupsName == "" {
		http.Error(w, "Groups name is required", http.StatusBadRequest)
		return
	}

	student, err := s.store.UpdateStudent(r.Context(), params)
	if err != nil {
		http.Error(w, "Failed to update student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func (s *Server) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = s.store.DeleteStudent(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Failed to delete student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.router))
}
