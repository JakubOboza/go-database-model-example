package users

//UserStoreMock base mock struct
type UserStoreMock struct {
	calls           map[string]int
	resultCallbacks map[string]resultCallbackFunction
	params          map[string][]interface{}
}

//SetupMock mocks and returns pointer to it, each call = fresh mock
func SetupMock() *UserStoreMock {
	mock := &UserStoreMock{}
	mock.calls = make(map[string]int)
	mock.resultCallbacks = make(map[string]resultCallbackFunction)
	mock.params = make(map[string][]interface{})
	Store = mock
	return mock
}

type resultCallbackFunction func() (*User, error)

//RegisterResultCallback registers callabcks that give us results we want in testing
func (mock *UserStoreMock) RegisterResultCallback(name string, foo resultCallbackFunction) {
	mock.resultCallbacks[name] = foo
}

func (mock *UserStoreMock) registerCall(name string, params ...interface{}) {
	mock.calls[name]++
	mock.params[name] = params
}

//CallsCount returns how many times given function was called
func (mock *UserStoreMock) CallsCount(name string) int {
	return mock.calls[name]
}

// INTERFACE IMPLEMENTATION:

//FindByEmail implements find mock
func (mock *UserStoreMock) FindByEmail(email string) (*User, error) {
	mock.registerCall("FindByEmail", email)
	return mock.resultCallbacks["FindByEmail"]()
}

//CreateUser implements create mock
func (mock *UserStoreMock) CreateUser(email string, password string) (*User, error) {
	mock.registerCall("CreateUser", email, password)
	return mock.resultCallbacks["CreateUser"]()
}
