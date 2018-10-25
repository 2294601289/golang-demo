package main
import "fmt"
func main(){
/*
	fmt.Println("Hello,World!")
	//定义单个变量
	var xx int
	xx = 11
	fmt.Println(xx)
	var xxx string = "gg"
	fmt.Println(xxx);
	//定义多变量
	var x1,x2,x3 int
	x1=1
	x2=2
	x3=3
	fmt.Println(x1+x2+x3)
	fmt.Println(xx,x1,x2,x3)
	//定义常量
	const c1 int=110
	fmt.Println(c1)
	//运算符
	x1=x1+x2
	x2=x2*x3
	fmt.Println(x1,x2,x1==x2,x1!=x3,x1>x2,true&&false,true||false,!false)
*/
/*
	//if语句
	var a1 int = 1
	if 1==2 {
		a1 = 2
	}
	fmt.Println(a1)
	//if else 语句
	if a1>0 {
		fmt.Println("a1大于0")
	}else if a1>-10 {
		fmt.Println("a1小于0大于-10")
	}else{
		fmt.Println("a1小于-10")
	}
	//switch语句
	switch a1 {
	case 1: fmt.Println("a1为1") 
	case 2: fmt.Println("a1为1") 
	case 3: fmt.Println("a1为1") 
	case 4: fmt.Println("a1为1") 
	case 5: fmt.Println("a1为1") 
	default:fmt.Println("switch默认执行")
	}
	//for 方法
	var i int;
	for i=0;i<10;i++{
		fmt.Println(i);
	}
	//for 的while用法
	var j int = 0;
	for j<10{
		j++
		if j==5{
			continue;
		}
		if j==7 {
			break;
		}
		fmt.Println(">>",j);
	}
*/
/*
	fmt.Println(say("林俊杰","jj"));
	fmt.Println(see("林俊杰"));
	var x,y = "aa","bb";
	jh(&x,&y);
	fmt.Println(x,y);
*/
	var b [10] int;
	b[0]=1;
	fmt.Print(b[0]);
	var b0 [3]int{3,4,5}
	fmt.Print(b0[2]);
	var b1 [...] int{1,2,3,45,6,7,8}
	


}
//初始化方法
func init(){
	fmt.Println("init");
}
//定义方法
func say( x,y string) string {
	fmt.Println("调用：say方法！");
	return x+y;
}
func see(x string) (string,string) {
	fmt.Println("调用：see方法！");
	return "jj",x;
}
//交换：涉及地址知识
func jh(x,y *string){
	var z string=*x;
	*x=*y;
	*y=z;
}