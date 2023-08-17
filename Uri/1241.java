import java.util.Scanner;

class Comp {
    private String string1;
    private String string2;
    private boolean compare;

    public Comp(String string1, String string2) {
        this.string1 = string1;
        this.string2 = string2;
        this.compare = verificaString();
    }

    public boolean verificaString() {
        if (string2.length() > string1.length()) {
            return false;
        } else {
            int startIndex = string1.length() - string2.length();
            String substring = string1.substring(startIndex);
            return substring.equals(string2);
        }
    }
  public boolean getCompare(){
    return this.compare;
  }
}

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        scanner.nextLine();

        for (int i = 0; i < n; i++) {
            String[] input = scanner.nextLine().split(" ");
            String A = input[0];
            String B = input[1];

            Comp comp = new Comp(A, B);
            if (comp.getCompare()) {
                System.out.println("encaixa");
            } else {
                System.out.println("nao encaixa");
            }
        }

        scanner.close();
    }
}
