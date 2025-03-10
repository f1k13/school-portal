package utils

func PtrToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
