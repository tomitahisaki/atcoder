#include <iostream>
using namespace std;

int main() {
  int A, B;
  cin >> A >> B;

  // ここにプログラムを追記
  int i = 0;
  cout << "A:";
  while (i < A) {
    cout << "]";
    i++;
  }
  cout << endl;

  int j = 0;
  cout << "B:";
  while (j < B) {
    cout << "]";
    j++;
  }
  cout << endl;
}

