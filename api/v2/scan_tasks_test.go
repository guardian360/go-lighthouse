package v2

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// recordingHTTPClient keeps the last request it was asked to send and answers
// with an empty scan task.
type recordingHTTPClient struct {
	request *http.Request
	body    []byte
}

func (r *recordingHTTPClient) Do(req *http.Request) (*http.Response, error) {
	r.request = req
	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		r.body = body
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Body:       io.NopCloser(bytes.NewBufferString(`{"data": {}}`)),
	}, nil
}

func newRecordingScanTaskAPI() (*ScanTaskAPI, *recordingHTTPClient) {
	recorder := &recordingHTTPClient{}
	c := &client.Client{BaseURL: "https://lighthouse.example", Client: recorder}
	return NewScanTaskAPI(c, "task-1"), recorder
}

func TestScanTaskAPI_Stop(t *testing.T) {
	t.Run("posts to the task's stop route with no body", func(t *testing.T) {
		scanTask, recorder := newRecordingScanTaskAPI()

		_, err := scanTask.Stop()

		require.NoError(t, err)
		assert.Equal(t, http.MethodPost, recorder.request.Method)
		assert.Equal(t, "https://lighthouse.example/api/v2/scan-tasks/task-1/stop", recorder.request.URL.String())
		assert.Empty(t, recorder.body, "Stop sends no body, as it always has")
	})

	t.Run("StopWith sends its payload as the JSON body", func(t *testing.T) {
		scanTask, recorder := newRecordingScanTaskAPI()

		_, err := scanTask.StopWith(api.APIRequestPayload{"coverage": map[string]any{"protocols": []string{"tcp"}}})

		require.NoError(t, err)
		assert.Equal(t, http.MethodPost, recorder.request.Method)
		assert.Equal(t, "https://lighthouse.example/api/v2/scan-tasks/task-1/stop", recorder.request.URL.String())
		var sent map[string]any
		require.NoError(t, json.Unmarshal(recorder.body, &sent))
		assert.Equal(t, map[string]any{"coverage": map[string]any{"protocols": []any{"tcp"}}}, sent)
	})

	t.Run("StopWith with no payload behaves exactly like Stop", func(t *testing.T) {
		scanTask, recorder := newRecordingScanTaskAPI()

		_, err := scanTask.StopWith(nil)

		require.NoError(t, err)
		assert.Empty(t, recorder.body)
	})
}
