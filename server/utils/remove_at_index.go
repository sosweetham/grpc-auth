package utils

import "errors"

func RemoveAtIndex[T any](slice []T, s int) ([]T, error) {
    if s < 0 {
        return nil, errors.New(`index cannot be negative`)
    }
    if s > len(slice) {
        return nil, errors.New(`index out of slice bounds`)
    }
    return append(slice[:s], slice[s+1:]...), nil
}