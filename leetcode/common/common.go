package common

type ArrayStack struct {
	datas []string
	size  int
}

func (a *ArrayStack) IsEmpty() bool {
	if a.size == 0 {
		return true
	}

	return false
}

func (a *ArrayStack) Push(data string) {
	a.datas = append(a.datas, data)
	a.size++
}

func (a *ArrayStack) Pop() string {
	if a.IsEmpty() {
		return ""
	}

	tmp := a.datas[a.size-1]
	a.datas = a.datas[:a.size-1]
	a.size--

	return tmp
}
