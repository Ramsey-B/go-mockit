package mockit

import "github.com/stretchr/testify/mock"

type MockMethod struct {
	MethodName string
	mock       *mock.Mock
}

func (method *MockMethod) When(arguments ...interface{}) *MockResult {
	return &MockResult{
		call: method.mock.On(method.MethodName, arguments...),
	}
}
