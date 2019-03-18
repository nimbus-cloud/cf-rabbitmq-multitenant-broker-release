// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"rabbitmq-service-broker/rabbithutch"
	"sync"
)

type FakeRabbitHutch struct {
	CreateUserStub        func(string, string, string) (string, error)
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createUserReturns struct {
		result1 string
		result2 error
	}
	createUserReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	EnsureVHostExistsStub        func(string) error
	ensureVHostExistsMutex       sync.RWMutex
	ensureVHostExistsArgsForCall []struct {
		arg1 string
	}
	ensureVHostExistsReturns struct {
		result1 error
	}
	ensureVHostExistsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRabbitHutch) CreateUser(arg1 string, arg2 string, arg3 string) (string, error) {
	fake.createUserMutex.Lock()
	ret, specificReturn := fake.createUserReturnsOnCall[len(fake.createUserArgsForCall)]
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateUser", []interface{}{arg1, arg2, arg3})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRabbitHutch) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeRabbitHutch) CreateUserCalls(stub func(string, string, string) (string, error)) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = stub
}

func (fake *FakeRabbitHutch) CreateUserArgsForCall(i int) (string, string, string) {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	argsForCall := fake.createUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRabbitHutch) CreateUserReturns(result1 string, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRabbitHutch) CreateUserReturnsOnCall(i int, result1 string, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	if fake.createUserReturnsOnCall == nil {
		fake.createUserReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createUserReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRabbitHutch) EnsureVHostExists(arg1 string) error {
	fake.ensureVHostExistsMutex.Lock()
	ret, specificReturn := fake.ensureVHostExistsReturnsOnCall[len(fake.ensureVHostExistsArgsForCall)]
	fake.ensureVHostExistsArgsForCall = append(fake.ensureVHostExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("EnsureVHostExists", []interface{}{arg1})
	fake.ensureVHostExistsMutex.Unlock()
	if fake.EnsureVHostExistsStub != nil {
		return fake.EnsureVHostExistsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.ensureVHostExistsReturns
	return fakeReturns.result1
}

func (fake *FakeRabbitHutch) EnsureVHostExistsCallCount() int {
	fake.ensureVHostExistsMutex.RLock()
	defer fake.ensureVHostExistsMutex.RUnlock()
	return len(fake.ensureVHostExistsArgsForCall)
}

func (fake *FakeRabbitHutch) EnsureVHostExistsCalls(stub func(string) error) {
	fake.ensureVHostExistsMutex.Lock()
	defer fake.ensureVHostExistsMutex.Unlock()
	fake.EnsureVHostExistsStub = stub
}

func (fake *FakeRabbitHutch) EnsureVHostExistsArgsForCall(i int) string {
	fake.ensureVHostExistsMutex.RLock()
	defer fake.ensureVHostExistsMutex.RUnlock()
	argsForCall := fake.ensureVHostExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRabbitHutch) EnsureVHostExistsReturns(result1 error) {
	fake.ensureVHostExistsMutex.Lock()
	defer fake.ensureVHostExistsMutex.Unlock()
	fake.EnsureVHostExistsStub = nil
	fake.ensureVHostExistsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRabbitHutch) EnsureVHostExistsReturnsOnCall(i int, result1 error) {
	fake.ensureVHostExistsMutex.Lock()
	defer fake.ensureVHostExistsMutex.Unlock()
	fake.EnsureVHostExistsStub = nil
	if fake.ensureVHostExistsReturnsOnCall == nil {
		fake.ensureVHostExistsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.ensureVHostExistsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRabbitHutch) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.ensureVHostExistsMutex.RLock()
	defer fake.ensureVHostExistsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRabbitHutch) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ rabbithutch.RabbitHutch = new(FakeRabbitHutch)
