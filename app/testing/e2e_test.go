package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"dot/app/handler"
	"dot/app/models/domain"
	"dot/app/repositories"
	"dot/app/services"
	"dot/routes"
)

func setupApp() (*echo.Echo, *gorm.DB) {
	user := "root"
	password := "Buat$and1"
	host := "127.0.0.1"
	port := "3306"
	dbname := "dot_rest_api"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		panic(err)
	}
	userRepo := repositories.NewAuthRepository(db)
	authService := services.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	e := echo.New()
	routes.SetupRoutes(e, authHandler)

	return e, db
}

func TestLoginFlow(t *testing.T) {
	e, _ := setupApp()
	loginPayload := map[string]string{
		"email":    "cobas@gmail.com",
		"password": "password",
	}
	loginBody, _ := json.Marshal(loginPayload)

	req := httptest.NewRequest(http.MethodPost, "/testing/login", bytes.NewReader(loginBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var loginResp struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}
	err := json.Unmarshal(rec.Body.Bytes(), &loginResp)
	assert.NoError(t, err)
	assert.Equal(t, "Login successful", loginResp.Message)
	assert.NotEmpty(t, loginResp.Token)
}
