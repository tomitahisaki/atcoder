#include <iostream>
using namespace std;

int main() {
  int A, B;
  string op;
  cin >> A >> op >> B;

  if (op == "+") {
    cout << A + B << endl;
  }

  if (op == "-") {
    cout << A - B << endl;
  }

  if (op == "*") {
    cout << A * B << endl;
  }

  if (op == "/") {
    if (B == 0) {
      cout << "error" << endl;
    } else {
      cout << A / B << endl;
    }
  }

  if (op == "?" || op == "=" || op == "!") {
    cout << "error" << endl;
  }

  return 0;
}

