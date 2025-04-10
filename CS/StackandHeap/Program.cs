namespace StackNS {
  public class StackClass {
    int stackValue = 10; // Save in the stack memory area
    
    public static void PrintStackValue() {
      Console.WriteLine("Stack Value: " + stackValue);
    }
  }
}

namespace HeapNS {
  public class HeapClass {
    int heapValue = new int(); // Save in the heap memory area

    public static void PrintHeapValue() {
      Console.WriteLine("Heap Value: " + heapValue);
    }
  }
}

class Compile {
  static void Main(string[] args) {
    StackNS.StackClass.PrintStackValue();
    HeapNS.HeapClass.PrintHeapValue();
  }
}