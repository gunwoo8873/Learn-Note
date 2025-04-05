#include <iostream>
using namespace std;

// Outside module use do import keyword
// SourceFile in the create module interface [.ixx] file name
import ModuleA;

int main() {
    // The using imported module in the module namespace to name
    cout << "Result of f() : " << ModuleA_1_NS::fn() << endl;

    return 0;
}
