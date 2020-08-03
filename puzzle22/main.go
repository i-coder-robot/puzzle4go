package main
//[intermediate] 从切片中删除一个元素，下面的算法实现正确的是（）
//A.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			if i== len(*s) - 1 {
//				*s = (*s)[:i]
//			}else {
//				*s = append((*s)[:i],(*s)[i + 2:]...)
//			}
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}

//B.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			*s = append((*s)[:i],(*s)[i + 1:])
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}

//C.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			delete(*s, v)
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}

//D.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			*s = append((*s)[:i],(*s)[i + 1:]...)
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}
//参考答案：D

//整个逻辑就是，找到你要删除的元素，然后把它前面的和后面的组合起来

func main() {
	
}
