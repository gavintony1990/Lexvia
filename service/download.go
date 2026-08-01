package service

import (
	"encoding/json"
	"net/http"

	"github.com/gavintony1990/Lexvia/service/httputil"
)

// WorkerRequest represents a worker request payload.
// Deprecated: use httputil.WorkerRequest instead.
type WorkerRequest = httputil.WorkerRequest

// DoWorkerRequest sends a request through the worker service.
// Deprecated: use httputil.DoWorkerRequest() instead.
func DoWorkerRequest(req *WorkerRequest) (*http.Response, error) {
	return httputil.DoWorkerRequest(req)
}

// DoDownloadRequest downloads a file from the given URL.
// Deprecated: use httputil.DoDownloadRequest() instead.
func DoDownloadRequest(url string, reason ...string) (*http.Response, error) {
	return httputil.DoDownloadRequest(url, reason...)
}

// Ensure json is imported for backward compatibility
var _ = json.Marshal