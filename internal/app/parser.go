package app

// // bytesToGoString конвертирует []byte в строку вида [200, 241, 15, ...]
// func bytesToGoString(data []byte) string {
// 	if len(data) == 0 {
// 		return "[]"
// 	}

// 	var builder strings.Builder
// 	builder.WriteString("[")

// 	for i, b := range data {
// 		if i > 0 {
// 			builder.WriteString(", ")
// 		}
// 		builder.WriteString(strconv.Itoa(int(b)))
// 	}

// 	builder.WriteString("]")
// 	return builder.String()
// }

//// parseByteArray парсит строку вида [200, 241, 15, ...] в []byte
// func parseByteArray(s string) ([]byte, error) {
// 	s = strings.TrimSpace(s)

// 	if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
// 		return nil, fmt.Errorf("invalid format: missing brackets")
// 	}
// 	s = s[1 : len(s)-1]

// 	if strings.TrimSpace(s) == "" {
// 		return []byte{}, nil
// 	}

// 	parts := strings.Split(s, ",")
// 	result := make([]byte, 0, len(parts))

// 	for _, part := range parts {
// 		part = strings.TrimSpace(part)
// 		if part == "" {
// 			continue
// 		}

// 		num, err := strconv.ParseUint(part, 10, 8)
// 		if err != nil {
// 			return nil, fmt.Errorf("invalid byte value: %s", part)
// 		}

// 		result = append(result, byte(num))
// 	}

// 	return result, nil
// }
