import java.util.Scanner;

class Main {

  public static void main(String[] args) {
    Scanner teclado = new Scanner(System.in);
    int num = teclado.nextInt();
    int horas = teclado.nextInt();
    double salario = teclado.nextDouble();
    System.out.println("NUMBER = " + num);
    System.out.printf("SALARY = U$ %.2f\n", horas * salario);
  }
}