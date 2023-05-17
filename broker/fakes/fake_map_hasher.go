// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/amansingh066/on-demand-service-broker/broker"
)

type FakeHasher struct {
	HashStub        func(map[string]string) string
	hashMutex       sync.RWMutex
	hashArgsForCall []struct {
		arg1 map[string]string
	}
	hashReturns struct {
		result1 string
	}
	hashReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHasher) Hash(arg1 map[string]string) string {
	fake.hashMutex.Lock()
	ret, specificReturn := fake.hashReturnsOnCall[len(fake.hashArgsForCall)]
	fake.hashArgsForCall = append(fake.hashArgsForCall, struct {
		arg1 map[string]string
	}{arg1})
	stub := fake.HashStub
	fakeReturns := fake.hashReturns
	fake.recordInvocation("Hash", []interface{}{arg1})
	fake.hashMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHasher) HashCallCount() int {
	fake.hashMutex.RLock()
	defer fake.hashMutex.RUnlock()
	return len(fake.hashArgsForCall)
}

func (fake *FakeHasher) HashCalls(stub func(map[string]string) string) {
	fake.hashMutex.Lock()
	defer fake.hashMutex.Unlock()
	fake.HashStub = stub
}

func (fake *FakeHasher) HashArgsForCall(i int) map[string]string {
	fake.hashMutex.RLock()
	defer fake.hashMutex.RUnlock()
	argsForCall := fake.hashArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHasher) HashReturns(result1 string) {
	fake.hashMutex.Lock()
	defer fake.hashMutex.Unlock()
	fake.HashStub = nil
	fake.hashReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeHasher) HashReturnsOnCall(i int, result1 string) {
	fake.hashMutex.Lock()
	defer fake.hashMutex.Unlock()
	fake.HashStub = nil
	if fake.hashReturnsOnCall == nil {
		fake.hashReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.hashReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeHasher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.hashMutex.RLock()
	defer fake.hashMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHasher) recordInvocation(key string, args []interface{}) {
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

var _ broker.Hasher = new(FakeHasher)
