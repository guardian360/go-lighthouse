package v2

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// stopRequest is what a stub Lighthouse received when a scan task was stopped.
type stopRequest struct {
	method      string
	path        string
	contentType string
	body        []byte
}

// stopScanTask stops scan task "task-1" against a stub Lighthouse and returns
// what the stub received.
func stopScanTask(t *testing.T, stop func(*ScanTaskAPI) (*ScanTaskAPIResponse, error)) stopRequest {
	t.Helper()

	var received stopRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		received = stopRequest{method: r.Method, path: r.URL.Path, contentType: r.Header.Get("Content-Type"), body: body}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data": {}}`))
	}))
	t.Cleanup(server.Close)

	_, err := stop(NewScanTaskAPI(client.New(server.URL), "task-1"))
	require.NoError(t, err)

	return received
}

func TestScanTaskAPI_Stop(t *testing.T) {
	t.Run("posts to the task's stop route with no body", func(t *testing.T) {
		received := stopScanTask(t, (*ScanTaskAPI).Stop)

		assert.Equal(t, http.MethodPost, received.method)
		assert.Equal(t, "/api/v2/scan-tasks/task-1/stop", received.path)
		assert.Empty(t, received.body, "Stop sends no body, as it always has")
	})

	t.Run("a nil report is the same request as Stop", func(t *testing.T) {
		stopped := stopScanTask(t, (*ScanTaskAPI).Stop)
		reported := stopScanTask(t, func(s *ScanTaskAPI) (*ScanTaskAPIResponse, error) { return s.StopAndReport(nil) })

		assert.Equal(t, stopped, reported)
	})
}

func TestScanTaskAPI_StopAndReport(t *testing.T) {
	t.Run("sends the report as the JSON body of the stop request", func(t *testing.T) {
		report := api.APIRequestPayload{"coverage": map[string]any{"protocols": []string{"tcp"}}}

		received := stopScanTask(t, func(s *ScanTaskAPI) (*ScanTaskAPIResponse, error) { return s.StopAndReport(report) })

		assert.Equal(t, http.MethodPost, received.method)
		assert.Equal(t, "/api/v2/scan-tasks/task-1/stop", received.path)
		assert.Equal(t, "application/json", received.contentType)
		var sent map[string]any
		require.NoError(t, json.Unmarshal(received.body, &sent))
		assert.Equal(t, map[string]any{"coverage": map[string]any{"protocols": []any{"tcp"}}}, sent)
	})
}
