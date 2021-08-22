package gohttpmock

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"sync"
)

var (
	mockUpServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	mocks map[string]*Mock
}

func StartMockServer() {
	mockUpServer.serverMutex.Lock()
	defer mockUpServer.serverMutex.Unlock()

	mockUpServer.enabled = true
}

func StopMockServer() {
	mockUpServer.serverMutex.Lock()
	defer mockUpServer.serverMutex.Unlock()

	mockUpServer.enabled = false
}

func DeleteMocks() {
	mockUpServer.serverMutex.Lock()
	defer mockUpServer.serverMutex.Unlock()

	mockUpServer.mocks = make(map[string]*Mock)
}

func AddMock(mock Mock) {
	mockUpServer.serverMutex.Lock()
	defer mockUpServer.serverMutex.Unlock()

	key := mockUpServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockUpServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)

	if body == "" {
		return ""
	}

	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")

	return body
}

func GetMock(method, url, body string) *Mock {
	if !mockUpServer.enabled {
		return nil
	}
	if mock := mockUpServer.mocks[mockUpServer.getMockKey(method, url, body)]; mock != nil {
		return mock
	}

	mock := Mock{
		Error: errors.New("no mock matching"),
	}

	return &mock
}
