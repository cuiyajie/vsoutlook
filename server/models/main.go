package models

import "vsoutlook.com/vsoutlook/infra/utils"

func AsBasic[T Basicable](t T) any {
	return t.AsBasic()
}

func AsBasics[T Basicable](t []T) []any {
	var result = make([]any, 0, len(t))
	for _, v := range t {
		row := v.AsBasic()
		if row != nil {
			result = append(result, row)
		}
	}
	return result
}

func AsBasicsFor[T BasicableFor](t []T, uid string) []any {
	var result = make([]any, 0, len(t))
	for _, v := range t {
		row := v.AsBasicFor(uid)
		if row != nil {
			result = append(result, row)
		}
	}
	return result
}

func IsDeleted[T Model](t T) bool {
	return t.IsDeleted()
}

func IsActive[T Model](t T) bool {
	return !t.IsDeleted()
}

func FilterActive[T Model](rows []T) []T {
	return utils.Filter(rows, IsActive[T])
}
