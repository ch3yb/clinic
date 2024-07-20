package utils

func GetStringPointer(s string) *string {
	return &s
}

func GetIntPointer(t int64) *int {
	var x = int(t)
	return &x
}
func GetInt64Pointer(t int64) *int64 {
	return &t
}
