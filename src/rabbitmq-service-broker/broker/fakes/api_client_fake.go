// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"rabbitmq-service-broker/broker"
	"sync"

	rabbithole "github.com/michaelklishin/rabbit-hole"
)

type FakeAPIClient struct {
	DeleteUserStub        func(string) (*http.Response, error)
	deleteUserMutex       sync.RWMutex
	deleteUserArgsForCall []struct {
		arg1 string
	}
	deleteUserReturns struct {
		result1 *http.Response
		result2 error
	}
	deleteUserReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	DeleteVhostStub        func(string) (*http.Response, error)
	deleteVhostMutex       sync.RWMutex
	deleteVhostArgsForCall []struct {
		arg1 string
	}
	deleteVhostReturns struct {
		result1 *http.Response
		result2 error
	}
	deleteVhostReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	GetVhostStub        func(string) (*rabbithole.VhostInfo, error)
	getVhostMutex       sync.RWMutex
	getVhostArgsForCall []struct {
		arg1 string
	}
	getVhostReturns struct {
		result1 *rabbithole.VhostInfo
		result2 error
	}
	getVhostReturnsOnCall map[int]struct {
		result1 *rabbithole.VhostInfo
		result2 error
	}
	ListUsersStub        func() ([]rabbithole.UserInfo, error)
	listUsersMutex       sync.RWMutex
	listUsersArgsForCall []struct {
	}
	listUsersReturns struct {
		result1 []rabbithole.UserInfo
		result2 error
	}
	listUsersReturnsOnCall map[int]struct {
		result1 []rabbithole.UserInfo
		result2 error
	}
	ProtocolPortsStub        func() (map[string]rabbithole.Port, error)
	protocolPortsMutex       sync.RWMutex
	protocolPortsArgsForCall []struct {
	}
	protocolPortsReturns struct {
		result1 map[string]rabbithole.Port
		result2 error
	}
	protocolPortsReturnsOnCall map[int]struct {
		result1 map[string]rabbithole.Port
		result2 error
	}
	PutPolicyStub        func(string, string, rabbithole.Policy) (*http.Response, error)
	putPolicyMutex       sync.RWMutex
	putPolicyArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 rabbithole.Policy
	}
	putPolicyReturns struct {
		result1 *http.Response
		result2 error
	}
	putPolicyReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	PutUserStub        func(string, rabbithole.UserSettings) (*http.Response, error)
	putUserMutex       sync.RWMutex
	putUserArgsForCall []struct {
		arg1 string
		arg2 rabbithole.UserSettings
	}
	putUserReturns struct {
		result1 *http.Response
		result2 error
	}
	putUserReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	PutVhostStub        func(string, rabbithole.VhostSettings) (*http.Response, error)
	putVhostMutex       sync.RWMutex
	putVhostArgsForCall []struct {
		arg1 string
		arg2 rabbithole.VhostSettings
	}
	putVhostReturns struct {
		result1 *http.Response
		result2 error
	}
	putVhostReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	UpdatePermissionsInStub        func(string, string, rabbithole.Permissions) (*http.Response, error)
	updatePermissionsInMutex       sync.RWMutex
	updatePermissionsInArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 rabbithole.Permissions
	}
	updatePermissionsInReturns struct {
		result1 *http.Response
		result2 error
	}
	updatePermissionsInReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPIClient) DeleteUser(arg1 string) (*http.Response, error) {
	fake.deleteUserMutex.Lock()
	ret, specificReturn := fake.deleteUserReturnsOnCall[len(fake.deleteUserArgsForCall)]
	fake.deleteUserArgsForCall = append(fake.deleteUserArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteUser", []interface{}{arg1})
	fake.deleteUserMutex.Unlock()
	if fake.DeleteUserStub != nil {
		return fake.DeleteUserStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIClient) DeleteUserCallCount() int {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return len(fake.deleteUserArgsForCall)
}

func (fake *FakeAPIClient) DeleteUserCalls(stub func(string) (*http.Response, error)) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = stub
}

func (fake *FakeAPIClient) DeleteUserArgsForCall(i int) string {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	argsForCall := fake.deleteUserArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAPIClient) DeleteUserReturns(result1 *http.Response, result2 error) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = nil
	fake.deleteUserReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) DeleteUserReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = nil
	if fake.deleteUserReturnsOnCall == nil {
		fake.deleteUserReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.deleteUserReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) DeleteVhost(arg1 string) (*http.Response, error) {
	fake.deleteVhostMutex.Lock()
	ret, specificReturn := fake.deleteVhostReturnsOnCall[len(fake.deleteVhostArgsForCall)]
	fake.deleteVhostArgsForCall = append(fake.deleteVhostArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteVhost", []interface{}{arg1})
	fake.deleteVhostMutex.Unlock()
	if fake.DeleteVhostStub != nil {
		return fake.DeleteVhostStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteVhostReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIClient) DeleteVhostCallCount() int {
	fake.deleteVhostMutex.RLock()
	defer fake.deleteVhostMutex.RUnlock()
	return len(fake.deleteVhostArgsForCall)
}

func (fake *FakeAPIClient) DeleteVhostCalls(stub func(string) (*http.Response, error)) {
	fake.deleteVhostMutex.Lock()
	defer fake.deleteVhostMutex.Unlock()
	fake.DeleteVhostStub = stub
}

func (fake *FakeAPIClient) DeleteVhostArgsForCall(i int) string {
	fake.deleteVhostMutex.RLock()
	defer fake.deleteVhostMutex.RUnlock()
	argsForCall := fake.deleteVhostArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAPIClient) DeleteVhostReturns(result1 *http.Response, result2 error) {
	fake.deleteVhostMutex.Lock()
	defer fake.deleteVhostMutex.Unlock()
	fake.DeleteVhostStub = nil
	fake.deleteVhostReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) DeleteVhostReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.deleteVhostMutex.Lock()
	defer fake.deleteVhostMutex.Unlock()
	fake.DeleteVhostStub = nil
	if fake.deleteVhostReturnsOnCall == nil {
		fake.deleteVhostReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.deleteVhostReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) GetVhost(arg1 string) (*rabbithole.VhostInfo, error) {
	fake.getVhostMutex.Lock()
	ret, specificReturn := fake.getVhostReturnsOnCall[len(fake.getVhostArgsForCall)]
	fake.getVhostArgsForCall = append(fake.getVhostArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetVhost", []interface{}{arg1})
	fake.getVhostMutex.Unlock()
	if fake.GetVhostStub != nil {
		return fake.GetVhostStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getVhostReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIClient) GetVhostCallCount() int {
	fake.getVhostMutex.RLock()
	defer fake.getVhostMutex.RUnlock()
	return len(fake.getVhostArgsForCall)
}

func (fake *FakeAPIClient) GetVhostCalls(stub func(string) (*rabbithole.VhostInfo, error)) {
	fake.getVhostMutex.Lock()
	defer fake.getVhostMutex.Unlock()
	fake.GetVhostStub = stub
}

func (fake *FakeAPIClient) GetVhostArgsForCall(i int) string {
	fake.getVhostMutex.RLock()
	defer fake.getVhostMutex.RUnlock()
	argsForCall := fake.getVhostArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAPIClient) GetVhostReturns(result1 *rabbithole.VhostInfo, result2 error) {
	fake.getVhostMutex.Lock()
	defer fake.getVhostMutex.Unlock()
	fake.GetVhostStub = nil
	fake.getVhostReturns = struct {
		result1 *rabbithole.VhostInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) GetVhostReturnsOnCall(i int, result1 *rabbithole.VhostInfo, result2 error) {
	fake.getVhostMutex.Lock()
	defer fake.getVhostMutex.Unlock()
	fake.GetVhostStub = nil
	if fake.getVhostReturnsOnCall == nil {
		fake.getVhostReturnsOnCall = make(map[int]struct {
			result1 *rabbithole.VhostInfo
			result2 error
		})
	}
	fake.getVhostReturnsOnCall[i] = struct {
		result1 *rabbithole.VhostInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) ListUsers() ([]rabbithole.UserInfo, error) {
	fake.listUsersMutex.Lock()
	ret, specificReturn := fake.listUsersReturnsOnCall[len(fake.listUsersArgsForCall)]
	fake.listUsersArgsForCall = append(fake.listUsersArgsForCall, struct {
	}{})
	fake.recordInvocation("ListUsers", []interface{}{})
	fake.listUsersMutex.Unlock()
	if fake.ListUsersStub != nil {
		return fake.ListUsersStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listUsersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIClient) ListUsersCallCount() int {
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	return len(fake.listUsersArgsForCall)
}

func (fake *FakeAPIClient) ListUsersCalls(stub func() ([]rabbithole.UserInfo, error)) {
	fake.listUsersMutex.Lock()
	defer fake.listUsersMutex.Unlock()
	fake.ListUsersStub = stub
}

func (fake *FakeAPIClient) ListUsersReturns(result1 []rabbithole.UserInfo, result2 error) {
	fake.listUsersMutex.Lock()
	defer fake.listUsersMutex.Unlock()
	fake.ListUsersStub = nil
	fake.listUsersReturns = struct {
		result1 []rabbithole.UserInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) ListUsersReturnsOnCall(i int, result1 []rabbithole.UserInfo, result2 error) {
	fake.listUsersMutex.Lock()
	defer fake.listUsersMutex.Unlock()
	fake.ListUsersStub = nil
	if fake.listUsersReturnsOnCall == nil {
		fake.listUsersReturnsOnCall = make(map[int]struct {
			result1 []rabbithole.UserInfo
			result2 error
		})
	}
	fake.listUsersReturnsOnCall[i] = struct {
		result1 []rabbithole.UserInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) ProtocolPorts() (map[string]rabbithole.Port, error) {
	fake.protocolPortsMutex.Lock()
	ret, specificReturn := fake.protocolPortsReturnsOnCall[len(fake.protocolPortsArgsForCall)]
	fake.protocolPortsArgsForCall = append(fake.protocolPortsArgsForCall, struct {
	}{})
	fake.recordInvocation("ProtocolPorts", []interface{}{})
	fake.protocolPortsMutex.Unlock()
	if fake.ProtocolPortsStub != nil {
		return fake.ProtocolPortsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.protocolPortsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIClient) ProtocolPortsCallCount() int {
	fake.protocolPortsMutex.RLock()
	defer fake.protocolPortsMutex.RUnlock()
	return len(fake.protocolPortsArgsForCall)
}

func (fake *FakeAPIClient) ProtocolPortsCalls(stub func() (map[string]rabbithole.Port, error)) {
	fake.protocolPortsMutex.Lock()
	defer fake.protocolPortsMutex.Unlock()
	fake.ProtocolPortsStub = stub
}

func (fake *FakeAPIClient) ProtocolPortsReturns(result1 map[string]rabbithole.Port, result2 error) {
	fake.protocolPortsMutex.Lock()
	defer fake.protocolPortsMutex.Unlock()
	fake.ProtocolPortsStub = nil
	fake.protocolPortsReturns = struct {
		result1 map[string]rabbithole.Port
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) ProtocolPortsReturnsOnCall(i int, result1 map[string]rabbithole.Port, result2 error) {
	fake.protocolPortsMutex.Lock()
	defer fake.protocolPortsMutex.Unlock()
	fake.ProtocolPortsStub = nil
	if fake.protocolPortsReturnsOnCall == nil {
		fake.protocolPortsReturnsOnCall = make(map[int]struct {
			result1 map[string]rabbithole.Port
			result2 error
		})
	}
	fake.protocolPortsReturnsOnCall[i] = struct {
		result1 map[string]rabbithole.Port
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) PutPolicy(arg1 string, arg2 string, arg3 rabbithole.Policy) (*http.Response, error) {
	fake.putPolicyMutex.Lock()
	ret, specificReturn := fake.putPolicyReturnsOnCall[len(fake.putPolicyArgsForCall)]
	fake.putPolicyArgsForCall = append(fake.putPolicyArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 rabbithole.Policy
	}{arg1, arg2, arg3})
	fake.recordInvocation("PutPolicy", []interface{}{arg1, arg2, arg3})
	fake.putPolicyMutex.Unlock()
	if fake.PutPolicyStub != nil {
		return fake.PutPolicyStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.putPolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIClient) PutPolicyCallCount() int {
	fake.putPolicyMutex.RLock()
	defer fake.putPolicyMutex.RUnlock()
	return len(fake.putPolicyArgsForCall)
}

func (fake *FakeAPIClient) PutPolicyCalls(stub func(string, string, rabbithole.Policy) (*http.Response, error)) {
	fake.putPolicyMutex.Lock()
	defer fake.putPolicyMutex.Unlock()
	fake.PutPolicyStub = stub
}

func (fake *FakeAPIClient) PutPolicyArgsForCall(i int) (string, string, rabbithole.Policy) {
	fake.putPolicyMutex.RLock()
	defer fake.putPolicyMutex.RUnlock()
	argsForCall := fake.putPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAPIClient) PutPolicyReturns(result1 *http.Response, result2 error) {
	fake.putPolicyMutex.Lock()
	defer fake.putPolicyMutex.Unlock()
	fake.PutPolicyStub = nil
	fake.putPolicyReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) PutPolicyReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.putPolicyMutex.Lock()
	defer fake.putPolicyMutex.Unlock()
	fake.PutPolicyStub = nil
	if fake.putPolicyReturnsOnCall == nil {
		fake.putPolicyReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.putPolicyReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) PutUser(arg1 string, arg2 rabbithole.UserSettings) (*http.Response, error) {
	fake.putUserMutex.Lock()
	ret, specificReturn := fake.putUserReturnsOnCall[len(fake.putUserArgsForCall)]
	fake.putUserArgsForCall = append(fake.putUserArgsForCall, struct {
		arg1 string
		arg2 rabbithole.UserSettings
	}{arg1, arg2})
	fake.recordInvocation("PutUser", []interface{}{arg1, arg2})
	fake.putUserMutex.Unlock()
	if fake.PutUserStub != nil {
		return fake.PutUserStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.putUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIClient) PutUserCallCount() int {
	fake.putUserMutex.RLock()
	defer fake.putUserMutex.RUnlock()
	return len(fake.putUserArgsForCall)
}

func (fake *FakeAPIClient) PutUserCalls(stub func(string, rabbithole.UserSettings) (*http.Response, error)) {
	fake.putUserMutex.Lock()
	defer fake.putUserMutex.Unlock()
	fake.PutUserStub = stub
}

func (fake *FakeAPIClient) PutUserArgsForCall(i int) (string, rabbithole.UserSettings) {
	fake.putUserMutex.RLock()
	defer fake.putUserMutex.RUnlock()
	argsForCall := fake.putUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPIClient) PutUserReturns(result1 *http.Response, result2 error) {
	fake.putUserMutex.Lock()
	defer fake.putUserMutex.Unlock()
	fake.PutUserStub = nil
	fake.putUserReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) PutUserReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.putUserMutex.Lock()
	defer fake.putUserMutex.Unlock()
	fake.PutUserStub = nil
	if fake.putUserReturnsOnCall == nil {
		fake.putUserReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.putUserReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) PutVhost(arg1 string, arg2 rabbithole.VhostSettings) (*http.Response, error) {
	fake.putVhostMutex.Lock()
	ret, specificReturn := fake.putVhostReturnsOnCall[len(fake.putVhostArgsForCall)]
	fake.putVhostArgsForCall = append(fake.putVhostArgsForCall, struct {
		arg1 string
		arg2 rabbithole.VhostSettings
	}{arg1, arg2})
	fake.recordInvocation("PutVhost", []interface{}{arg1, arg2})
	fake.putVhostMutex.Unlock()
	if fake.PutVhostStub != nil {
		return fake.PutVhostStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.putVhostReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIClient) PutVhostCallCount() int {
	fake.putVhostMutex.RLock()
	defer fake.putVhostMutex.RUnlock()
	return len(fake.putVhostArgsForCall)
}

func (fake *FakeAPIClient) PutVhostCalls(stub func(string, rabbithole.VhostSettings) (*http.Response, error)) {
	fake.putVhostMutex.Lock()
	defer fake.putVhostMutex.Unlock()
	fake.PutVhostStub = stub
}

func (fake *FakeAPIClient) PutVhostArgsForCall(i int) (string, rabbithole.VhostSettings) {
	fake.putVhostMutex.RLock()
	defer fake.putVhostMutex.RUnlock()
	argsForCall := fake.putVhostArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPIClient) PutVhostReturns(result1 *http.Response, result2 error) {
	fake.putVhostMutex.Lock()
	defer fake.putVhostMutex.Unlock()
	fake.PutVhostStub = nil
	fake.putVhostReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) PutVhostReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.putVhostMutex.Lock()
	defer fake.putVhostMutex.Unlock()
	fake.PutVhostStub = nil
	if fake.putVhostReturnsOnCall == nil {
		fake.putVhostReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.putVhostReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) UpdatePermissionsIn(arg1 string, arg2 string, arg3 rabbithole.Permissions) (*http.Response, error) {
	fake.updatePermissionsInMutex.Lock()
	ret, specificReturn := fake.updatePermissionsInReturnsOnCall[len(fake.updatePermissionsInArgsForCall)]
	fake.updatePermissionsInArgsForCall = append(fake.updatePermissionsInArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 rabbithole.Permissions
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdatePermissionsIn", []interface{}{arg1, arg2, arg3})
	fake.updatePermissionsInMutex.Unlock()
	if fake.UpdatePermissionsInStub != nil {
		return fake.UpdatePermissionsInStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updatePermissionsInReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIClient) UpdatePermissionsInCallCount() int {
	fake.updatePermissionsInMutex.RLock()
	defer fake.updatePermissionsInMutex.RUnlock()
	return len(fake.updatePermissionsInArgsForCall)
}

func (fake *FakeAPIClient) UpdatePermissionsInCalls(stub func(string, string, rabbithole.Permissions) (*http.Response, error)) {
	fake.updatePermissionsInMutex.Lock()
	defer fake.updatePermissionsInMutex.Unlock()
	fake.UpdatePermissionsInStub = stub
}

func (fake *FakeAPIClient) UpdatePermissionsInArgsForCall(i int) (string, string, rabbithole.Permissions) {
	fake.updatePermissionsInMutex.RLock()
	defer fake.updatePermissionsInMutex.RUnlock()
	argsForCall := fake.updatePermissionsInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAPIClient) UpdatePermissionsInReturns(result1 *http.Response, result2 error) {
	fake.updatePermissionsInMutex.Lock()
	defer fake.updatePermissionsInMutex.Unlock()
	fake.UpdatePermissionsInStub = nil
	fake.updatePermissionsInReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) UpdatePermissionsInReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.updatePermissionsInMutex.Lock()
	defer fake.updatePermissionsInMutex.Unlock()
	fake.UpdatePermissionsInStub = nil
	if fake.updatePermissionsInReturnsOnCall == nil {
		fake.updatePermissionsInReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.updatePermissionsInReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	fake.deleteVhostMutex.RLock()
	defer fake.deleteVhostMutex.RUnlock()
	fake.getVhostMutex.RLock()
	defer fake.getVhostMutex.RUnlock()
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	fake.protocolPortsMutex.RLock()
	defer fake.protocolPortsMutex.RUnlock()
	fake.putPolicyMutex.RLock()
	defer fake.putPolicyMutex.RUnlock()
	fake.putUserMutex.RLock()
	defer fake.putUserMutex.RUnlock()
	fake.putVhostMutex.RLock()
	defer fake.putVhostMutex.RUnlock()
	fake.updatePermissionsInMutex.RLock()
	defer fake.updatePermissionsInMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAPIClient) recordInvocation(key string, args []interface{}) {
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

var _ broker.APIClient = new(FakeAPIClient)
