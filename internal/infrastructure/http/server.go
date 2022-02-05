package http

type Server struct {

}

func (s Server) Run() error {
	for true {
		
	}
	return nil
}

func (s Server) Stop() error {
	return nil
}

func NewHTTPServer() Server {
	return Server{}
}