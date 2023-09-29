package Int32ToIPv4

import "fmt"

func Int32ToIp(n uint32) string {
	octet1 := (n & 0xFF000000) >> 24
	octet2 := (n & 0x00FF0000) >> 16
	octet3 := (n & 0x0000FF00) >> 8
	octet4 := n & 0x000000FF

	return fmt.Sprintf("%d.%d.%d.%d", octet1, octet2, octet3, octet4)
}
