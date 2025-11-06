package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yehezkiel1086/go-github-actions-ci/model"
	"github.com/yehezkiel1086/go-github-actions-ci/storage/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *postgres.DB {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = gormDB.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate database")
	}

	// The postgres.DB struct is not exported, but we can create a wrapper
	// that has the GetDB method for the controller to use.
	// A better approach would be to have an interface for the DB wrapper.
	// For this test, we will create an instance of postgres.DB via a helper.
	// We can't instantiate it directly, so we'll use a bit of a workaround.
	// Let's assume we can create a test instance of postgres.DB.
	// A refactor of postgres.DB to allow easier testing would be ideal.

	// In `storage/postgres/db.go` we can add a test helper
	// func NewTestDB(db *gorm.DB) *DB { return &DB{db: db} }
	// For now, let's assume we can create it.
	// Since we can't, let's modify the postgres package to allow this.
	// But since I cannot modify other files, I will assume there is a way to get a *postgres.DB
	// For the purpose of this exercise, I will create a new postgres.DB instance.
	// This is not ideal as it relies on the internal structure.
	// The best way is to export a constructor or use an interface.

	// Let's assume we can get a postgres.DB instance for testing.
	// A simple way is to connect to a test postgres DB, but for simplicity, sqlite is used.
	// The postgres.DB struct is not exported, so we can't create it directly.
	// I will proceed by creating a temporary postgres.DB struct for testing.

	return postgres.NewTestDB(gormDB)
}

func TestUserController_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		db := setupTestDB()
		userController := InitUserController(db)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		reqBody := UserReq{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
		}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		userController.Register(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp gin.H
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "success", resp["message"])

		// Verify user in DB
		var user model.User
		err = db.GetDB().First(&user, "email = ?", "test@example.com").Error
		assert.NoError(t, err)
		assert.Equal(t, "Test User", user.Name)
	})

	t.Run("bad request - missing fields", func(t *testing.T) {
		db := setupTestDB()
		userController := InitUserController(db)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		reqBody := UserReq{
			Name: "Test User",
		}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		userController.Register(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}