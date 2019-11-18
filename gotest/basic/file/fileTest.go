package main

func main() {

	/*err := os.Mkdir("test",os.ModeDir)
	fmt.Println(err)*/

	/*file , err := os.Create("test.txt")
	if err != nil{
		fmt.Println(err)
	}
	_,err1 := file.Write([]byte("开始操作文件"))
	if err1 != nil{
		fmt.Println(err1)
	}*/

	/*file ,err :=os.Open("test.txt")
	if err != nil{
		fmt.Println(err)
	}
	buffer := make([]byte ,1024)
	file.Read(buffer)
	fmt.Println(string(buffer))*/

	/*file ,err :=os.OpenFile("test.txt",os.O_APPEND,0777)
	if err != nil{
		fmt.Println(err)
	}
	buffer := make([]byte ,1024)
	file.Read(buffer)
	fmt.Println(string(buffer))*/

}
