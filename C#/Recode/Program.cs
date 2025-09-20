// 'recode' 키워드는 Tuple 형식과 동일한 구조에서 여러 데이터타입 값을 선언하고자 할 경우에 사용한다.
public record Point(int X, int Y, int Z);

public class Program
{
    static void Main()
    {
        // recode의 저장된 값을 불러오기 위해 새로운 Instance를 생성
        Point pt = new Point(5, 10, 3);

        // 'with'는 recode와 Tuple 같은 불변[Immutable] 데이터 구조를 다룰 때 활용
        var pt2 = pt with { X = 7, Y = 11 };
        Console.WriteLine($"Value update Point Result : {pt} and {pt2}");
    }
}