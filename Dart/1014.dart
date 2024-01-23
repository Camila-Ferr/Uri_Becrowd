import 'dart:io';

void main() {

  double km = double.parse(stdin.readLineSync()!);
  double litros = double.parse(stdin.readLineSync()!);
  print("${(km/litros).toStringAsFixed(3)} km/l");
}