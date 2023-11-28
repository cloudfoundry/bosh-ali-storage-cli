// Code generated by counterfeiter. DO NOT EDIT.
package clientfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-ali-storage-cli/client"
)

type FakeStorageClient struct {
	DownloadStub        func(string, string) error
	downloadMutex       sync.RWMutex
	downloadArgsForCall []struct {
		arg1 string
		arg2 string
	}
	downloadReturns struct {
		result1 error
	}
	downloadReturnsOnCall map[int]struct {
		result1 error
	}
	UploadStub        func(string, string) error
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		arg1 string
		arg2 string
	}
	uploadReturns struct {
		result1 error
	}
	uploadReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorageClient) Download(arg1 string, arg2 string) error {
	fake.downloadMutex.Lock()
	ret, specificReturn := fake.downloadReturnsOnCall[len(fake.downloadArgsForCall)]
	fake.downloadArgsForCall = append(fake.downloadArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.DownloadStub
	fakeReturns := fake.downloadReturns
	fake.recordInvocation("Download", []interface{}{arg1, arg2})
	fake.downloadMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStorageClient) DownloadCallCount() int {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return len(fake.downloadArgsForCall)
}

func (fake *FakeStorageClient) DownloadCalls(stub func(string, string) error) {
	fake.downloadMutex.Lock()
	defer fake.downloadMutex.Unlock()
	fake.DownloadStub = stub
}

func (fake *FakeStorageClient) DownloadArgsForCall(i int) (string, string) {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	argsForCall := fake.downloadArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStorageClient) DownloadReturns(result1 error) {
	fake.downloadMutex.Lock()
	defer fake.downloadMutex.Unlock()
	fake.DownloadStub = nil
	fake.downloadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageClient) DownloadReturnsOnCall(i int, result1 error) {
	fake.downloadMutex.Lock()
	defer fake.downloadMutex.Unlock()
	fake.DownloadStub = nil
	if fake.downloadReturnsOnCall == nil {
		fake.downloadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageClient) Upload(arg1 string, arg2 string) error {
	fake.uploadMutex.Lock()
	ret, specificReturn := fake.uploadReturnsOnCall[len(fake.uploadArgsForCall)]
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.UploadStub
	fakeReturns := fake.uploadReturns
	fake.recordInvocation("Upload", []interface{}{arg1, arg2})
	fake.uploadMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStorageClient) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *FakeStorageClient) UploadCalls(stub func(string, string) error) {
	fake.uploadMutex.Lock()
	defer fake.uploadMutex.Unlock()
	fake.UploadStub = stub
}

func (fake *FakeStorageClient) UploadArgsForCall(i int) (string, string) {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	argsForCall := fake.uploadArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStorageClient) UploadReturns(result1 error) {
	fake.uploadMutex.Lock()
	defer fake.uploadMutex.Unlock()
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageClient) UploadReturnsOnCall(i int, result1 error) {
	fake.uploadMutex.Lock()
	defer fake.uploadMutex.Unlock()
	fake.UploadStub = nil
	if fake.uploadReturnsOnCall == nil {
		fake.uploadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStorageClient) recordInvocation(key string, args []interface{}) {
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

var _ client.StorageClient = new(FakeStorageClient)