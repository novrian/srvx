package srvx

type Server struct {
	Name               string
	Host               string
	User               string
	Port               int
	Remote_Tunnel_Port int
	Host_Tunnel_Port   int
}

func (srvx *Server) Label() string {
	var label string
	label = srvx.Name

	if srvx.Host_Tunnel_Port != 0 || srvx.Remote_Tunnel_Port != 0 {
		label += " ["
		if srvx.Host_Tunnel_Port != 0 {
			label += fmt.Sprintf("%d", srvx.Host_Tunnel_Port)
		}
		if srvx.Remote_Tunnel_Port != 0 {
			label += fmt.Sprintf(":%d", srvx.Remote_Tunnel_Port)
		}
		label += "]"
	}

	return label
}

func (server *Server) GenerateArgs() []string {
	host := server.Host
	if server.Port != 0 {
		host += ":" + server.Port
	}
	return []string{""}
}
