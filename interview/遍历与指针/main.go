package main

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhao", Age: 24},
		{Name: "chen", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus { // 使用零时变量stu的地址来传给m的，而且零时变量stu每次的地址都是不会变的
		m[stu.Name] = &stu
	}

	for k, v := range m {
		println(k, "--", v.Name)
	}
}
