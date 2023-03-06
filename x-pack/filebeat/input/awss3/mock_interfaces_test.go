// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package awss3 is a generated GoMock package.
package awss3

import (
	context "context"
	reflect "reflect"
	time "time"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	types "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	gomock "github.com/golang/mock/gomock"

	beat "github.com/elastic/beats/v7/libbeat/beat"
	aws "github.com/elastic/beats/v7/x-pack/libbeat/common/aws"
	logp "github.com/elastic/elastic-agent-libs/logp"
)

// MockSQSAPI is a mock of sqsAPI interface.
type MockSQSAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSQSAPIMockRecorder
}

// MockSQSAPIMockRecorder is the mock recorder for MockSQSAPI.
type MockSQSAPIMockRecorder struct {
	mock *MockSQSAPI
}

// NewMockSQSAPI creates a new mock instance.
func NewMockSQSAPI(ctrl *gomock.Controller) *MockSQSAPI {
	mock := &MockSQSAPI{ctrl: ctrl}
	mock.recorder = &MockSQSAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQSAPI) EXPECT() *MockSQSAPIMockRecorder {
	return m.recorder
}

// ChangeMessageVisibility mocks base method.
func (m *MockSQSAPI) ChangeMessageVisibility(ctx context.Context, msg *types.Message, timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeMessageVisibility", ctx, msg, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeMessageVisibility indicates an expected call of ChangeMessageVisibility.
func (mr *MockSQSAPIMockRecorder) ChangeMessageVisibility(ctx, msg, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMessageVisibility", reflect.TypeOf((*MockSQSAPI)(nil).ChangeMessageVisibility), ctx, msg, timeout)
}

// DeleteMessage mocks base method.
func (m *MockSQSAPI) DeleteMessage(ctx context.Context, msg *types.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockSQSAPIMockRecorder) DeleteMessage(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockSQSAPI)(nil).DeleteMessage), ctx, msg)
}

// GetQueueAttributes mocks base method.
func (m *MockSQSAPI) GetQueueAttributes(ctx context.Context, attr []types.QueueAttributeName) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueueAttributes", ctx, attr)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueueAttributes indicates an expected call of GetQueueAttributes.
func (mr *MockSQSAPIMockRecorder) GetQueueAttributes(ctx, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueAttributes", reflect.TypeOf((*MockSQSAPI)(nil).GetQueueAttributes), ctx, attr)
}

// ReceiveMessage mocks base method.
func (m *MockSQSAPI) ReceiveMessage(ctx context.Context, maxMessages int) ([]types.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage", ctx, maxMessages)
	ret0, _ := ret[0].([]types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MockSQSAPIMockRecorder) ReceiveMessage(ctx, maxMessages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockSQSAPI)(nil).ReceiveMessage), ctx, maxMessages)
}

// MocksqsReceiver is a mock of sqsReceiver interface.
type MocksqsReceiver struct {
	ctrl     *gomock.Controller
	recorder *MocksqsReceiverMockRecorder
}

// MocksqsReceiverMockRecorder is the mock recorder for MocksqsReceiver.
type MocksqsReceiverMockRecorder struct {
	mock *MocksqsReceiver
}

// NewMocksqsReceiver creates a new mock instance.
func NewMocksqsReceiver(ctrl *gomock.Controller) *MocksqsReceiver {
	mock := &MocksqsReceiver{ctrl: ctrl}
	mock.recorder = &MocksqsReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksqsReceiver) EXPECT() *MocksqsReceiverMockRecorder {
	return m.recorder
}

// ReceiveMessage mocks base method.
func (m *MocksqsReceiver) ReceiveMessage(ctx context.Context, maxMessages int) ([]types.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage", ctx, maxMessages)
	ret0, _ := ret[0].([]types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MocksqsReceiverMockRecorder) ReceiveMessage(ctx, maxMessages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MocksqsReceiver)(nil).ReceiveMessage), ctx, maxMessages)
}

// MocksqsDeleter is a mock of sqsDeleter interface.
type MocksqsDeleter struct {
	ctrl     *gomock.Controller
	recorder *MocksqsDeleterMockRecorder
}

// MocksqsDeleterMockRecorder is the mock recorder for MocksqsDeleter.
type MocksqsDeleterMockRecorder struct {
	mock *MocksqsDeleter
}

// NewMocksqsDeleter creates a new mock instance.
func NewMocksqsDeleter(ctrl *gomock.Controller) *MocksqsDeleter {
	mock := &MocksqsDeleter{ctrl: ctrl}
	mock.recorder = &MocksqsDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksqsDeleter) EXPECT() *MocksqsDeleterMockRecorder {
	return m.recorder
}

// DeleteMessage mocks base method.
func (m *MocksqsDeleter) DeleteMessage(ctx context.Context, msg *types.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MocksqsDeleterMockRecorder) DeleteMessage(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MocksqsDeleter)(nil).DeleteMessage), ctx, msg)
}

// MocksqsVisibilityChanger is a mock of sqsVisibilityChanger interface.
type MocksqsVisibilityChanger struct {
	ctrl     *gomock.Controller
	recorder *MocksqsVisibilityChangerMockRecorder
}

// MocksqsVisibilityChangerMockRecorder is the mock recorder for MocksqsVisibilityChanger.
type MocksqsVisibilityChangerMockRecorder struct {
	mock *MocksqsVisibilityChanger
}

// NewMocksqsVisibilityChanger creates a new mock instance.
func NewMocksqsVisibilityChanger(ctrl *gomock.Controller) *MocksqsVisibilityChanger {
	mock := &MocksqsVisibilityChanger{ctrl: ctrl}
	mock.recorder = &MocksqsVisibilityChangerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksqsVisibilityChanger) EXPECT() *MocksqsVisibilityChangerMockRecorder {
	return m.recorder
}

// ChangeMessageVisibility mocks base method.
func (m *MocksqsVisibilityChanger) ChangeMessageVisibility(ctx context.Context, msg *types.Message, timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeMessageVisibility", ctx, msg, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeMessageVisibility indicates an expected call of ChangeMessageVisibility.
func (mr *MocksqsVisibilityChangerMockRecorder) ChangeMessageVisibility(ctx, msg, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMessageVisibility", reflect.TypeOf((*MocksqsVisibilityChanger)(nil).ChangeMessageVisibility), ctx, msg, timeout)
}

// MocksqsAttributeGetter is a mock of sqsAttributeGetter interface.
type MocksqsAttributeGetter struct {
	ctrl     *gomock.Controller
	recorder *MocksqsAttributeGetterMockRecorder
}

// MocksqsAttributeGetterMockRecorder is the mock recorder for MocksqsAttributeGetter.
type MocksqsAttributeGetterMockRecorder struct {
	mock *MocksqsAttributeGetter
}

// NewMocksqsAttributeGetter creates a new mock instance.
func NewMocksqsAttributeGetter(ctrl *gomock.Controller) *MocksqsAttributeGetter {
	mock := &MocksqsAttributeGetter{ctrl: ctrl}
	mock.recorder = &MocksqsAttributeGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksqsAttributeGetter) EXPECT() *MocksqsAttributeGetterMockRecorder {
	return m.recorder
}

// GetQueueAttributes mocks base method.
func (m *MocksqsAttributeGetter) GetQueueAttributes(ctx context.Context, attr []types.QueueAttributeName) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueueAttributes", ctx, attr)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueueAttributes indicates an expected call of GetQueueAttributes.
func (mr *MocksqsAttributeGetterMockRecorder) GetQueueAttributes(ctx, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueAttributes", reflect.TypeOf((*MocksqsAttributeGetter)(nil).GetQueueAttributes), ctx, attr)
}

// MockSQSProcessor is a mock of sqsProcessor interface.
type MockSQSProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockSQSProcessorMockRecorder
}

// MockSQSProcessorMockRecorder is the mock recorder for MockSQSProcessor.
type MockSQSProcessorMockRecorder struct {
	mock *MockSQSProcessor
}

// NewMockSQSProcessor creates a new mock instance.
func NewMockSQSProcessor(ctrl *gomock.Controller) *MockSQSProcessor {
	mock := &MockSQSProcessor{ctrl: ctrl}
	mock.recorder = &MockSQSProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQSProcessor) EXPECT() *MockSQSProcessorMockRecorder {
	return m.recorder
}

// ProcessSQS mocks base method.
func (m *MockSQSProcessor) ProcessSQS(ctx context.Context, msg *types.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessSQS", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessSQS indicates an expected call of ProcessSQS.
func (mr *MockSQSProcessorMockRecorder) ProcessSQS(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSQS", reflect.TypeOf((*MockSQSProcessor)(nil).ProcessSQS), ctx, msg)
}

// MockS3API is a mock of s3API interface.
type MockS3API struct {
	ctrl     *gomock.Controller
	recorder *MockS3APIMockRecorder
}

// MockS3APIMockRecorder is the mock recorder for MockS3API.
type MockS3APIMockRecorder struct {
	mock *MockS3API
}

// NewMockS3API creates a new mock instance.
func NewMockS3API(ctrl *gomock.Controller) *MockS3API {
	mock := &MockS3API{ctrl: ctrl}
	mock.recorder = &MockS3APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3API) EXPECT() *MockS3APIMockRecorder {
	return m.recorder
}

// CopyObject mocks base method.
func (m *MockS3API) CopyObject(ctx context.Context, from_bucket, to_bucket, from_key, to_key string) (*s3.CopyObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyObject", ctx, from_bucket, to_bucket, from_key, to_key)
	ret0, _ := ret[0].(*s3.CopyObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CopyObject indicates an expected call of CopyObject.
func (mr *MockS3APIMockRecorder) CopyObject(ctx, from_bucket, to_bucket, from_key, to_key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyObject", reflect.TypeOf((*MockS3API)(nil).CopyObject), ctx, from_bucket, to_bucket, from_key, to_key)
}

// DeleteObject mocks base method.
func (m *MockS3API) DeleteObject(ctx context.Context, bucket, key string) (*s3.DeleteObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", ctx, bucket, key)
	ret0, _ := ret[0].(*s3.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockS3APIMockRecorder) DeleteObject(ctx, bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockS3API)(nil).DeleteObject), ctx, bucket, key)
}

// GetObject mocks base method.
func (m *MockS3API) GetObject(ctx context.Context, bucket, key string) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", ctx, bucket, key)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockS3APIMockRecorder) GetObject(ctx, bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3API)(nil).GetObject), ctx, bucket, key)
}

// ListObjectsPaginator mocks base method.
func (m *MockS3API) ListObjectsPaginator(bucket, prefix string) s3Pager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectsPaginator", bucket, prefix)
	ret0, _ := ret[0].(s3Pager)
	return ret0
}

// ListObjectsPaginator indicates an expected call of ListObjectsPaginator.
func (mr *MockS3APIMockRecorder) ListObjectsPaginator(bucket, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsPaginator", reflect.TypeOf((*MockS3API)(nil).ListObjectsPaginator), bucket, prefix)
}

// Mocks3Getter is a mock of s3Getter interface.
type Mocks3Getter struct {
	ctrl     *gomock.Controller
	recorder *Mocks3GetterMockRecorder
}

// Mocks3GetterMockRecorder is the mock recorder for Mocks3Getter.
type Mocks3GetterMockRecorder struct {
	mock *Mocks3Getter
}

// NewMocks3Getter creates a new mock instance.
func NewMocks3Getter(ctrl *gomock.Controller) *Mocks3Getter {
	mock := &Mocks3Getter{ctrl: ctrl}
	mock.recorder = &Mocks3GetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocks3Getter) EXPECT() *Mocks3GetterMockRecorder {
	return m.recorder
}

// GetObject mocks base method.
func (m *Mocks3Getter) GetObject(ctx context.Context, bucket, key string) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", ctx, bucket, key)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *Mocks3GetterMockRecorder) GetObject(ctx, bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*Mocks3Getter)(nil).GetObject), ctx, bucket, key)
}

// Mocks3Mover is a mock of s3Mover interface.
type Mocks3Mover struct {
	ctrl     *gomock.Controller
	recorder *Mocks3MoverMockRecorder
}

// Mocks3MoverMockRecorder is the mock recorder for Mocks3Mover.
type Mocks3MoverMockRecorder struct {
	mock *Mocks3Mover
}

// NewMocks3Mover creates a new mock instance.
func NewMocks3Mover(ctrl *gomock.Controller) *Mocks3Mover {
	mock := &Mocks3Mover{ctrl: ctrl}
	mock.recorder = &Mocks3MoverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocks3Mover) EXPECT() *Mocks3MoverMockRecorder {
	return m.recorder
}

// CopyObject mocks base method.
func (m *Mocks3Mover) CopyObject(ctx context.Context, from_bucket, to_bucket, from_key, to_key string) (*s3.CopyObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyObject", ctx, from_bucket, to_bucket, from_key, to_key)
	ret0, _ := ret[0].(*s3.CopyObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CopyObject indicates an expected call of CopyObject.
func (mr *Mocks3MoverMockRecorder) CopyObject(ctx, from_bucket, to_bucket, from_key, to_key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyObject", reflect.TypeOf((*Mocks3Mover)(nil).CopyObject), ctx, from_bucket, to_bucket, from_key, to_key)
}

// DeleteObject mocks base method.
func (m *Mocks3Mover) DeleteObject(ctx context.Context, bucket, key string) (*s3.DeleteObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", ctx, bucket, key)
	ret0, _ := ret[0].(*s3.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *Mocks3MoverMockRecorder) DeleteObject(ctx, bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*Mocks3Mover)(nil).DeleteObject), ctx, bucket, key)
}

// Mocks3Lister is a mock of s3Lister interface.
type Mocks3Lister struct {
	ctrl     *gomock.Controller
	recorder *Mocks3ListerMockRecorder
}

// Mocks3ListerMockRecorder is the mock recorder for Mocks3Lister.
type Mocks3ListerMockRecorder struct {
	mock *Mocks3Lister
}

// NewMocks3Lister creates a new mock instance.
func NewMocks3Lister(ctrl *gomock.Controller) *Mocks3Lister {
	mock := &Mocks3Lister{ctrl: ctrl}
	mock.recorder = &Mocks3ListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocks3Lister) EXPECT() *Mocks3ListerMockRecorder {
	return m.recorder
}

// ListObjectsPaginator mocks base method.
func (m *Mocks3Lister) ListObjectsPaginator(bucket, prefix string) s3Pager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectsPaginator", bucket, prefix)
	ret0, _ := ret[0].(s3Pager)
	return ret0
}

// ListObjectsPaginator indicates an expected call of ListObjectsPaginator.
func (mr *Mocks3ListerMockRecorder) ListObjectsPaginator(bucket, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsPaginator", reflect.TypeOf((*Mocks3Lister)(nil).ListObjectsPaginator), bucket, prefix)
}

// MockS3Pager is a mock of s3Pager interface.
type MockS3Pager struct {
	ctrl     *gomock.Controller
	recorder *MockS3PagerMockRecorder
}

// MockS3PagerMockRecorder is the mock recorder for MockS3Pager.
type MockS3PagerMockRecorder struct {
	mock *MockS3Pager
}

// NewMockS3Pager creates a new mock instance.
func NewMockS3Pager(ctrl *gomock.Controller) *MockS3Pager {
	mock := &MockS3Pager{ctrl: ctrl}
	mock.recorder = &MockS3PagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3Pager) EXPECT() *MockS3PagerMockRecorder {
	return m.recorder
}

// HasMorePages mocks base method.
func (m *MockS3Pager) HasMorePages() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasMorePages")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasMorePages indicates an expected call of HasMorePages.
func (mr *MockS3PagerMockRecorder) HasMorePages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasMorePages", reflect.TypeOf((*MockS3Pager)(nil).HasMorePages))
}

// NextPage mocks base method.
func (m *MockS3Pager) NextPage(ctx context.Context, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NextPage", varargs...)
	ret0, _ := ret[0].(*s3.ListObjectsV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextPage indicates an expected call of NextPage.
func (mr *MockS3PagerMockRecorder) NextPage(ctx interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextPage", reflect.TypeOf((*MockS3Pager)(nil).NextPage), varargs...)
}

// MockS3ObjectHandlerFactory is a mock of s3ObjectHandlerFactory interface.
type MockS3ObjectHandlerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockS3ObjectHandlerFactoryMockRecorder
}

// MockS3ObjectHandlerFactoryMockRecorder is the mock recorder for MockS3ObjectHandlerFactory.
type MockS3ObjectHandlerFactoryMockRecorder struct {
	mock *MockS3ObjectHandlerFactory
}

// NewMockS3ObjectHandlerFactory creates a new mock instance.
func NewMockS3ObjectHandlerFactory(ctrl *gomock.Controller) *MockS3ObjectHandlerFactory {
	mock := &MockS3ObjectHandlerFactory{ctrl: ctrl}
	mock.recorder = &MockS3ObjectHandlerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3ObjectHandlerFactory) EXPECT() *MockS3ObjectHandlerFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockS3ObjectHandlerFactory) Create(ctx context.Context, log *logp.Logger, client beat.Client, acker *aws.EventACKTracker, obj s3EventV2) s3ObjectHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, log, client, acker, obj)
	ret0, _ := ret[0].(s3ObjectHandler)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockS3ObjectHandlerFactoryMockRecorder) Create(ctx, log, client, acker, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockS3ObjectHandlerFactory)(nil).Create), ctx, log, client, acker, obj)
}

// MockS3ObjectHandler is a mock of s3ObjectHandler interface.
type MockS3ObjectHandler struct {
	ctrl     *gomock.Controller
	recorder *MockS3ObjectHandlerMockRecorder
}

// MockS3ObjectHandlerMockRecorder is the mock recorder for MockS3ObjectHandler.
type MockS3ObjectHandlerMockRecorder struct {
	mock *MockS3ObjectHandler
}

// NewMockS3ObjectHandler creates a new mock instance.
func NewMockS3ObjectHandler(ctrl *gomock.Controller) *MockS3ObjectHandler {
	mock := &MockS3ObjectHandler{ctrl: ctrl}
	mock.recorder = &MockS3ObjectHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3ObjectHandler) EXPECT() *MockS3ObjectHandlerMockRecorder {
	return m.recorder
}

// FinalizeS3Object mocks base method.
func (m *MockS3ObjectHandler) FinalizeS3Object() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeS3Object")
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeS3Object indicates an expected call of FinalizeS3Object.
func (mr *MockS3ObjectHandlerMockRecorder) FinalizeS3Object() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeS3Object", reflect.TypeOf((*MockS3ObjectHandler)(nil).FinalizeS3Object))
}

// ProcessS3Object mocks base method.
func (m *MockS3ObjectHandler) ProcessS3Object() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessS3Object")
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessS3Object indicates an expected call of ProcessS3Object.
func (mr *MockS3ObjectHandlerMockRecorder) ProcessS3Object() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessS3Object", reflect.TypeOf((*MockS3ObjectHandler)(nil).ProcessS3Object))
}

// Wait mocks base method.
func (m *MockS3ObjectHandler) Wait() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Wait")
}

// Wait indicates an expected call of Wait.
func (mr *MockS3ObjectHandlerMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockS3ObjectHandler)(nil).Wait))
}
