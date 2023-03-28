package main

func f1() {
	println("in defer function2")
}

func deferCase1() {
	println("start defer case 1")

	defer func() {
		println("in defer function1")
	}()
	defer f1()
	println("defer case end")
}
//参数预计算
func deferCase2(){
  i := 1 
  //传参
  defer func(j int){
    println("defer j:",j)
  }(i + 1)
  //闭包
  defer func(){
    println("defer j:",i)
  }()
  i = 99
  println("i: ",i)
}
func main() {
	deferCase1()
  deferCase2()
}
