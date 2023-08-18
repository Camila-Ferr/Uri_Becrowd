import 'dart:io';

void main() {

  double tempo = double.parse(stdin.readLineSync()!);
  double velocidade = double.parse(stdin.readLineSync()!);
  print("${((velocidade*tempo)/12).toStringAsFixed(3)}");
}