package apiserver_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vielendanke/restful-service/internal/app/model"
)

func TestSever_HandleUserCreate(t *testing.T) {
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"username": "example@example.org",
				"password": "password",
				"nickname": "nickname",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid request",
			payload:      "invalid body",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/registration", b)
			testServer.ServeHTTP(rec, req)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleFindAllUsers(t *testing.T) {
	u := model.MockUser()
	token, _ := testServer.CreateToken(u)

	testCases := []struct {
		name         string
		payload      string
		expectedCode int
	}{
		{
			name:         "valid token",
			payload:      token,
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid token",
			payload:      "invalidtoken",
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/auth/users/", nil)
			req.Header.Set("Authorization", tc.payload)
			testServer.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleFindAllPosts(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts/", nil)
	testServer.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestServer_HandleFindAllPostsByUserID(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/posts/%s/", uuid.New().String()), nil)
	testServer.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestServer_HandleUserLogin(t *testing.T) {
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid request",
			payload: map[string]string{
				"username": "username@mail.ru",
				"password": "userpassword",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid request",
			payload: "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/login", b)
			testServer.ServeHTTP(rec, req)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandlePostSave(t *testing.T) {
	validPayload := map[string]string{
		"name":    "Valid post name",
		"content": "Valid post content about post name",
	}
	u := model.MockUser()

	token, _ := testServer.CreateToken(u)

	testCases := []struct {
		name         string
		payload      interface{}
		token        string
		expectedCode int
	}{
		{
			name:         "valid post and token",
			payload:      validPayload,
			token:        token,
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid post",
			payload:      "invalid",
			token:        token,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid token",
			payload:      validPayload,
			token:        "invalid",
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/auth/posts/", b)
			req.Header.Set("Authorization", tc.token)
			testServer.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.NoError(t, err)
		})
	}
}

func TestServer_HandleUserCabinet(t *testing.T) {
	u := model.MockUser()

	token, _ := testServer.CreateToken(u)

	testCases := []struct {
		name         string
		payload      string
		expectedCode int
	}{
		{
			name:         "valid token",
			payload:      token,
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid token",
			payload:      "invalid",
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/auth/cabinet/", nil)
			req.Header.Set("Authorization", tc.payload)
			testServer.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleUsersPostInCabinet(t *testing.T) {
	u := model.MockUser()

	token, _ := testServer.CreateToken(u)

	testCases := []struct {
		name         string
		payload      string
		expectedCode int
	}{
		{
			name:         "valid token",
			payload:      token,
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid token",
			payload:      "invalid",
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/auth/cabinet/posts/", nil)
			req.Header.Set("Authorization", tc.payload)
			testServer.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
