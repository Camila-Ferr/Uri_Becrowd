#include <iostream>
using namespace std;

int main() {
  int casos, dias;
  float comida;
  cin >> casos;

  for (int i = 0; i < casos; i++) {
    cin >> comida;
    dias = 0;

    while (comida > 1) {
      dias += 1;
      comida = comida / 2;
    }
    printf("%d dias\n", dias);
  }
}