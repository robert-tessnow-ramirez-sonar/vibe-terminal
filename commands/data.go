package commands

type Command struct {
	Name    string
	Execute string
}

// Categories holds our command groups mapped to macOS-friendly commands
var Categories = map[string][]Command{
	"Networking": {
		{"List Interfaces", "ifconfig"},
		{"Active Connections", "lsof -i -P -n | grep LISTEN"},
		{"DNS Lookup", "dig google.com"},
		{"Ping Router", "ping -c 4 192.168.1.1"},
	},
	"System": {
		{"Disk Usage", "df -h"},
		{"Memory Stats", "top -l 1 | grep PhysMem"}, // macOS alternative to free -m
		{"Hardware Info", "system_profiler SPHardwareDataType"},
	},
}

func GetCategoryNames() []string {
	// Hardcoding the order guarantees the UI dropdown always looks the same
	return []string{"Networking", "System"}
}
