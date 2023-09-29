package operator

func Ternary[T any](isTrue bool, ifTrue T, ifFalse T) T {
	if isTrue {
		return ifTrue
	}
	return ifFalse
}
