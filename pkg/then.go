package mockit

import "github.com/stretchr/testify/mock"

type MockResult struct {
	Call *mock.Call
}

func (result *MockResult) Then(returnValues ...interface{}) *MockResult {
	result.Call.Return(returnValues)
	return result
}

func (result *MockResult) Times(times int) *MockResult {
	result.Call.Times(times)
	return result
}
