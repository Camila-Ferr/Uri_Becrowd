#include <iostream>
using namespace std;

float calculaDeslocamento(int tempo, float vo) {
  // s = so + vot + at2/2
  return ((vo * tempo) + (vo * tempo));
}
int main() {
  int velocidade, tempo;

  while (cin >> velocidade >> tempo) {
    printf("%d\n", int(calculaDeslocamento(tempo, velocidade)));
  }
}