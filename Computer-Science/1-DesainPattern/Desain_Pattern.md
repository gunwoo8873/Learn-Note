# Desain Pattern

## Desain Paatern이란?
디자인 패턴은 프로그램을 설계할 때 발생했던 문제점들을 객체 간의 상호 관계 등을 이용하여 해결할 수 있도록 하나의 `규약` 형태로 만들어 놓은 것을 의미한다

### 1. Singleton Pattern
---
하나의 `클래스[Class]`에 오직 하나의 `인스턴스[Instance]`만 가지는 패턴이며, 하나의 클래스를 기반으로 여러 갱의 개별적인 인스턴스를 만들 수 있지만, 그렇게 하지 않고 하나의 클래스를 기반으로 단 하나의 인스턴스를 만들어 이를 기반으로 로직을 만드는 데 쓰이고, 주로 Database 연결 모듈에 많이 사용한다.

### 1-2. Singleton Pattern 장단점
---
* 장점
  * 하나의 인스턴스를 생성하고 해당 인스턴스를 다른 Module들이 공유하여 사용하기 때문에 인스턴스 생성 비용이 감소한다.

* [단점](#1-3-signleton-pattern-단점)
  * 하나의 인스턴스를 생성하고 사용하는 Module들이 늘어날 수록 의존성이 상승한다.

### 1-3. Signleton Pattern 단점
---
`TDD[Test Driven Development]`에 대한 문제가 발생할 가능성이 크며, 주로 TDD를 할 땐 단위 테스트를 위한 방식이나, 서로 독립적인 테스트여야 하기 때문에 어떤 순서로든 실행이 가능해야 한다. 하지만 미리 생성된 하나의 Instance를 기반으로 구현하는 패턴이므로 각 테스트마다 `독립적인` 인스턴스를 만들기 어렵다. 또한 Signleton Pattern은 사용하기 쉽고 실용적이지만 모듈간의 결합을 강하게 만들 수 있다는 단점이 존재한다. 이경우 `의존성 주입[DI, Dependency Injection]`을 통해 모듈간의 완화로 해결할 수 있다.

### 1-3-2. Dependency Injection 장단점
* 장점
  * Module들을 쉽게 교체할 수 있는 구조화가 되어 테스트에 적합하며 쉽게 마이크레이션이 가능하다.
  * 구현할 때 추상적인 레이어를 넣고 이를 기반으로 구현체를 넣어 주기 때문에 애플리케이션 의존성 방향이 일관성을 띈다.
  * 

* 단점

### Singleton Pattern Code
* Go
  ```go
  type Obj struct {
	  a int
  }

  var (
    instance *Obj
    once     sync.Once
  )

  func GetInstance() *Obj {
    once.Do(func() {
      instance = &Obj{a: 1}
    })

    return instance
  }

  func main() {
    objA := GetInstance()
    objB := GetInstance()

    fmt.Println(objA == objB)
  }
  ```