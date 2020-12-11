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
	s.router.HandleFunc("/registration", s.saveUser).Methods("POST")

	postsRouter := s.router.PathPrefix("/posts").Subrouter()
	postsRouter.HandleFunc("/", s.findAllPosts).Methods("GET")
	postsRouter.HandleFunc("/{userID}/", s.handleUserPostsByUserID).Methods("GET")

	usersRouter := s.router.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("/", s.findAllUsers).Methods("GET")

	secure := s.router.PathPrefix("/auth").Subrouter()
	secure.Use(s.authenticateUser)
	secure.HandleFunc("/posts/", s.savePost).Methods("POST")
	secure.HandleFunc("/cabinet/", s.handleUserCabinet).Methods("GET")
	secure.HandleFunc("/cabinet/posts/", s.handleAllUserPostsInCabinet).Methods("GET")
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
			validTime, err := time.Parse(time.RFC3339, expiration.(string))
			if !time.Now().Before(validTime) || err != nil {
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

func (s *server) handleUserPostsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	posts, err := s.service.PostService().FindAllPostsByUserID(userID)
	if err != nil {
		s.errorRespond(w, err, 500)
		return
	}
	jsonResponse(w, 200, posts)
}

func (s *server) handleUserCabinet(w http.ResponseWriter, r *http.Request) {
	if user := r.Context().Value(ctxKeyUser).(*model.User); user != nil {
		jsonResponse(w, 200, user)
		return
	}
	s.errorRespond(w, fmt.Errorf("User in contenxt not found"), 401)
}

func (s *server) handleAllUserPostsInCabinet(w http.ResponseWriter, r *http.Request) {
	if user := r.Context().Value(ctxKeyUser).(*model.User); user != nil {
		posts, err := s.service.PostService().FindAllPostsByUserID(user.ID)
		if err != nil {
			s.errorRespond(w, err, 500)
			return
		}
		jsonResponse(w, 200, posts)
		return
	}
	s.errorRespond(w, fmt.Errorf("User in context not found"), 401)
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
	jsonResponse(w, 200, user)
}

func (s *server) savePost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ctxKeyUser).(*model.User)
	saveRequest := &postSaveRequest{}
	json.NewDecoder(r.Body).Decode(saveRequest)
	post := &model.Post{
		Name:    saveRequest.Name,
		Content: saveRequest.Content,
		UserID:  user.ID,
	}
	post.BeforeSaving()
	if err := post.Validate(); err != nil {
		s.errorRespond(w, err, 400)
		return
	}
	s.service.PostService().SavePost(post)
	w.WriteHeader(201)
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
	jsonResponse(w, 200, posts)
}

func (s *server) findAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.service.UserService().FindAllUsers()
	if err != nil {
		s.errorRespond(w, err, 500)
		return
	}
	jsonResponse(w, 200, users)
}

func (s *server) errorRespond(w http.ResponseWriter, err error, status int) {
	s.logger.Debug(err)
	errorMap := map[string]string{
		"Error message": err.Error(),
	}
	jsonResponse(w, status, errorMap)
}

func (s *server) createToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   user.Username,
		"role":       user.Authority,
		"expiration": time.Now().Add(time.Hour * time.Duration(s.config.TokenValidTime)),
	})
	tokenString, err := token.SignedString([]byte(s.config.TokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func jsonResponse(w http.ResponseWriter, status int, body interface{}) {
	jr := &JSONResponse{w, status}
	jr.CreateJSONResponse(status, body)
}
