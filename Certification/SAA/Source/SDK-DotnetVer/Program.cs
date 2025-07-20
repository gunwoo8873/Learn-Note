// Dotnet use librarys
using System;
using System.Threading.Tasks;

// Amazon API librarys
using Amazon.Runtime;

namespace SDK_S3Bucket
{
    public class S3_List
    {
        public static async Task(string[] args)
        {
            var ssoCreds = LoadSSOCredentials("default");
        }
    }
}