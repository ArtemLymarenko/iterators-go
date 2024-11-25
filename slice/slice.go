package slice

import (
	"iter"
)

type iterator[T comparable] iter.Seq2[int, T]

func From[T comparable](s []T) iterator[T] {
	return func(yield func(idx int, value T) bool) {
		for idx, value := range s {
			if !yield(idx, value) {
				return
			}
		}
	}
}

func (iterator iterator[T]) Collect() (result []T) {
	for _, v := range iterator {
		result = append(result, v)
	}

	return result
}

func (iterator iterator[T]) ForEach(operate func(idx int, value T)) iterator[T] {
	for idx, value := range iterator {
		operate(idx, value)
	}

	return iterator
}

func (iterator iterator[T]) Map(operate func(idx int, value T) T) iterator[T] {
	return func(yield func(i int, value T) bool) {
		for idx, value := range iterator {
			if !yield(idx, operate(idx, value)) {
				return
			}
		}
	}
}

func (iterator iterator[T]) Reverse() iterator[T] {
	arr := iterator.Collect()
	for i, j := 0, len(arr)-1; i < len(arr)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return From(arr)
}

func (iterator iterator[T]) Filter(validate func(value T) bool) iterator[T] {
	return func(yield func(i int, value T) bool) {
		for idx, value := range iterator {
			if validate(value) && !yield(idx, value) {
				return
			}
		}
	}
}

func (iterator iterator[T]) Reduce(reduce func(acc T, current T) T, initValue T) T {
	acc := initValue
	for _, value := range iterator {
		acc = reduce(acc, value)
	}

	return acc
}

func (iterator iterator[T]) Find(validate func(value T) bool) (idx int, value T) {
	for idx, value = range iterator {
		if validate(value) {
			return idx, value
		}
	}

	return idx, value
}

func (iterator iterator[T]) Every(validate func(value T) bool) bool {
	for _, value := range iterator {
		if !validate(value) {
			return false
		}
	}

	return true
}

func (iterator iterator[T]) Fill(valueToFill T, startIdx, endIdx int) iterator[T] {
	return func(yield func(i int, value T) bool) {
		for idx, _ := range iterator {
			if idx >= startIdx || idx <= endIdx {
				if !yield(idx, valueToFill) {
					return
				}
			}
		}
	}
}
