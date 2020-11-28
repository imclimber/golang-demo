package utils

// SwapByPointers ...
func SwapByPointers(f *int, s *int) {
	if *f > *s {
		tmp := *f
		*f = *s
		*s = tmp
	}
}


// SwapDirectly ...
func SwapDirectly(f *int, s *int){
	*f, *s = *s, *f
}