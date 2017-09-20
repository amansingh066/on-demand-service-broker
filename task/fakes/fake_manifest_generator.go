// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"log"
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/task"
)

type FakeManifestGenerator struct {
	GenerateManifestStub        func(deploymentName, planID string, requestParams map[string]interface{}, oldManifest []byte, previousPlanID *string, logger *log.Logger) (task.RawBoshManifest, error)
	generateManifestMutex       sync.RWMutex
	generateManifestArgsForCall []struct {
		deploymentName string
		planID         string
		requestParams  map[string]interface{}
		oldManifest    []byte
		previousPlanID *string
		logger         *log.Logger
	}
	generateManifestReturns struct {
		result1 task.RawBoshManifest
		result2 error
	}
	generateManifestReturnsOnCall map[int]struct {
		result1 task.RawBoshManifest
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifestGenerator) GenerateManifest(deploymentName string, planID string, requestParams map[string]interface{}, oldManifest []byte, previousPlanID *string, logger *log.Logger) (task.RawBoshManifest, error) {
	var oldManifestCopy []byte
	if oldManifest != nil {
		oldManifestCopy = make([]byte, len(oldManifest))
		copy(oldManifestCopy, oldManifest)
	}
	fake.generateManifestMutex.Lock()
	ret, specificReturn := fake.generateManifestReturnsOnCall[len(fake.generateManifestArgsForCall)]
	fake.generateManifestArgsForCall = append(fake.generateManifestArgsForCall, struct {
		deploymentName string
		planID         string
		requestParams  map[string]interface{}
		oldManifest    []byte
		previousPlanID *string
		logger         *log.Logger
	}{deploymentName, planID, requestParams, oldManifestCopy, previousPlanID, logger})
	fake.recordInvocation("GenerateManifest", []interface{}{deploymentName, planID, requestParams, oldManifestCopy, previousPlanID, logger})
	fake.generateManifestMutex.Unlock()
	if fake.GenerateManifestStub != nil {
		return fake.GenerateManifestStub(deploymentName, planID, requestParams, oldManifest, previousPlanID, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.generateManifestReturns.result1, fake.generateManifestReturns.result2
}

func (fake *FakeManifestGenerator) GenerateManifestCallCount() int {
	fake.generateManifestMutex.RLock()
	defer fake.generateManifestMutex.RUnlock()
	return len(fake.generateManifestArgsForCall)
}

func (fake *FakeManifestGenerator) GenerateManifestArgsForCall(i int) (string, string, map[string]interface{}, []byte, *string, *log.Logger) {
	fake.generateManifestMutex.RLock()
	defer fake.generateManifestMutex.RUnlock()
	return fake.generateManifestArgsForCall[i].deploymentName, fake.generateManifestArgsForCall[i].planID, fake.generateManifestArgsForCall[i].requestParams, fake.generateManifestArgsForCall[i].oldManifest, fake.generateManifestArgsForCall[i].previousPlanID, fake.generateManifestArgsForCall[i].logger
}

func (fake *FakeManifestGenerator) GenerateManifestReturns(result1 task.RawBoshManifest, result2 error) {
	fake.GenerateManifestStub = nil
	fake.generateManifestReturns = struct {
		result1 task.RawBoshManifest
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestGenerator) GenerateManifestReturnsOnCall(i int, result1 task.RawBoshManifest, result2 error) {
	fake.GenerateManifestStub = nil
	if fake.generateManifestReturnsOnCall == nil {
		fake.generateManifestReturnsOnCall = make(map[int]struct {
			result1 task.RawBoshManifest
			result2 error
		})
	}
	fake.generateManifestReturnsOnCall[i] = struct {
		result1 task.RawBoshManifest
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateManifestMutex.RLock()
	defer fake.generateManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManifestGenerator) recordInvocation(key string, args []interface{}) {
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

var _ task.ManifestGenerator = new(FakeManifestGenerator)
