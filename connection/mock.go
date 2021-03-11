package connection

type MockConnection struct {
	client *MockClient
}

type MockClient struct{}

func NewMockConnection(config Config) (conn Connection, err error) {
	conn = MockConnection{client: &MockClient{}}
	return
}

func (conn MockConnection) Interface() (i interface{}) {
	return conn.client
}

func (conn MockConnection) Close() (err error) {
	return
}
