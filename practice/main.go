package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎来到澳门皇家赌场")
	fmt.Println("1.添加女郎")
	fmt.Println("2.编辑学生信息")
	fmt.Println("3.展示所有女郎信息")
	fmt.Println("4.退出系统")
}

func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入女郎信息")
	fmt.Println("请输入女郎的编号：")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入女郎的姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入女郎的班级:")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStudentMgr()
	for {
		// 1-打印系统菜单
		showMenu()
		// 2-等待用户选择执行选项
		var input int
		fmt.Print("请输入你要操作的序号")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是", input)
		// 3-执行用户选择的动作
		switch input {
		case 1:
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			sm.showStudent()
		case 4:
			os.Exit(0)
		}
	}
}
