//go:build linux

package gateway

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	pathRouteFile    = "/proc/net/route"
	flagRouteUp      = 0x0001 /* route usable, based on linux/route.h */
	flagRouteGateway = 0x0002 /* route is gateway, based on linux/route.h */
)

func DetectDefaultGateway() (string, error) {
	routeFile, err := os.OpenFile(pathRouteFile, os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	defer routeFile.Close()

	spaceRemover := regexp.MustCompile("\\s+")

	reader := bufio.NewReader(routeFile)
	_, _, err = reader.ReadLine() // lines are rather short so let's ignore isPrefix, first line can be skipped
	for err == nil {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		modLine := spaceRemover.ReplaceAllString(string(line), " ")
		modLine = strings.Trim(modLine, " ")
		parts := strings.Split(modLine, " ")
		fmt.Println("Number of elements:", len(parts), "Slice content:", parts)

		if isUsableGateway(parts[3]) {
			return convertHexToIP(parts[2]), nil
		}
	}

	return "", fmt.Errorf("usable route not found")
}

func isUsableGateway(flags string) bool {
	hexFlags, err := strconv.ParseInt(flags, 16, 16)
	if err != nil {
		return false
	}

	if (hexFlags & flagRouteGateway) == 0 {
		return false
	}

	if (hexFlags & flagRouteUp) == 0 {
		return false
	}

	return true
}

func convertHexToIP(in string) string {
	strA := in[0:2]
	strB := in[2:4]
	strC := in[4:6]
	strD := in[6:8]

	fmt.Println("CheckPoint:", strA, strB, strC, strD)

	A, err := strconv.ParseUint(strA, 16, 8)
	if err != nil {
		return "0.0.0.0"
	}

	B, err := strconv.ParseUint(strB, 16, 8)
	if err != nil {
		return "0.0.0.0"
	}

	C, err := strconv.ParseUint(strC, 16, 8)
	if err != nil {
		return "0.0.0.0"
	}

	D, err := strconv.ParseUint(strD, 16, 8)
	if err != nil {
		return "0.0.0.0"
	}

	return fmt.Sprintf("%d.%d.%d.%d", D, C, B, A)
}
