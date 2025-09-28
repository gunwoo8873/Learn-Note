using System;
using System.Net;

class StandardMethods
{
    static void NormalReturnMethod()
    {
        //...
    }

    static int IntegerReturnMethod()
    {
        //...
    }
}


class AsyncMethods
{
    // 'string[]'는 매게 변수 타입
    static async Task AsyncTaskMethod(string[] args)
    {
        return await Task.FromResult(args);
    }

    static async Task<int> IntegerAsyncTaskMethod()
    {
        return await Task.FromResult<int>(0);
    }
}