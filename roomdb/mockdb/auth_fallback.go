// Code generated by counterfeiter. DO NOT EDIT.
package mockdb

import (
	"context"
	"sync"

	"github.com/ssb-ngi-pointer/go-ssb-room/roomdb"
)

type FakeAuthFallbackService struct {
	CheckStub        func(string, string) (interface{}, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 string
		arg2 string
	}
	checkReturns struct {
		result1 interface{}
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	CreateResetTokenStub        func(context.Context, int64, int64) (string, error)
	createResetTokenMutex       sync.RWMutex
	createResetTokenArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 int64
	}
	createResetTokenReturns struct {
		result1 string
		result2 error
	}
	createResetTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	SetPasswordStub        func(context.Context, int64, string) error
	setPasswordMutex       sync.RWMutex
	setPasswordArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 string
	}
	setPasswordReturns struct {
		result1 error
	}
	setPasswordReturnsOnCall map[int]struct {
		result1 error
	}
	SetPasswordWithTokenStub        func(context.Context, string, string) error
	setPasswordWithTokenMutex       sync.RWMutex
	setPasswordWithTokenArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	setPasswordWithTokenReturns struct {
		result1 error
	}
	setPasswordWithTokenReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthFallbackService) Check(arg1 string, arg2 string) (interface{}, error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CheckStub
	fakeReturns := fake.checkReturns
	fake.recordInvocation("Check", []interface{}{arg1, arg2})
	fake.checkMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthFallbackService) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeAuthFallbackService) CheckCalls(stub func(string, string) (interface{}, error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeAuthFallbackService) CheckArgsForCall(i int) (string, string) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthFallbackService) CheckReturns(result1 interface{}, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthFallbackService) CheckReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthFallbackService) CreateResetToken(arg1 context.Context, arg2 int64, arg3 int64) (string, error) {
	fake.createResetTokenMutex.Lock()
	ret, specificReturn := fake.createResetTokenReturnsOnCall[len(fake.createResetTokenArgsForCall)]
	fake.createResetTokenArgsForCall = append(fake.createResetTokenArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 int64
	}{arg1, arg2, arg3})
	stub := fake.CreateResetTokenStub
	fakeReturns := fake.createResetTokenReturns
	fake.recordInvocation("CreateResetToken", []interface{}{arg1, arg2, arg3})
	fake.createResetTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthFallbackService) CreateResetTokenCallCount() int {
	fake.createResetTokenMutex.RLock()
	defer fake.createResetTokenMutex.RUnlock()
	return len(fake.createResetTokenArgsForCall)
}

func (fake *FakeAuthFallbackService) CreateResetTokenCalls(stub func(context.Context, int64, int64) (string, error)) {
	fake.createResetTokenMutex.Lock()
	defer fake.createResetTokenMutex.Unlock()
	fake.CreateResetTokenStub = stub
}

func (fake *FakeAuthFallbackService) CreateResetTokenArgsForCall(i int) (context.Context, int64, int64) {
	fake.createResetTokenMutex.RLock()
	defer fake.createResetTokenMutex.RUnlock()
	argsForCall := fake.createResetTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAuthFallbackService) CreateResetTokenReturns(result1 string, result2 error) {
	fake.createResetTokenMutex.Lock()
	defer fake.createResetTokenMutex.Unlock()
	fake.CreateResetTokenStub = nil
	fake.createResetTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthFallbackService) CreateResetTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.createResetTokenMutex.Lock()
	defer fake.createResetTokenMutex.Unlock()
	fake.CreateResetTokenStub = nil
	if fake.createResetTokenReturnsOnCall == nil {
		fake.createResetTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createResetTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthFallbackService) SetPassword(arg1 context.Context, arg2 int64, arg3 string) error {
	fake.setPasswordMutex.Lock()
	ret, specificReturn := fake.setPasswordReturnsOnCall[len(fake.setPasswordArgsForCall)]
	fake.setPasswordArgsForCall = append(fake.setPasswordArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.SetPasswordStub
	fakeReturns := fake.setPasswordReturns
	fake.recordInvocation("SetPassword", []interface{}{arg1, arg2, arg3})
	fake.setPasswordMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthFallbackService) SetPasswordCallCount() int {
	fake.setPasswordMutex.RLock()
	defer fake.setPasswordMutex.RUnlock()
	return len(fake.setPasswordArgsForCall)
}

func (fake *FakeAuthFallbackService) SetPasswordCalls(stub func(context.Context, int64, string) error) {
	fake.setPasswordMutex.Lock()
	defer fake.setPasswordMutex.Unlock()
	fake.SetPasswordStub = stub
}

func (fake *FakeAuthFallbackService) SetPasswordArgsForCall(i int) (context.Context, int64, string) {
	fake.setPasswordMutex.RLock()
	defer fake.setPasswordMutex.RUnlock()
	argsForCall := fake.setPasswordArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAuthFallbackService) SetPasswordReturns(result1 error) {
	fake.setPasswordMutex.Lock()
	defer fake.setPasswordMutex.Unlock()
	fake.SetPasswordStub = nil
	fake.setPasswordReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthFallbackService) SetPasswordReturnsOnCall(i int, result1 error) {
	fake.setPasswordMutex.Lock()
	defer fake.setPasswordMutex.Unlock()
	fake.SetPasswordStub = nil
	if fake.setPasswordReturnsOnCall == nil {
		fake.setPasswordReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPasswordReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthFallbackService) SetPasswordWithToken(arg1 context.Context, arg2 string, arg3 string) error {
	fake.setPasswordWithTokenMutex.Lock()
	ret, specificReturn := fake.setPasswordWithTokenReturnsOnCall[len(fake.setPasswordWithTokenArgsForCall)]
	fake.setPasswordWithTokenArgsForCall = append(fake.setPasswordWithTokenArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.SetPasswordWithTokenStub
	fakeReturns := fake.setPasswordWithTokenReturns
	fake.recordInvocation("SetPasswordWithToken", []interface{}{arg1, arg2, arg3})
	fake.setPasswordWithTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthFallbackService) SetPasswordWithTokenCallCount() int {
	fake.setPasswordWithTokenMutex.RLock()
	defer fake.setPasswordWithTokenMutex.RUnlock()
	return len(fake.setPasswordWithTokenArgsForCall)
}

func (fake *FakeAuthFallbackService) SetPasswordWithTokenCalls(stub func(context.Context, string, string) error) {
	fake.setPasswordWithTokenMutex.Lock()
	defer fake.setPasswordWithTokenMutex.Unlock()
	fake.SetPasswordWithTokenStub = stub
}

func (fake *FakeAuthFallbackService) SetPasswordWithTokenArgsForCall(i int) (context.Context, string, string) {
	fake.setPasswordWithTokenMutex.RLock()
	defer fake.setPasswordWithTokenMutex.RUnlock()
	argsForCall := fake.setPasswordWithTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAuthFallbackService) SetPasswordWithTokenReturns(result1 error) {
	fake.setPasswordWithTokenMutex.Lock()
	defer fake.setPasswordWithTokenMutex.Unlock()
	fake.SetPasswordWithTokenStub = nil
	fake.setPasswordWithTokenReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthFallbackService) SetPasswordWithTokenReturnsOnCall(i int, result1 error) {
	fake.setPasswordWithTokenMutex.Lock()
	defer fake.setPasswordWithTokenMutex.Unlock()
	fake.SetPasswordWithTokenStub = nil
	if fake.setPasswordWithTokenReturnsOnCall == nil {
		fake.setPasswordWithTokenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPasswordWithTokenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthFallbackService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.createResetTokenMutex.RLock()
	defer fake.createResetTokenMutex.RUnlock()
	fake.setPasswordMutex.RLock()
	defer fake.setPasswordMutex.RUnlock()
	fake.setPasswordWithTokenMutex.RLock()
	defer fake.setPasswordWithTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthFallbackService) recordInvocation(key string, args []interface{}) {
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

var _ roomdb.AuthFallbackService = new(FakeAuthFallbackService)
