package mapper

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	connect()
	info, err := engine.DBMetas()
	assert.NoError(t, err)
	assert.NotNil(t, info)
}
