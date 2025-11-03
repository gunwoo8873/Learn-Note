using System;

namespace Variable {
    class Program {
        static void Integer(string[] args) {
            int Num = 10; // int of value type is integer [-1, 0, ...etc]
            Console.WriteLine(Num);
        }

        static void String(string[] args) {
            string Str = "String"; // string of reference type is a sequence of characters
            Console.WriteLine(Str);

            char Char = 'C';
            Console.WriteLine(Char);
        }

        static void Boolean(string[] args) {
            bool BoolTrue = true, BoolFalse = false; // bool of value type is true or false
            Console.WriteLine(BoolTrue);
            Console.WriteLine(BoolFalse);
        }

        static void Object(string[] args) {
            object Obj = "Object", Obj2 = 123; // object of reference type is the base type of all data types
            Console.WriteLine(Obj);
            Console.WriteLine(Obj2);
        }

        static void Initailization(string[] args) {
            int Num; // Declaration
            Num = 20; // Initialization
            Console.WriteLine(Num);
        }

        static void Null(string[] args) {
            string Null = null; // null means no value
            Console.WriteLine(Null);
        }

        static void MultiVariables(string[] args) {
            int a = 1, b = 3, c = 5;
            a = b = c;
            Console.WriteLine("{0} {1} {2}", a, b, c);
        }

        static void Main(string[] args) {
            Integer(args);
            String(args);
            Boolean(args);
            Object(args);
            Initailization(args);
            Null(args);
            MultiVariables(args);
        }
    }
}