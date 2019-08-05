package main

import (
	"gin_demo/pkg/setting"
	"gin_demo/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

//func main() {
//	m1 := make([]string, 1)
//	m1[0] = "test"
//	fmt.Println("调用 func1 前 m1 值:", m1, cap(m1))
//	func1(m1)
//	fmt.Println("调用 func1 后 m1 值:", m1, cap(m1))
//
//	m2 := make(map[string]string, 2)
//	m2["x"] = "test"
//	fmt.Println("调用 func2 前 m2 值:", m2)
//	func2(m2)
//	fmt.Println("调用 func2 后 m2 值:", m2)
//
//
//}
//
//func func1 (a []string) {
//	a = append(a, "val1")
//	a[0] = "sdfsd"
//	fmt.Println("func1中:", a, cap(a))
//}
//
//func func2 (a map[string]string) {
//	a["x"] = "xxx"
//	a["y"] = "yyy"
//	fmt.Println("func2中:", a)
//}
