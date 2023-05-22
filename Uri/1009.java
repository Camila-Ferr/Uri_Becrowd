import java.util.Scanner;

class Main {

  public static void main(String[] args) {
    Scanner teclado = new Scanner(System.in);
    teclado.next();
    double fixo = teclado.nextDouble();
    double vendas = teclado.nextDouble();
    System.out.printf("TOTAL = R$ %.2f\n", (fixo + (vendas * 0.15)));
  }
}