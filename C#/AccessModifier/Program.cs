namespace PrivateNS
{
    // Private은 생략되어 있지만 기본값으로 외부에서 접근 권한이 없다.
    class Top_Class
    {
        private class classB {}

        protected class ClassC { }

        internal class ClassD { }
    }

    // Private 키워드가 붙어 외부에서 접근 권한이 없다.
}


namespace PublicNS
{
    public class Public_Top_Class
    {
        
    }
}