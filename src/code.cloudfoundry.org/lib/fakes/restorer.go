// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type Restorer struct {
	RestoreStub        func(string) error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct {
		arg1 string
	}
	restoreReturns struct {
		result1 error
	}
	restoreReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Restorer) Restore(arg1 string) error {
	fake.restoreMutex.Lock()
	ret, specificReturn := fake.restoreReturnsOnCall[len(fake.restoreArgsForCall)]
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RestoreStub
	fakeReturns := fake.restoreReturns
	fake.recordInvocation("Restore", []interface{}{arg1})
	fake.restoreMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Restorer) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *Restorer) RestoreCalls(stub func(string) error) {
	fake.restoreMutex.Lock()
	defer fake.restoreMutex.Unlock()
	fake.RestoreStub = stub
}

func (fake *Restorer) RestoreArgsForCall(i int) string {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	argsForCall := fake.restoreArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Restorer) RestoreReturns(result1 error) {
	fake.restoreMutex.Lock()
	defer fake.restoreMutex.Unlock()
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *Restorer) RestoreReturnsOnCall(i int, result1 error) {
	fake.restoreMutex.Lock()
	defer fake.restoreMutex.Unlock()
	fake.RestoreStub = nil
	if fake.restoreReturnsOnCall == nil {
		fake.restoreReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restoreReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Restorer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Restorer) recordInvocation(key string, args []interface{}) {
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
