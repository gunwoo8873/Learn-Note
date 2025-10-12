using System;

namespace ReservedWordNS
{
    public class Container(int v)
    {
        private int _i = v;
    }

    class Program
    {
        static void Main()
        {
            // string = Stack area
            // str = Heap area
            string str = "This is String type";
            Console.WriteLine(str);

            int i = 100;
            i = 150; // Update and Reference Type
            Console.WriteLine(i);
        }
    }
}