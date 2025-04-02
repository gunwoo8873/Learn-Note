package main

import "fmt"

//////////////////////////////////////////////////////
// Interface
//////////////////////////////////////////////////////
// 1. 객체 간 상호작용을 정의
// 2. 덕 타이핑을 사용하면 서비스 사용자 중심의 코딩 가능
// 3. 타입 변한해 구체화된 타입으로 변환
//////////////////////////////////////////////////////

// Interface is define method to bind object type
type Interface interface {
	Method() string
}

// Create struct value data type
type Structer struct {
	name string
}

// Method is reference Structer type
func (s Structer) Method() string {
	return fmt.Sprintf("name: %s", s.name)
}

func main() {
	value := Structer{"golang"} // value is Structer type

	var structer Interface // structer is Interface change to type
	structer = value       // the value is assigned to structer

	fmt.Println(structer.Method()) // Structer reference method call
}
