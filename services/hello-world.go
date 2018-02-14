package services

// HelloWorldService contains methods for finding and manipulating resources
type HelloWorldService struct {
}

// NewHelloWorldService creates an instance of HelloWorldService
func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}

// SayHello say hello
func (hs HelloWorldService) SayHello() string {
	return "Hello World!"
}
