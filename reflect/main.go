package main

import "reflect"

// 定义一个Enum类型
// type Enum int

// const (
//     Zero Enum = 0
// )
//
// type Student struct {
//     Name string
//     Age  int
// }
//
// func main()  {
//     // var stu Student
//     // typeofSud := reflect.TypeOf(stu)
//     // fmt.Println(typeofSud.Name(),typeofSud.Kind())
//     // // 获取Zero常量的反射类型对象
//     // typeOfZero := reflect.TypeOf(Zero)
//     //
//     // // 显 示反射类型对象的名称和种类
//     // fmt.Println(typeOfZero.Name(), typeOfZero.Kind())
//     // 定义一个Student类型的指针变量
//     var stu = &Student{Name:"kitty", Age: 20}
//
//     // 获取结构体实例的反射类型对象
//     typeOfStu := reflect.TypeOf(stu)
//     // 取类型的元素
//     typeOfStu = typeOfStu.Elem()
//
//     // 显示反射类型对象的名称和种类
//     fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfStu.Name(), typeOfStu.Kind())
// }

// func main()  {
//     
//     // 声明一个空结构体
//     type cat struct {
//         Name string
//         
//         // 带有结构体tag的字段
//         Type int `json:"type" id:"100"`
//     }
//     
//     
//     // 创建cat的实例
//     ins := cat{Name: "mimi", Type: 1}
//     // 获取结构体实例的反射类型对象
//     typeOfCat := reflect.TypeOf(ins)
//     // 遍历结构体所有成员
//     for i := 0; i < typeOfCat.NumField(); i++ {
//         
//         // 获取每个成员的结构体字段类型
//         fieldType := typeOfCat.Field(i)
//         
//         // 输出成员名和tag
//         fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
//     }
//     // 通过字段名, 找到字段类型信息
//     if catType, ok := typeOfCat.FieldByName("Type"); ok {
//         
//         // 从tag中取出需要的tag
//         fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
//     }
// }

type MyMath struct {
    Pi float64
}

// 普通函数
func (myMath MyMath) Sum(a, b int) int {
    return a + b
}

func (myMath MyMath) Dec(a, b int) int {
    return a - b
}
func main()  {
    var myMath =MyMath{pi:3.14159}
    rV := reflect.ValueOf(myMath)
    
    // 获取方法数量
    MymathNumfunc := rV.NumMethod()
}