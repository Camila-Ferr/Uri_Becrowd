import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;

class Main {
  public static void main(String[] args) {
    Scanner scanner = new Scanner(System.in);
    
    Set<Integer> figurinhas = new HashSet<>();
    Integer total, compradas;
    
    total = scanner.nextInt();
    compradas = scanner.nextInt();

    for (int i = 0; i<compradas;i++){
      figurinhas.add(scanner.nextInt());
    }
    System.out.println(total -  figurinhas.size());
    
  }
}