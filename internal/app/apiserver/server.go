package apiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/service"
	"github.com/vielendanke/restful-service/internal/app/sqlstore"
)

const (
	ctxKeyUser ctxKey = iota + 1
)

type ctxKey int8

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
	s.router.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"*"}),
		handlers.AllowedMethods([]string{"*"}),
		handlers.AllowCredentials(),
	))
	s.router.HandleFunc("/login", s.handleUserLogin)

	postsRouter := s.router.PathPrefix("/posts").Subrouter()
	usersRouter := s.router.PathPrefix("/users").Subrouter()

	postsRouter.HandleFunc("/", s.findAllPosts).Methods("GET")

	usersRouter.HandleFunc("/", s.findAllUsers).Methods("GET")
	usersRouter.HandleFunc("/", s.saveUser).Methods("POST")

	secure := s.router.PathPrefix("/auth").Subrouter()
	secure.Use(s.authenticateUser)
	secure.HandleFunc("/posts/", s.savePost).Methods("POST")
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Token is not valid")
			}
			return []byte(s.config.TokenSecret), nil
		})
		if err != nil {
			s.errorRespond(w, err, 401)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username := claims["username"]
			expiration := claims["expiration"]
			tm := time.Unix(expiration.(int64), 0)
			if !time.Now().Before(tm) {
				s.errorRespond(w, fmt.Errorf("Token is expired"), 401)
				return
			}
			user, err := s.service.UserService().FindByUsername(username.(string))
			if err != nil {
				s.errorRespond(w, err, 401)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, user)))
		} else {
			s.errorRespond(w, fmt.Errorf("Token is not valid"), 401)
			return
		}
	})
}

func (s *server) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	loginRequest := &userLoginRequest{}

	json.NewDecoder(r.Body).Decode(loginRequest)

	user, err := s.service.UserService().Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		s.errorRespond(w, err, 404)
		return
	}
	token, err := s.createToken(user)
	if err != nil {
		s.errorRespond(w, err, 500)
		return
	}
	w.Header().Set("Authorization", token)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
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

func (s *server) createToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   user.Username,
		"role":       user.Authority,
		"expiration": time.Now().Add(time.Hour * 15).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.config.TokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
