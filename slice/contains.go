package slice

import "reflect"

func ContainsString(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func Contains[T any](sl []T, v T) bool {
	var _v interface{} = v
	for _, vv := range sl {
		var _vv interface{} = vv
		if _vv == _v {
			return true
		}
		//if fmt.Sprintf("%v", vv) == fmt.Sprintf("%v", v) {
		//	return true
		//}
	}
	return false
}

func Equals[T any](source, target []T) bool {
	// If one is nil, the other must also be nil.
	if (source == nil) != (target == nil) {
		return false
	}

	if len(source) != len(target) {
		return false
	}

	for i := range source {
		blnEq := reflect.DeepEqual(source[i], target[i])
		if !blnEq {
			return false
		}
	}

	return true
}

// HasCommon source里面的元素在target中存在，则返回true
func HasCommon[T any](source, target []T) bool {
	for _, v := range source {
		if Contains(target, v) {
			return true
		}
	}
	return false
}

// DiffSlice 返回两个slice的不同的元素, 以source为标准
// 返回在target中不存在的source元素列表
func DiffSlice[T any](source, target []T) (diffSlice []T) {
	for _, v := range source {
		if !Contains(target, v) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

// CommonSlice 返回两个slice相同的元素, 以source为标准
// 返回在target中存在的source元素列表
func CommonSlice[T any](source, target []T) (commonSlice []T) {
	for _, v := range source {
		if Contains(target, v) {
			commonSlice = append(commonSlice, v)
		}
	}
	return
}
