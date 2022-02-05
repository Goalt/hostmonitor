package grpc

type Server struct {

}

func (s Server) Run() error {
	return nil
}

func (s Server) Stop() error {
	return nil
}

func NewHTTPServer() Server {
	return Server{}
}