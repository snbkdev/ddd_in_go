package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

type MockCookieStockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCookieStockCheckerMockRecorder
}

type MockCookieStockCheckerMockRecorder struct {
	mock *MockCookieStockChecker
}

func NewMockCookieStockChecker(ctrl *gomock.Controller) *MockCookieStockChecker {
	mock := &MockCookieStockChecker{ctrl: ctrl}
	mock.recorder = &MockCookieStockCheckerMockRecorder{mock}
	return mock
}

func (m *MockCookieStockChecker) EXPECT() *MockCookieStockCheckerMockRecorder {
	return m.recorder
}

func (m *MockCookieStockChecker) AmountInStock(arg0 context.Context) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AmountInStock", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

func (mr *MockCookieStockCheckerMockRecorder) AmountInStock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AmountInStock", reflect.TypeOf((*MockCookieStockChecker)(nil).AmountInStock), arg0)
}

type MockCardCharger struct {
	ctrl     *gomock.Controller
	recorder *MockCardChargerMockRecorder
}

type MockCardChargerMockRecorder struct {
	mock *MockCardCharger
}

func NewMockCardCharger(ctrl *gomock.Controller) *MockCardCharger {
	mock := &MockCardCharger{ctrl: ctrl}
	mock.recorder = &MockCardChargerMockRecorder{mock}
	return mock
}

func (m *MockCardCharger) EXPECT() *MockCardChargerMockRecorder {
	return m.recorder
}

func (m *MockCardCharger) ChargeCard(arg0 context.Context, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChargeCard", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCardChargerMockRecorder) ChargeCard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChargeCard", reflect.TypeOf((*MockCardCharger)(nil).ChargeCard), arg0, arg1, arg2)
}

type MockEmailSender struct {
	ctrl     *gomock.Controller
	recorder *MockEmailSenderMockRecorder
}

type MockEmailSenderMockRecorder struct {
	mock *MockEmailSender
}

func NewMockEmailSender(ctrl *gomock.Controller) *MockEmailSender {
	mock := &MockEmailSender{ctrl: ctrl}
	mock.recorder = &MockEmailSenderMockRecorder{mock}
	return mock
}

func (m *MockEmailSender) EXPECT() *MockEmailSenderMockRecorder {
	return m.recorder
}

func (m *MockEmailSender) SendEmailReceipt(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmailReceipt", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEmailSenderMockRecorder) SendEmailReceipt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmailReceipt", reflect.TypeOf((*MockEmailSender)(nil).SendEmailReceipt), arg0, arg1)
}