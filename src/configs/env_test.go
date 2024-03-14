package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvMongoURI(t *testing.T) {

	// use table method to check both the error test and non error test.

	actualURL := EnvMongoURI()
	expectedURL := "mongodb://localhost:27017"
	// if assert.NoErrorf(t, err, "Unexpected error encountered:  %s", err) {
	// 	assert.Equal(t, expectedURL, actualURL)
	// }
	assert.Equal(t, expectedURL, actualURL)

}

func TestEnvDatabase(t *testing.T) {

	actualDB := EnvDatabase()
	expectedDB := "EmployeeManagement"

	assert.Equal(t, expectedDB, actualDB)
}
func TestEnvCollection(t *testing.T) {

	actualCollection := EnvCollection()
	expectedCollection := "EmployeeDetails"

	assert.Equal(t, expectedCollection, actualCollection)

}
