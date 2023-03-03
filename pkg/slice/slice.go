// Package slice provides utility functions for working with slices.
package slice

import "reflect"

// ToSet returns the slice as a set.
func ToSet[S ~[]E, E comparable](s S) map[E]bool {
	set := make(map[E]bool, len(s))

	for _, e := range s {
		set[e] = true
	}

	return set
}

// Filter returns a slice of all elements for which keep returns true.
func Filter[S ~[]E, E any](s S, keep func(e E) bool) S {
	var kept S

	for _, e := range s {
		if keep(e) {
			kept = append(kept, e)
		}
	}

	return kept
}

// Difference returns the difference of two slices.
func Difference[S ~[]E, E comparable](a, b S) []E {
	setA := ToSet(a)

	var diff []E

	Each(b, func(e E) {
		if !setA[e] {
			diff = append(diff, e)
		}
	})

	return diff
}

// Match returns the difference of two slices.
func Match[S ~[]E, E comparable](a, b S) []E {
	setA := ToSet(a)

	var diff []E

	Each(b, func(e E) {
		if !setA[e] {
			diff = append(diff, e)
		}
	})

	return diff
}

// Intersection returns the intersection of two slices.
func Intersection[S ~[]E, E comparable](a, b S) []E {
	setA := ToSet(a)

	var diff []E

	Each(b, func(e E) {
		if setA[e] {
			diff = append(diff, e)
		}
	})

	return diff
}

// Each applies fn to each element of the slice.
func Each[S ~[]E, E any](s S, fn func(e E)) {
	for _, e := range s {
		fn(e)
	}
}

// Includes returns true if s is found in the target
func Includes[S ~[]E, E comparable](s S, target E) bool {
	for _, e := range s {
		if e == target {
			return true
		}
	}

	return false
}

// ReflectCompare returns true if a and b are equal
func ReflectCompare[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}
