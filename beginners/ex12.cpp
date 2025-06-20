#include <iostream>
using namespace std;

int main() {
  string S;
  cin >> S;

  // ここにプログラムを追記
  int count = 1;

  for(int i = 0; i < S.size(); i++) {
    if (S.at(i) == '+') {
      count++;
    }
    else if (S.at(i) == '-') {
      count--;
    }
  }
  cout << count << endl;
}

