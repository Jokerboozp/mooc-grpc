package handler

type HelloService struct {
}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

const HelloServiceName = "handler/HelloService"
