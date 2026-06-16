package util

import "encoding/json"

// RepairInvalidJSONStringEscapes repairs common invalid JSON string escapes
// such as Windows paths emitted as "C:\Users\name". It only returns changed
// output when the repaired payload is valid JSON.
func RepairInvalidJSONStringEscapes(payload []byte) ([]byte, bool) {
	if json.Valid(payload) {
		return payload, false
	}
	repaired, changed := repairJSONStringEscapes(payload)
	if !changed || !json.Valid(repaired) {
		return payload, false
	}
	return repaired, true
}

func repairJSONStringEscapes(payload []byte) ([]byte, bool) {
	out := make([]byte, 0, len(payload)+8)
	changed := false
	inString := false
	repairRestOfString := false
	for i := 0; i < len(payload); i++ {
		c := payload[i]
		if !inString {
			out = append(out, c)
			if c == '"' {
				inString = true
				repairRestOfString = false
			}
			continue
		}
		if c == '"' {
			inString = false
			repairRestOfString = false
			out = append(out, c)
			continue
		}
		if c != '\\' {
			out = append(out, c)
			continue
		}
		if i+1 >= len(payload) {
			out = append(out, '\\', '\\')
			changed = true
			continue
		}
		next := payload[i+1]
		if repairRestOfString {
			if next == '\\' || next == '"' {
				out = append(out, '\\', next)
				i++
			} else {
				out = append(out, '\\', '\\')
				changed = true
			}
			continue
		}
		switch next {
		case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
			out = append(out, '\\', next)
			i++
		case 'u':
			if i+5 < len(payload) && isJSONHex(payload[i+2]) && isJSONHex(payload[i+3]) && isJSONHex(payload[i+4]) && isJSONHex(payload[i+5]) {
				out = append(out, payload[i:i+6]...)
				i += 5
			} else {
				out = append(out, '\\', '\\')
				changed = true
				repairRestOfString = true
			}
		default:
			out = append(out, '\\', '\\')
			changed = true
			repairRestOfString = true
		}
	}
	return out, changed
}

func isJSONHex(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}
