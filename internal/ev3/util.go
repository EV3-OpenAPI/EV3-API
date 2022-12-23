package ev3

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"time"
)

var (
	macMappings = map[string]string{
		"74:DA:38:CD:33:C6": "ev3-001",
		"74:DA:38:CD:33:AE": "ev3-002",
		"74:DA:38:F4:51:51": "ev3-003",
		"74:DA:38:CD:33:C8": "ev3-004",
		"74:DA:38:F4:51:5C": "ev3-005",
		"74:DA:38:F4:51:44": "ev3-006", // metl
		"74:DA:38:F4:51:67": "ev3-007",
		"74:DA:38:F4:51:77": "ev3-008",
		"74:DA:38:F4:51:2E": "ev3-009",
		"74:DA:38:F4:51:87": "ev3-010",
		"74:DA:38:CD:33:9E": "ev3-011",
		"74:DA:38:6F:05:CC": "ev3-012",
		"74:DA:38:F4:51:5B": "ev3-013",
		"74:DA:38:F4:51:4B": "ev3-014",
		"74:DA:38:F4:51:73": "ev3-015",
		"74:DA:38:6F:02:D9": "ev3-016",
		"74:DA:38:F4:51:30": "ev3-017",
		"74:DA:38:F4:51:3C": "ev3-018",
		"74:DA:38:CD:33:9C": "ev3-019",
		"74:DA:38:F4:51:5E": "ev3-020",
		"74:DA:38:F4:51:6F": "ev3-021",
		"74:DA:38:F4:51:7C": "ev3-022",
		"74:DA:38:F4:51:6C": "ev3-023",
		"74:DA:38:6F:03:03": "ev3-024",
		"74:DA:38:F4:51:38": "ev3-025",
		"74:DA:38:F4:51:33": "ev3-026",
		"74:DA:38:F4:51:50": "ev3-027",
		"74:DA:38:F4:51:53": "ev3-028",
		"74:DA:38:F4:51:62": "ev3-029",
		"74:DA:38:CD:33:C3": "ev3-030",
		"74:DA:38:F4:51:3D": "ev3-031",
		"74:DA:38:54:7E:03": "ev3-032",
		"74:DA:38:F4:51:32": "ev3-033",
		"74:DA:38:F4:51:75": "ev3-034",
		"74:DA:38:F4:51:57": "ev3-035",
		"74:DA:38:CD:33:C4": "ev3-036",
		"74:DA:38:6F:02:F8": "ev3-037",
		"74:DA:38:6F:02:AC": "ev3-038",
		"74:DA:38:CD:33:AF": "ev3-039",
		"74:DA:38:6F:02:C3": "ev3-040",
		"74:DA:38:F4:51:64": "ev3-041",
		"74:DA:38:CD:33:A5": "ev3-042",
		"74:DA:38:CD:33:A6": "ev3-043",
		"74:DA:38:6F:02:AD": "ev3-044",
		"74:DA:38:F4:51:4E": "ev3-045",
		"74:DA:38:CD:33:C7": "ev3-046",
		"74:DA:38:6F:02:9F": "ev3-047",
		"74:DA:38:CD:33:C9": "ev3-048",
		"74:DA:38:CD:33:B0": "ev3-049",
		"74:DA:38:6F:02:D5": "ev3-050",
		"74:DA:38:CD:33:B1": "ev3-051",
		"74:DA:38:F4:51:3B": "ev3-052",
		"74:DA:38:F4:51:78": "ev3-053",
		"74:DA:38:0E:49:0C": "ev3-054",
		"74:DA:38:6F:05:83": "ev3-055",
		"74:DA:38:F4:51:65": "ev3-056",
		"74:DA:38:F4:51:8D": "ev3-057",
		"74:DA:38:CD:33:CA": "ev3-058",
		"74:DA:38:F4:51:42": "ev3-059",
		"74:DA:38:6F:02:BD": "ev3-060",
		"74:DA:38:6F:02:EE": "ev3-061",
		"74:DA:38:F4:51:86": "ev3-062",
		"74:DA:38:6F:02:BA": "ev3-063",
		"74:DA:38:C9:A8:A6": "ev3-064",
		"74:DA:38:CD:33:B2": "ev3-065",
		"74:DA:38:6F:02:AF": "ev3-066",

		"74:DA:38:6F:02:B1": "ev3-090", // metl
		"74:DA:38:6F:02:E0": "ev3-091", // metl
		"74:DA:38:6F:02:B0": "ev3-092", // metl
	}

	ifIdx = -1
)

func GetHostname() string {
	findWlanInterface()

	netIf, err := net.InterfaceByIndex(ifIdx)
	if err != nil {
		return "ev3-000"
	}

	return GetHostnameForMac(netIf.HardwareAddr.String())
}

func GetHostnameForMac(mac string) string {
	mac = strings.ToUpper(mac)
	if hostname, ok := macMappings[mac]; ok {
		return hostname
	}

	return "ev3-000"
}

func SetSystemHostname(hostname string) error {
	if len(hostname) <= 0 {
		return fmt.Errorf("hostname cannot be empty")
	}

	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo /usr/bin/hostnamectl set-hostname %s", hostname))
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func DurationMs(ms int32) time.Duration {
	return time.Duration(ms) * 1000 * 1000
}

func findWlanInterface() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("ERROR - No WLAN interface found. Info will not be avaiable")
	}

	for _, intf := range interfaces {
		if strings.HasPrefix(intf.Name, "wlx") {
			ifIdx = intf.Index
		}
	}
}

func GetWlanInterfaceIndex() int {
	if ifIdx == -1 {
		findWlanInterface()
	}

	return ifIdx
}
