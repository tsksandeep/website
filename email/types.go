package email

type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

type Body interface {
	GetName() string
	GetEmail() string
	GetMessage() string
	ToString() string
}
