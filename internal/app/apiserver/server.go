package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/service"
	"github.com/vielendanke/restful-service/internal/app/sqlstore"
)

type server struct {
	logger  *logrus.Logger
	router  *mux.Router
	config  *Config
	service *service.Service
}

func newServer(store *sqlstore.Store, config *Config) (*server, error) {
	server := &server{
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		config:  config,
		service: service.NewService(store),
	}
	if err := server.configureLogger(); err != nil {
		return nil, err
	}
	server.configureRouter()
	return server, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/login", s.handleUserLogin)

	postsRouter := s.router.PathPrefix("/posts").Subrouter()
	usersRouter := s.router.PathPrefix("/users").Subrouter()

	postsRouter.HandleFunc("/", s.findAllPosts).Methods("GET")
	postsRouter.HandleFunc("/", s.savePost).Methods("POST")

	usersRouter.HandleFunc("/", s.findAllUsers).Methods("GET")
	usersRouter.HandleFunc("/", s.saveUser).Methods("POST")
}

func (s *server) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	loginRequest := &userLoginRequest{}

	json.NewDecoder(r.Body).Decode(loginRequest)

	user, err := s.service.UserService().Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		s.errorRespond(w, err, 404)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (s *server) savePost(w http.ResponseWriter, r *http.Request) {

}

func (s *server) saveUser(w http.ResponseWriter, r *http.Request) {
	userRequest := &userSaveRequest{}

	json.NewDecoder(r.Body).Decode(userRequest)

	user := &model.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
		Nickname: userRequest.Nickname,
	}
	err := s.service.UserService().SaveUser(user)
	if err != nil {
		s.errorRespond(w, err, 500)
	}
	w.WriteHeader(201)
}

func (s *server) findAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := s.service.PostService().FindAllPosts()
	if err != nil {
		s.errorRespond(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func (s *server) findAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.service.UserService().FindAllUsers()
	if err != nil {
		s.errorRespond(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (s *server) errorRespond(w http.ResponseWriter, err error, status int) {
	s.logger.Debug(err)
	errorMap := map[string]string{
		"Error message": err.Error(),
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorMap)
}
