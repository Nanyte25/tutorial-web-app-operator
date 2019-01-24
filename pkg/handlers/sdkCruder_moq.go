// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handlers

import (
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"sync"
)

var (
	lockSdkCruderMockCreate sync.RWMutex
	lockSdkCruderMockDelete sync.RWMutex
	lockSdkCruderMockGet    sync.RWMutex
	lockSdkCruderMockList   sync.RWMutex
	lockSdkCruderMockUpdate sync.RWMutex
)

// SdkCruderMock is a mock implementation of SdkCruder.
//
//     func TestSomethingThatUsesSdkCruder(t *testing.T) {
//
//         // make and configure a mocked SdkCruder
//         mockedSdkCruder := &SdkCruderMock{
//             CreateFunc: func(object sdk.Object) error {
// 	               panic("mock out the Create method")
//             },
//             DeleteFunc: func(object sdk.Object, opts ...sdk.DeleteOption) error {
// 	               panic("mock out the Delete method")
//             },
//             GetFunc: func(object sdk.Object, opts ...sdk.GetOption) error {
// 	               panic("mock out the Get method")
//             },
//             ListFunc: func(namespace string, into sdk.Object, opts ...sdk.ListOption) error {
// 	               panic("mock out the List method")
//             },
//             UpdateFunc: func(object sdk.Object) error {
// 	               panic("mock out the Update method")
//             },
//         }
//
//         // use mockedSdkCruder in code that requires SdkCruder
//         // and then make assertions.
//
//     }
type SdkCruderMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(object sdk.Object) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(object sdk.Object, opts ...sdk.DeleteOption) error

	// GetFunc mocks the Get method.
	GetFunc func(object sdk.Object, opts ...sdk.GetOption) error

	// ListFunc mocks the List method.
	ListFunc func(namespace string, into sdk.Object, opts ...sdk.ListOption) error

	// UpdateFunc mocks the Update method.
	UpdateFunc func(object sdk.Object) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Object is the object argument value.
			Object sdk.Object
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Object is the object argument value.
			Object sdk.Object
			// Opts is the opts argument value.
			Opts []sdk.DeleteOption
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Object is the object argument value.
			Object sdk.Object
			// Opts is the opts argument value.
			Opts []sdk.GetOption
		}
		// List holds details about calls to the List method.
		List []struct {
			// Namespace is the namespace argument value.
			Namespace string
			// Into is the into argument value.
			Into sdk.Object
			// Opts is the opts argument value.
			Opts []sdk.ListOption
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Object is the object argument value.
			Object sdk.Object
		}
	}
}

// Create calls CreateFunc.
func (mock *SdkCruderMock) Create(object sdk.Object) error {
	if mock.CreateFunc == nil {
		panic("SdkCruderMock.CreateFunc: method is nil but SdkCruder.Create was just called")
	}
	callInfo := struct {
		Object sdk.Object
	}{
		Object: object,
	}
	lockSdkCruderMockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	lockSdkCruderMockCreate.Unlock()
	return mock.CreateFunc(object)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedSdkCruder.CreateCalls())
func (mock *SdkCruderMock) CreateCalls() []struct {
	Object sdk.Object
} {
	var calls []struct {
		Object sdk.Object
	}
	lockSdkCruderMockCreate.RLock()
	calls = mock.calls.Create
	lockSdkCruderMockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *SdkCruderMock) Delete(object sdk.Object, opts ...sdk.DeleteOption) error {
	if mock.DeleteFunc == nil {
		panic("SdkCruderMock.DeleteFunc: method is nil but SdkCruder.Delete was just called")
	}
	callInfo := struct {
		Object sdk.Object
		Opts   []sdk.DeleteOption
	}{
		Object: object,
		Opts:   opts,
	}
	lockSdkCruderMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockSdkCruderMockDelete.Unlock()
	return mock.DeleteFunc(object, opts...)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedSdkCruder.DeleteCalls())
func (mock *SdkCruderMock) DeleteCalls() []struct {
	Object sdk.Object
	Opts   []sdk.DeleteOption
} {
	var calls []struct {
		Object sdk.Object
		Opts   []sdk.DeleteOption
	}
	lockSdkCruderMockDelete.RLock()
	calls = mock.calls.Delete
	lockSdkCruderMockDelete.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *SdkCruderMock) Get(object sdk.Object, opts ...sdk.GetOption) error {
	if mock.GetFunc == nil {
		panic("SdkCruderMock.GetFunc: method is nil but SdkCruder.Get was just called")
	}
	callInfo := struct {
		Object sdk.Object
		Opts   []sdk.GetOption
	}{
		Object: object,
		Opts:   opts,
	}
	lockSdkCruderMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockSdkCruderMockGet.Unlock()
	return mock.GetFunc(object, opts...)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedSdkCruder.GetCalls())
func (mock *SdkCruderMock) GetCalls() []struct {
	Object sdk.Object
	Opts   []sdk.GetOption
} {
	var calls []struct {
		Object sdk.Object
		Opts   []sdk.GetOption
	}
	lockSdkCruderMockGet.RLock()
	calls = mock.calls.Get
	lockSdkCruderMockGet.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *SdkCruderMock) List(namespace string, into sdk.Object, opts ...sdk.ListOption) error {
	if mock.ListFunc == nil {
		panic("SdkCruderMock.ListFunc: method is nil but SdkCruder.List was just called")
	}
	callInfo := struct {
		Namespace string
		Into      sdk.Object
		Opts      []sdk.ListOption
	}{
		Namespace: namespace,
		Into:      into,
		Opts:      opts,
	}
	lockSdkCruderMockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	lockSdkCruderMockList.Unlock()
	return mock.ListFunc(namespace, into, opts...)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedSdkCruder.ListCalls())
func (mock *SdkCruderMock) ListCalls() []struct {
	Namespace string
	Into      sdk.Object
	Opts      []sdk.ListOption
} {
	var calls []struct {
		Namespace string
		Into      sdk.Object
		Opts      []sdk.ListOption
	}
	lockSdkCruderMockList.RLock()
	calls = mock.calls.List
	lockSdkCruderMockList.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *SdkCruderMock) Update(object sdk.Object) error {
	if mock.UpdateFunc == nil {
		panic("SdkCruderMock.UpdateFunc: method is nil but SdkCruder.Update was just called")
	}
	callInfo := struct {
		Object sdk.Object
	}{
		Object: object,
	}
	lockSdkCruderMockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	lockSdkCruderMockUpdate.Unlock()
	return mock.UpdateFunc(object)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedSdkCruder.UpdateCalls())
func (mock *SdkCruderMock) UpdateCalls() []struct {
	Object sdk.Object
} {
	var calls []struct {
		Object sdk.Object
	}
	lockSdkCruderMockUpdate.RLock()
	calls = mock.calls.Update
	lockSdkCruderMockUpdate.RUnlock()
	return calls
}
