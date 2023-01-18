#include <iostream>
using namespace std;

int mdc(int maior, int menor) {
  int temp;
  while (menor != 0) {
    temp = menor;
    menor = maior % menor;
    maior = temp;
  }
  return temp;
}

int main() {
  int casos, fig1, fig2,maior,menor;
  cin >> casos;

  for (int i = 0; i < casos; i++) {
    cin >> fig1 >> fig2;

    if (fig1 < fig2){
      maior = fig2;
      menor = fig1;
    }
    else{
      maior = fig1;
      menor = fig2;
    }
    printf("%d\n", mdc(maior, menor));
  }
}