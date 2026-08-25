package model

import "fmt"

var allowedStatuses = map[string]bool{"new": true, "queued": true, "processing": true, "resolved": true, "closed": true, "archived": true}

func ValidateStatus(s string) error {
	if !allowedStatuses[s] {
		return fmt.Errorf("unsupported status %q", s)
	}
	return nil
}
func NextStatus(current, requested string) error {
	if err := ValidateStatus(requested); err != nil {
		return err
	}
	if current == "archived" && requested != "archived" {
		return fmt.Errorf("archived record immutable")
	}
	return nil
}
func NormalizeFilter(f QueryFilter) QueryFilter {
	if f.Limit <= 0 || f.Limit > 100 {
		f.Limit = 100
	}
	return f
}
