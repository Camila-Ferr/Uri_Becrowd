#include <iostream>
using namespace std;

int main() {
  int testes,n1,n2;
  cin >> testes;
  for (int i = 0; i < testes; i++) {
    int soma = 0;
    cin >> n1 >> n2;
    int menor =n1>n2? n2:n1;
    int maior = n1>n2? n1:n2;

    for (int j=menor+1; j<maior; j++){
      soma += j%2? j: 0;
    }
    printf("%d\n", soma);
  }
}