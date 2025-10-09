using System;
using DotNetEnv;

class Config
{
    private readonly string configpath;
    private readonly string tokenvalue;

    // public Config(string configpath, string tokenvalue)
    // {
    //     this.configpath = configpath;
    //     this.tokenvalue = tokenvalue;
    // }

    private static string GetConfigPath(string configpath)
    {
        return this.configpath;
    }

    private static string GetTokenValue(string tokenvalue)
    {
        return this.tokenvalue;
    }
    
    static void Main()
    {
        Config config = new Config(".env", "TOKENTEST");
        var ConfigPath = GetConfigPath.configpath;
        var TokenValue = GetTokenValue.tokenvalue;

        Console.WriteLine("Path File :" + ConfigPath);
        Console.WriteLine("Token Value :" + TokenValue);
    }
}