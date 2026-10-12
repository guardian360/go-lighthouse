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

func TestScanTaskAPI_Get_TargetsAndPhases(t *testing.T) {
	var query string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data": {
			"id": "task-1",
			"phases": ["discovery", "vulnerability"],
			"targets": [
				{"id": 1, "scanobject_id": "so-1", "scan_result_id": "sr-1", "host_discovery_id": null, "target": "https://example.test:8443", "template": "ssl-weak-cipher"},
				{"id": 2, "scanobject_id": "so-1", "scan_result_id": null, "host_discovery_id": "hd-1", "target": "10.0.0.5", "template": null}
			]
		}}`))
	}))
	t.Cleanup(server.Close)

	response, err := NewScanTaskAPI(client.New(server.URL), "task-1").With("targets").Get()
	require.NoError(t, err)

	assert.Equal(t, "with=targets", query)
	task := response.Data
	assert.Equal(t, []string{"discovery", "vulnerability"}, task.Phases)
	require.Len(t, task.Targets, 2)

	finding, host := task.Targets[0], task.Targets[1]
	require.NotNil(t, finding.ScanResultID)
	assert.Equal(t, "sr-1", *finding.ScanResultID)
	assert.Nil(t, finding.HostDiscoveryID)
	require.NotNil(t, finding.Template)
	assert.Equal(t, "ssl-weak-cipher", *finding.Template)

	require.NotNil(t, host.HostDiscoveryID)
	assert.Equal(t, "hd-1", *host.HostDiscoveryID)
	assert.Nil(t, host.ScanResultID)
	assert.Nil(t, host.Template, "a host target has no template")
	assert.Equal(t, "10.0.0.5", host.Target)
}

func TestScanTask_NullPhasesMeanTheProfile(t *testing.T) {
	var task ScanTask
	require.NoError(t, json.Unmarshal([]byte(`{"id": "task-1", "phases": null}`), &task))

	assert.Nil(t, task.Phases)
}
