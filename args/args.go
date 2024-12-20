package argUtils

import "flag"

type Settings struct {
	StartPort     *int
	EndPort       *int
	IP            *string
	Timeout       *int
	Retry         *int
	Delay         *int
	Protocol      *string
	ListOpenPorts *bool
}

func ParseArgs() Settings {
	parsed := Settings{}

	parsed.EndPort = flag.Int("end", 65535, "The last port to search. Searches in a range from start port to end port")
	parsed.StartPort = flag.Int("start", 0, "The starting port to search. Searches in a range from start port to end port.")
	parsed.Timeout = flag.Int("timeout", 0, "Time in seconds to wait before timing out if a port is unavailable.")
	parsed.IP = flag.String("ip", "127.0.0.1", "The IP to check open ports on.")
	parsed.Delay = flag.Int("delay", 0, "Delay in seconds before connecting to another port")
	parsed.Retry = flag.Int("retrycon", 1, "How many times to retry a port.")
	parsed.Protocol = flag.String("protocol", "tcp", "Protocol to search ports on. (e.g tcp, udp)")
	parsed.ListOpenPorts = flag.Bool("print", false, "Print port number if it is open.")
	flag.Parse()

	return parsed
}
