using System;

namespace ArrayNS
{
    class Program
    {
        static void Main()
        {
            //// 단일차원 배열 [None Index]
            int[] NomalArr = new int[10];

            //// 2차원 배열 = [X(Column), Y(Row)]
            int[,] MultiInterArrA = new int[2, 5]
            {
                {00, 01, 02, 03, 04},
                {05, 06, 07, 08, 09},
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
            int[][] ImmutableArr = new int[4][];
            ImmmutableArr[0] = new int[4];
            ImmmutableArr[1] = new int[5];
            ImmmutableArr[2] = new int[2];
            ImmmutableArr[3] = new int[3];
        }
    }
}