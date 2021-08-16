package mockit

import (
	"github.com/stretchr/testify/mock"
)

type MockIt struct {
	mock.Mock
}

func (m *MockIt) Given(methodName string) *MockMethod {
	method := &MockMethod{
		MethodName: methodName,
		mock:       &m.Mock,
	}

	return method
}
