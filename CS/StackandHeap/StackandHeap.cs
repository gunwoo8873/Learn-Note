namespace BoxingNS {
  public class BoxingCS {
    public void BoxingFN() {
      int value = 10;
      object value_box = value;

      Console.WriteLine("Boxing: " + value_box);
    }
  }
}
// Boxing은 값 형식의 타입을 Object 라는 참조 형식으로 형변환을 시도한다. 사용 목적이 여러 방식으로 존재하며 보통 파라미터로 전달하거나,
// 기존에 저장된 스택영역에서 Heap영역에 값형식을 저장한다. 

namespace StackNS {
  public class StackCS {
    public static void StackFN() {
      int a = 10; // Value type to 4 bytes integer type
      int b;  // Reference type to 8 bytes reference type
      b = 20;

      Console.WriteLine("Value Type: " + a);
      Console.WriteLine("Reference Type: " + b);
    }
  }
}
// Stack Memory는 LIFO[Last In and First Out] 구조로 동작하며, 함수 호출 시 스택 프레임이 생성된다.
// Stack에 저장되는 변수의 메모리 주소는 스택 프레임의 위치에 따라 결정된다.
// 같은 프로그램을 N번 실행을 하더라도 실제 스택 메모리 주소는 일관되지 않을 수 있으므로,
// 이는 운영 체제와 런타임 환경에 대한 스택 메모리의 시작 위치를 다르게 설정할 수 있다.

namespace HeapNS {
  public class HeapCS {
    int[] heapArr = new int[5];
  }
}
// Heap Memory는 동적으로 할당되어, 메모리 주소는 가용 메모리 블록에 따라 결정된다.
// Heap Memory 주소는 가비지 컬렉터[GC, Garbage Collector]에 의해 관리되며, 프로그램 실행 중 메모리 할당/해제 상태에 따라 달라진다.
// Stack Memory와 같은 저장된 객체의 메모리 주소도 일관되지 않을 수 있다. 그렇기에 직접 관리할 수 없어 메모리 누수[Memory Leak]가 발생할 수 있다.

class Compiler {
  public static void Main(string[] args) {
    BoxingNS.BoxingCS.BoxingFN();
  }
}