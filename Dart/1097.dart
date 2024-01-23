import 'dart:io';

void main() {
  int max = 5;
  for (int i = 1; i <= 9; i += 2) {
    max = max + 2;
    for (int j = max; j > max-3; j--) {
      print("I=$i J=$j");
    }
  }
}