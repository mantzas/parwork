package parwork

type TestWork struct{}

func (t TestWork) ID() string      { return "1" }
func (t TestWork) Do()             {}
func (t TestWork) GetError() error { return nil }

func testGenerator() Work {
	return TestWork{}
}

func testReporter(w Work) {
}
