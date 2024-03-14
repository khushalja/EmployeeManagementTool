package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"EmployeeManagementTool/src/model"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) error {
	args := m.Called(ctx, filter)
	return args.Error(0)
}
func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}
func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}
func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}
func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

type MockContext struct {
	mock.Mock
}

// WithTimeout mocks the WithTimeout method of the context.Context interface
func (m *MockContext) WithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	args := m.Called(ctx, timeout)
	return args.Get(0).(context.Context), args.Get(1).(context.CancelFunc)
}

type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) Connect(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(*mongo.Client), args.Error(1)
}

type MockCursor struct {
	mock.Mock
}

func (m *MockCursor) All(ctx context.Context, document interface{}) error {
	args := m.Called(ctx, document)
	return args.Error(0)
}

func Getrouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestLoginManager(t *testing.T) {

	mockCol := new(MockCollection)
	mockClient := new(MockMongoClient)
	ctx := context.TODO()
	mockContext := new(MockContext)

	mockContext.On("WithTimeout", ctx, mock.AnythingOfType("time.Duration"))
	resultContext, resultCancelFunc := mockContext.WithTimeout(ctx, 30*time.Second)

	assert.NotNil(t, resultContext)
	assert.NotNil(t, resultCancelFunc)

	mockCol.On("FindOne", ctx, bson.M{"employeeid": "testID"}).Return(nil)
	result := mockCol.FindOne(ctx, bson.M{"employeeid": "testID"})
	assert.Nil(t, result)

	mockClient.On("Connect", ctx, mock.Anything).Return(&mongo.Client{}, nil)

	client, err := mockClient.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	assert.NoError(t, err)
	assert.NotNil(t, client)

	r := Getrouter()
	r.POST("/employee/login", LoginManager())
	emp := &model.Employee{
		EmployeeId: 1212121,
		Password:   "letsRock",
	}
	empJson, _ := json.Marshal(emp)

	req, _ := http.NewRequest("POST", "/employee/login", bytes.NewBuffer(empJson))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockClient.AssertExpectations(t)
	mockCol.AssertExpectations(t)

}

func TestGetAllEmployees(t *testing.T) {
	mockCol := new(MockCollection)
	mockClient := new(MockMongoClient)
	mockCursor := new(MockCursor)
	ctx := context.TODO()
	mockContext := new(MockContext)

	mockContext.On("WithTimeout", ctx, mock.AnythingOfType("time.Duration"))
	resultContext, resultCancelFunc := mockContext.WithTimeout(ctx, 30*time.Second)
	assert.NotNil(t, resultContext)
	assert.NotNil(t, resultCancelFunc)

	mockCol.On("Find", ctx, bson.M{}, mock.Anything).Return(&mongo.Cursor{}, nil)
	result, err := mockCol.Find(ctx, bson.M{})
	assert.NotNil(t, result)
	assert.NoError(t, err)

	mockClient.On("Connect", ctx, mock.Anything).Return(&mongo.Client{}, nil)
	client, err := mockClient.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.NoError(t, err)
	assert.NotNil(t, client)

	mockCol.On("Find", ctx, mock.Anything).Return(&mongo.Cursor{}, nil)
	findResult, err := mockCol.Find(ctx, mock.Anything)
	assert.NoError(t, err)
	assert.NotNil(t, findResult)

	mockCursor.On("All", ctx, mock.Anything).Return(nil)
	allErr := mockCursor.All(ctx, mock.Anything)
	assert.NoError(t, allErr)

	r := Getrouter()
	r.GET("/employee/employeedetails", GetAllEmployees())
	req, _ := http.NewRequest("GET", "/employee/employeedetails", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetEmployee(t *testing.T) {

}
