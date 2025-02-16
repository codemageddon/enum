package status

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	_ "modernc.org/sqlite"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {

	t.Run("basic", func(t *testing.T) {
		s := StatusActive
		assert.Equal(t, "active", s.String())
	})

	t.Run("json", func(t *testing.T) {
		type Data struct {
			Status Status `json:"status"`
		}

		d := Data{Status: StatusActive}
		b, err := json.Marshal(d)
		require.NoError(t, err)
		assert.Equal(t, `{"status":"active"}`, string(b))

		var d2 Data
		err = json.Unmarshal([]byte(`{"status":"inactive"}`), &d2)
		require.NoError(t, err)
		assert.Equal(t, StatusInactive, d2.Status)
	})

	t.Run("sql", func(t *testing.T) {
		s := StatusActive

		// test Value() method
		v, err := s.Value()
		require.NoError(t, err)
		assert.Equal(t, "active", v)

		// test Scan from string
		var s2 Status
		err = s2.Scan("inactive")
		require.NoError(t, err)
		assert.Equal(t, StatusInactive, s2)

		// test Scan from []byte
		err = s2.Scan([]byte("blocked"))
		require.NoError(t, err)
		assert.Equal(t, StatusBlocked, s2)

		// test Scan from nil - should get first value from StatusValues()
		err = s2.Scan(nil)
		require.NoError(t, err)
		assert.Equal(t, StatusValues()[0], s2)

		// test invalid value
		err = s2.Scan(123)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid status value")
	})

	t.Run("sqlite", func(t *testing.T) {
		db, err := sql.Open("sqlite", ":memory:")
		require.NoError(t, err)
		defer db.Close()

		// create table with status column
		_, err = db.Exec(`CREATE TABLE test_status (id INTEGER PRIMARY KEY, status TEXT)`)
		require.NoError(t, err)

		// insert different status values
		statuses := []Status{StatusActive, StatusInactive, StatusBlocked}
		for i, s := range statuses {
			_, err = db.Exec(`INSERT INTO test_status (id, status) VALUES (?, ?)`, i+1, s)
			require.NoError(t, err)
		}

		// insert nil status
		_, err = db.Exec(`INSERT INTO test_status (id, status) VALUES (?, ?)`, 4, nil)
		require.NoError(t, err)

		// read and verify each status
		for i, expected := range statuses {
			var s Status
			err = db.QueryRow(`SELECT status FROM test_status WHERE id = ?`, i+1).Scan(&s)
			require.NoError(t, err)
			assert.Equal(t, expected, s)
		}

		// verify nil status gets first value
		var s Status
		err = db.QueryRow(`SELECT status FROM test_status WHERE id = 4`).Scan(&s)
		require.NoError(t, err)
		assert.Equal(t, StatusValues()[0], s)
	})

	t.Run("invalid", func(t *testing.T) {
		var d struct {
			Status Status `json:"status"`
		}
		err := json.Unmarshal([]byte(`{"status":"invalid"}`), &d)
		assert.Error(t, err)
	})
}

func ExampleStatus() {
	s := StatusActive
	fmt.Println(s.String())
	// Output: active
}
