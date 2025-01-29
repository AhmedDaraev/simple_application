package handlers

// Преобразует int в *int
func toIntPtr(value int) *int {
	return &value
}

// Преобразует string в *string
func toStringPtr(value string) *string {
	return &value
}

// Преобразует bool в *bool
func toBoolPtr(value bool) *bool {
	return &value
}
