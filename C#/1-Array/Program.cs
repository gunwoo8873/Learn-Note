using System;

namespace ArrayNS
{
    class Program
    {
        static void Main()
        {
            //// 단일차원 배열
            int[] NoneIndexArr = new int[25];

            //// 2차원 배열 = [X(Column), Y(Row)]
            int[,] MultiInterArrA = new int[2, 5]
            {
                {00, 01, 02, 03, 04},
                {05, 06, 07, 08, 09},

                // {00, 01, 02, 03, 04, 05, 06, 07, 08, 09},
                // {10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
                // {20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
                // {30, 31, 32, 33, 34, 35, 36, 37, 38, 39},
                // {40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
                // {50, 51, 52, 53, 54, 55, 56, 57, 58, 59}
            };

            //// 3차원 배열 = [X, Y(Column), Z(Row)]
            int[,,] MultIntegerArrB = new int[2, 3, 4]
            {
                {
                    {00, 01, 02, 03},
                    {05, 06, 07, 08},
                    {10, 11, 12, 13},
                },
                {
                    {14, 15, 16, 17},
                    {18, 19, 20, 21},
                    {22, 23, 24, 25},
                },
            };

            //// 가변 배열 = [Column <Mutable>][Row]
            int[][] ImmutableArr = new int[10][];
            ImmmutableArr[1] = new int[15]; // Row 15+
            Console.WriteLine(ImmutableArr[1]);

            string Text = "This is Text";

            char c1 = Text[1];
            Console.WriteLine(c1);
        }
    }
}