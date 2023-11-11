package day_13

import (
	"encoding/json"
	"strings"
)

func parse(input string) [][]interface{} {
	lines := strings.Split(input, "\n")
	var out [][]interface{}

	for _, line := range lines {
		if line != "" {
			var lineSlice []interface{}
			err := json.Unmarshal([]byte(line), &lineSlice)
			if err == nil {
				out = append(out, lineSlice)
			}
		}
	}

	return out
}

const (
	Unknown = "UNKNOWN"
	Ok      = "OK"
	Nok     = "NOK"
)

func comparePackets(packet1, packet2 []interface{}) string {
	// Compare all elements one-by-one, until we find an answer or either one runs out
	for i := 0; i < len(packet1) && i < len(packet2); i++ {

		switch element1 := packet1[i].(type) {

		case float64: // Int and ?
			if element2, ok := packet2[i].(float64); ok { // Int and int
				if element1 > element2 {
					return Nok
				}
				if element1 < element2 {
					return Ok
				}
			} else { // Int and list
				out := comparePackets([]interface{}{element1}, packet2[i].([]interface{}))
				if out == Ok || out == Nok {
					return out
				}
			}
		case []interface{}: // List and ?
			if element2, ok := packet2[i].(float64); ok { // List and int
				switch comparePackets(element1, []interface{}{element2}) {
				case Ok:
					return Ok
				case Nok:
					return Nok
				case Unknown:
					continue
				}
			} else { // List and list
				switch comparePackets(element1, packet2[i].([]interface{})) {
				case Ok:
					return Ok
				case Nok:
					return Nok
				case Unknown:
					continue
				}
			}
		}
	}

	// Either list has run out of elements and we have not found an answer
	if len(packet1) < len(packet2) {
		return Ok
	} else if len(packet1) > len(packet2) {
		return Nok
	} else {
		return Unknown
	}
}
