package schtasks

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestRunEveryMinutes also tests Get and ForceDelete
func TestRunEveryMinutes(t *testing.T) {
	// Create the task
	taskName := t.Name()
	_, err := RunEveryMinutes(taskName, 1, "dir")
	require.NoError(t, err)

	// Get task to verify it was created
	st, err := Get(taskName)
	require.NoError(t, err)
	require.Equal(t, taskName, st.TaskName)

	// Delete the task
	_, err = ForceDelete(taskName)
	require.NoError(t, err)

	// Get task to verify it was deleted
	st, err = Get(taskName)
	require.Error(t, err)
	require.Contains(t, err.Error(), fmt.Sprintf("task %s not found", taskName))
}

func TestTimeAtMinutes(t *testing.T) {
	t1Str := "2006-01-02T15:04:50Z"
	t2Str := "2006-01-02T15:04:55Z"
	t3Str := "2006-01-02T23:59:56Z"

	t1, err := time.Parse(time.RFC3339, t1Str)
	require.NoError(t, err)
	t2, err := time.Parse(time.RFC3339, t2Str)
	require.NoError(t, err)
	t3, err := time.Parse(time.RFC3339, t3Str)
	require.NoError(t, err)

	timeAt := TimeAtMinutes(t1, 1)
	require.Equal(t, "15:05", timeAt)
	timeAt = TimeAtMinutes(t1, 2)
	require.Equal(t, "15:06", timeAt)

	// Assuming PaddingSeconds = 5
	timeAt = TimeAtMinutes(t2, 1)
	require.Equal(t, "15:06", timeAt)
	timeAt = TimeAtMinutes(t2, 2)
	require.Equal(t, "15:07", timeAt)

	timeAt = TimeAtMinutes(t3, 1)
	require.Equal(t, "00:01", timeAt)
	timeAt = TimeAtMinutes(t3, 2)
	require.Equal(t, "00:02", timeAt)
}

func TestRunAtMinutes(t *testing.T) {
	// Create the task
	taskName := t.Name()
	_, err := RunAtMinutes(taskName, 1, "dir")
	require.NoError(t, err)

	// Get task to verify it was created
	st, err := Get(taskName)
	require.NoError(t, err)
	require.Equal(t, taskName, st.TaskName)

	// Delete the task
	_, err = ForceDelete(taskName)
	require.NoError(t, err)

	// Get task to verify it was deleted
	st, err = Get(taskName)
	require.Error(t, err)
	require.Contains(t, err.Error(), fmt.Sprintf("task %s not found", taskName))
}

