#include <cmath>
#include <iostream>
using namespace std;

int main() {
  int raio;
  cin >> raio;

  printf("VOLUME = %.3f\n", (4.0 / 3.0) * 3.14159 * (std::pow(raio, 3)));
}