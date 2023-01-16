package windows

import "fmt"

const WindowLengthError = "window length M must be a non-negative integer"

// _len_guards checks the window length
func _len_guards(M int) (bool, error) {
	if M != int(M) || M < 0 {
		return true, fmt.Errorf(WindowLengthError)
	}
	return M <= 1, nil
}

// _extend extends the window by 1 sample if needed for DFT-even symmetry
func _extend(M int, sym bool) (int, bool) {
	/* Extend window by 1 sample if needed for DFT-even symmetry */
	if !sym {
		return M + 1, true
	} else {
		return M, false
	}
}

// _truncate truncates the window if needed
func _truncate(w []float64, needed bool) []float64 {
	if needed {
		return w[:len(w)-1]
	} else {
		return w
	}
}
