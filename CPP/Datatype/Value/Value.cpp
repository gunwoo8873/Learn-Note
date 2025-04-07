#include <iostream>

using namespace std;

namespace Integer {
    // Short type integer value
    int int8_fn() {
        int integerType = 100000;
        return printf("%c", integerType);
    }

    int int16_fn() {
    }

    int int32_fn() {
    }

    int int64_fn() {
    }
}

namespace Uinteger {
}

namespace Float {
}

namespace Char {
    int Char_fn() {
        char str_CharType = 'a';
        return printf("%c", str_CharType);
    }
}

int main()
{
    cout << Char::Char_fn() << endl;
}