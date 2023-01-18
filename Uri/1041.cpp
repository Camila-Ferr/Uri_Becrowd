#include <iostream>
using namespace std;

bool classificaEixo(double eixo) {
  if (eixo > 0) {
    return true;
  }
  return false;
}

int verificaEixo(double eixox, double eixoy) {
  int resultado = 0;
  if (eixox == 0) {
    resultado += 1;
  }
  if (eixoy == 0) {
    resultado += 2;
  }
  return resultado;
}

int main() {
  double eixox, eixoy;
  int resultado;

  cin >> eixox >> eixoy;
  resultado = verificaEixo(eixox, eixoy);

  switch (resultado) {
  case 1:
    cout << "Eixo Y" << endl;
    break;

  case 2:
    cout << "Eixo X" << endl;
    break;

  case 3:
    cout << "Origem" << endl;
    break;

  default:
    if (classificaEixo(eixox) && classificaEixo(eixoy)) {
      cout << "Q1" << endl;
    } else if (!classificaEixo(eixox) && classificaEixo(eixoy)) {
      cout << "Q2" << endl;
    } else if (classificaEixo(eixox) && !classificaEixo(eixoy)) {
      cout << "Q4" << endl;
    } else {
      cout << "Q3" << endl;
    }
    break;
  }
}