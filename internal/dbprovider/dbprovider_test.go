package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSavePoints(t *testing.T) {
	cacheDB := ProvideCacheDB()

	// Test saving points with a valid ID
	id := "receipt1"
	points := 10
	err := cacheDB.SavePoints(id, points)
	require.NoError(t, err)

	// Retrieve points to verify they were saved correctly
	savedPoints, err := cacheDB.GetPointByReceiptID(id)
	require.NoError(t, err)
	assert.Equal(t, points, savedPoints)

	// Test saving points with an empty ID
	err = cacheDB.SavePoints("", points)
	assert.Error(t, err)
	assert.Equal(t, "id cannot be empty", err.Error())
}

func TestGetPointByReceiptID(t *testing.T) {
	cacheDB := ProvideCacheDB()

	// Test retrieval of points with a valid ID
	id := "receipt2"
	points := 20
	cacheDB.SavePoints(id, points)

	retrievedPoints, err := cacheDB.GetPointByReceiptID(id)
	require.NoError(t, err)
	assert.Equal(t, points, retrievedPoints)

	// Test retrieval with a non-existent ID
	nonExistentID := "nonexistent"
	_, err = cacheDB.GetPointByReceiptID(nonExistentID)
	assert.Error(t, err)
	assert.Equal(t, "points not found for the given id", err.Error())

	// Test retrieval with an empty ID
	_, err = cacheDB.GetPointByReceiptID("")
	assert.Error(t, err)
	assert.Equal(t, "id cannot be empty", err.Error())
}
