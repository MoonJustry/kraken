// Code generated by MockGen. DO NOT EDIT.
// Source: code.uber.internal/infra/kraken/origin/blobclient (interfaces: Client,Provider,ClusterClient,ClusterProvider,ClientResolver)

// Package mockblobclient is a generated GoMock package.
package mockblobclient

import (
	core "code.uber.internal/infra/kraken/core"
	blobclient "code.uber.internal/infra/kraken/origin/blobclient"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
	time "time"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Addr mocks base method
func (m *MockClient) Addr() string {
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(string)
	return ret0
}

// Addr indicates an expected call of Addr
func (mr *MockClientMockRecorder) Addr() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockClient)(nil).Addr))
}

// DeleteBlob mocks base method
func (m *MockClient) DeleteBlob(arg0 core.Digest) error {
	ret := m.ctrl.Call(m, "DeleteBlob", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBlob indicates an expected call of DeleteBlob
func (mr *MockClientMockRecorder) DeleteBlob(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBlob", reflect.TypeOf((*MockClient)(nil).DeleteBlob), arg0)
}

// DownloadBlob mocks base method
func (m *MockClient) DownloadBlob(arg0 string, arg1 core.Digest, arg2 io.Writer) error {
	ret := m.ctrl.Call(m, "DownloadBlob", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadBlob indicates an expected call of DownloadBlob
func (mr *MockClientMockRecorder) DownloadBlob(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadBlob", reflect.TypeOf((*MockClient)(nil).DownloadBlob), arg0, arg1, arg2)
}

// DuplicateUploadBlob mocks base method
func (m *MockClient) DuplicateUploadBlob(arg0 string, arg1 core.Digest, arg2 io.Reader, arg3 time.Duration) error {
	ret := m.ctrl.Call(m, "DuplicateUploadBlob", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DuplicateUploadBlob indicates an expected call of DuplicateUploadBlob
func (mr *MockClientMockRecorder) DuplicateUploadBlob(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DuplicateUploadBlob", reflect.TypeOf((*MockClient)(nil).DuplicateUploadBlob), arg0, arg1, arg2, arg3)
}

// ForceCleanup mocks base method
func (m *MockClient) ForceCleanup(arg0 time.Duration) error {
	ret := m.ctrl.Call(m, "ForceCleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceCleanup indicates an expected call of ForceCleanup
func (mr *MockClientMockRecorder) ForceCleanup(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceCleanup", reflect.TypeOf((*MockClient)(nil).ForceCleanup), arg0)
}

// GetMetaInfo mocks base method
func (m *MockClient) GetMetaInfo(arg0 string, arg1 core.Digest) (*core.MetaInfo, error) {
	ret := m.ctrl.Call(m, "GetMetaInfo", arg0, arg1)
	ret0, _ := ret[0].(*core.MetaInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaInfo indicates an expected call of GetMetaInfo
func (mr *MockClientMockRecorder) GetMetaInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaInfo", reflect.TypeOf((*MockClient)(nil).GetMetaInfo), arg0, arg1)
}

// GetPeerContext mocks base method
func (m *MockClient) GetPeerContext() (core.PeerContext, error) {
	ret := m.ctrl.Call(m, "GetPeerContext")
	ret0, _ := ret[0].(core.PeerContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeerContext indicates an expected call of GetPeerContext
func (mr *MockClientMockRecorder) GetPeerContext() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerContext", reflect.TypeOf((*MockClient)(nil).GetPeerContext))
}

// Locations mocks base method
func (m *MockClient) Locations(arg0 core.Digest) ([]string, error) {
	ret := m.ctrl.Call(m, "Locations", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Locations indicates an expected call of Locations
func (mr *MockClientMockRecorder) Locations(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Locations", reflect.TypeOf((*MockClient)(nil).Locations), arg0)
}

// OverwriteMetaInfo mocks base method
func (m *MockClient) OverwriteMetaInfo(arg0 core.Digest, arg1 int64) error {
	ret := m.ctrl.Call(m, "OverwriteMetaInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OverwriteMetaInfo indicates an expected call of OverwriteMetaInfo
func (mr *MockClientMockRecorder) OverwriteMetaInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OverwriteMetaInfo", reflect.TypeOf((*MockClient)(nil).OverwriteMetaInfo), arg0, arg1)
}

// ReplicateToRemote mocks base method
func (m *MockClient) ReplicateToRemote(arg0 string, arg1 core.Digest, arg2 string) error {
	ret := m.ctrl.Call(m, "ReplicateToRemote", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplicateToRemote indicates an expected call of ReplicateToRemote
func (mr *MockClientMockRecorder) ReplicateToRemote(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateToRemote", reflect.TypeOf((*MockClient)(nil).ReplicateToRemote), arg0, arg1, arg2)
}

// Stat mocks base method
func (m *MockClient) Stat(arg0 string, arg1 core.Digest) (*core.BlobInfo, error) {
	ret := m.ctrl.Call(m, "Stat", arg0, arg1)
	ret0, _ := ret[0].(*core.BlobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat
func (mr *MockClientMockRecorder) Stat(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockClient)(nil).Stat), arg0, arg1)
}

// StatLocal mocks base method
func (m *MockClient) StatLocal(arg0 string, arg1 core.Digest) (*core.BlobInfo, error) {
	ret := m.ctrl.Call(m, "StatLocal", arg0, arg1)
	ret0, _ := ret[0].(*core.BlobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatLocal indicates an expected call of StatLocal
func (mr *MockClientMockRecorder) StatLocal(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatLocal", reflect.TypeOf((*MockClient)(nil).StatLocal), arg0, arg1)
}

// TransferBlob mocks base method
func (m *MockClient) TransferBlob(arg0 core.Digest, arg1 io.Reader) error {
	ret := m.ctrl.Call(m, "TransferBlob", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferBlob indicates an expected call of TransferBlob
func (mr *MockClientMockRecorder) TransferBlob(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferBlob", reflect.TypeOf((*MockClient)(nil).TransferBlob), arg0, arg1)
}

// UploadBlob mocks base method
func (m *MockClient) UploadBlob(arg0 string, arg1 core.Digest, arg2 io.Reader) error {
	ret := m.ctrl.Call(m, "UploadBlob", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadBlob indicates an expected call of UploadBlob
func (mr *MockClientMockRecorder) UploadBlob(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadBlob", reflect.TypeOf((*MockClient)(nil).UploadBlob), arg0, arg1, arg2)
}

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// Provide mocks base method
func (m *MockProvider) Provide(arg0 string) blobclient.Client {
	ret := m.ctrl.Call(m, "Provide", arg0)
	ret0, _ := ret[0].(blobclient.Client)
	return ret0
}

// Provide indicates an expected call of Provide
func (mr *MockProviderMockRecorder) Provide(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockProvider)(nil).Provide), arg0)
}

// MockClusterClient is a mock of ClusterClient interface
type MockClusterClient struct {
	ctrl     *gomock.Controller
	recorder *MockClusterClientMockRecorder
}

// MockClusterClientMockRecorder is the mock recorder for MockClusterClient
type MockClusterClientMockRecorder struct {
	mock *MockClusterClient
}

// NewMockClusterClient creates a new mock instance
func NewMockClusterClient(ctrl *gomock.Controller) *MockClusterClient {
	mock := &MockClusterClient{ctrl: ctrl}
	mock.recorder = &MockClusterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterClient) EXPECT() *MockClusterClientMockRecorder {
	return m.recorder
}

// DownloadBlob mocks base method
func (m *MockClusterClient) DownloadBlob(arg0 string, arg1 core.Digest, arg2 io.Writer) error {
	ret := m.ctrl.Call(m, "DownloadBlob", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadBlob indicates an expected call of DownloadBlob
func (mr *MockClusterClientMockRecorder) DownloadBlob(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadBlob", reflect.TypeOf((*MockClusterClient)(nil).DownloadBlob), arg0, arg1, arg2)
}

// GetMetaInfo mocks base method
func (m *MockClusterClient) GetMetaInfo(arg0 string, arg1 core.Digest) (*core.MetaInfo, error) {
	ret := m.ctrl.Call(m, "GetMetaInfo", arg0, arg1)
	ret0, _ := ret[0].(*core.MetaInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaInfo indicates an expected call of GetMetaInfo
func (mr *MockClusterClientMockRecorder) GetMetaInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaInfo", reflect.TypeOf((*MockClusterClient)(nil).GetMetaInfo), arg0, arg1)
}

// OverwriteMetaInfo mocks base method
func (m *MockClusterClient) OverwriteMetaInfo(arg0 core.Digest, arg1 int64) error {
	ret := m.ctrl.Call(m, "OverwriteMetaInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OverwriteMetaInfo indicates an expected call of OverwriteMetaInfo
func (mr *MockClusterClientMockRecorder) OverwriteMetaInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OverwriteMetaInfo", reflect.TypeOf((*MockClusterClient)(nil).OverwriteMetaInfo), arg0, arg1)
}

// Owners mocks base method
func (m *MockClusterClient) Owners(arg0 core.Digest) ([]core.PeerContext, error) {
	ret := m.ctrl.Call(m, "Owners", arg0)
	ret0, _ := ret[0].([]core.PeerContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Owners indicates an expected call of Owners
func (mr *MockClusterClientMockRecorder) Owners(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owners", reflect.TypeOf((*MockClusterClient)(nil).Owners), arg0)
}

// ReplicateToRemote mocks base method
func (m *MockClusterClient) ReplicateToRemote(arg0 string, arg1 core.Digest, arg2 string) error {
	ret := m.ctrl.Call(m, "ReplicateToRemote", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplicateToRemote indicates an expected call of ReplicateToRemote
func (mr *MockClusterClientMockRecorder) ReplicateToRemote(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateToRemote", reflect.TypeOf((*MockClusterClient)(nil).ReplicateToRemote), arg0, arg1, arg2)
}

// Stat mocks base method
func (m *MockClusterClient) Stat(arg0 string, arg1 core.Digest) (*core.BlobInfo, error) {
	ret := m.ctrl.Call(m, "Stat", arg0, arg1)
	ret0, _ := ret[0].(*core.BlobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat
func (mr *MockClusterClientMockRecorder) Stat(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockClusterClient)(nil).Stat), arg0, arg1)
}

// UploadBlob mocks base method
func (m *MockClusterClient) UploadBlob(arg0 string, arg1 core.Digest, arg2 io.Reader) error {
	ret := m.ctrl.Call(m, "UploadBlob", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadBlob indicates an expected call of UploadBlob
func (mr *MockClusterClientMockRecorder) UploadBlob(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadBlob", reflect.TypeOf((*MockClusterClient)(nil).UploadBlob), arg0, arg1, arg2)
}

// MockClusterProvider is a mock of ClusterProvider interface
type MockClusterProvider struct {
	ctrl     *gomock.Controller
	recorder *MockClusterProviderMockRecorder
}

// MockClusterProviderMockRecorder is the mock recorder for MockClusterProvider
type MockClusterProviderMockRecorder struct {
	mock *MockClusterProvider
}

// NewMockClusterProvider creates a new mock instance
func NewMockClusterProvider(ctrl *gomock.Controller) *MockClusterProvider {
	mock := &MockClusterProvider{ctrl: ctrl}
	mock.recorder = &MockClusterProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterProvider) EXPECT() *MockClusterProviderMockRecorder {
	return m.recorder
}

// Provide mocks base method
func (m *MockClusterProvider) Provide(arg0 string) (blobclient.ClusterClient, error) {
	ret := m.ctrl.Call(m, "Provide", arg0)
	ret0, _ := ret[0].(blobclient.ClusterClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provide indicates an expected call of Provide
func (mr *MockClusterProviderMockRecorder) Provide(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockClusterProvider)(nil).Provide), arg0)
}

// MockClientResolver is a mock of ClientResolver interface
type MockClientResolver struct {
	ctrl     *gomock.Controller
	recorder *MockClientResolverMockRecorder
}

// MockClientResolverMockRecorder is the mock recorder for MockClientResolver
type MockClientResolverMockRecorder struct {
	mock *MockClientResolver
}

// NewMockClientResolver creates a new mock instance
func NewMockClientResolver(ctrl *gomock.Controller) *MockClientResolver {
	mock := &MockClientResolver{ctrl: ctrl}
	mock.recorder = &MockClientResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientResolver) EXPECT() *MockClientResolverMockRecorder {
	return m.recorder
}

// Resolve mocks base method
func (m *MockClientResolver) Resolve(arg0 core.Digest) ([]blobclient.Client, error) {
	ret := m.ctrl.Call(m, "Resolve", arg0)
	ret0, _ := ret[0].([]blobclient.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (mr *MockClientResolverMockRecorder) Resolve(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockClientResolver)(nil).Resolve), arg0)
}
