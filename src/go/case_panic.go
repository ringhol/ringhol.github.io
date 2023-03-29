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

//返回值，defer执行在returun之后
var g = 100
func f2()(int,*int){
  defer func(){
    g = 200
  }()
  println("f2 g: ",g)
  return g,&g
}

func deferCase3(){
  i,j := f2()
  println("i,j,g: ",i,*j,g)
}

func exceptionCase(){
  defer func(){
    err := recover()
    if err != nil{
      println("error dealing...  defer recover:",err)
    }
  }()
  println("start exceptionCase")
  panic("exceptionCase error")
  println("end exceptionCase")
}

func main() {
	// deferCase1()
	//   deferCase2()
	  // deferCase3()
  exceptionCase()
}
