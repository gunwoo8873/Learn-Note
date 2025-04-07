namespace STRING_NS {
  public class STRING_CLASS {
    public static string String() {
      // String data type size is EN[1 byte] and KO[2 byte]
      string str = "String DataType";
      return str;
    }
  }
}

// namespace INT_NS {
//   public class SIGN_INTEGER {
//     public static int Value() {
//     }
//   }
//   public class UNSIGNED_INTEGER {
//     public static uint Value() {
//     }
//   }
// }

// namespace FLOAT {}

namespace TYPE_CHANGE_NS {
  public class TYPE_CHANGE_CLASS {
    public static string TypeChange() {
      string str = "String DataType";
    }
  } 
}

class DataType_Compile_Class {
  static void Main(string[] args) {
    STRING_NS.STRING str = new STRING_NS.STRING();
    Console.WriteLine(str.String());
  }
}