import java.util.Scanner;

public class Matrix{
    private int n;
    private double [][] matrix;
    private double determinante;
    private boolean possibilidade = true;

    //Construtor que determina a dimensão da matriz, a matriz e seu determinante.
    Matrix(int n, double[][] matrix){
        this.n=n;
        this.matrix= matrix;
        this.determinante = calculaDeterminante();

    }
    //Molda a matriz seguindo o método de Gauss
    public void moldaMatrix() {

        double coeficiente;

        for (int k =0; k<n-1; k++){
            
            for (int i=k+1; i<n; i++){
                
                if (this.matrix[k][k]==0){
                    troca(k);
                    if (this.matrix[k][k] == 0){
                        this.possibilidade=false;
                        return;
                    }
                }

                coeficiente= this.matrix[i][k] / this.matrix[k][k];
                this.matrix[i][k] = 0;
                
                for (int j=k+1; j<n; j++){
                    this.matrix[i][j] -= coeficiente * this.matrix[k][j];

                }
            } 
        }
    }
    //Calcula o determinate segundo o método de Gauss
    public double calculaDeterminante() {
        
        moldaMatrix();

        double determinante =1;
        
        if (this.possibilidade){
            
            System.out.println();
            System.out.println("Matriz triangularizada :");
            printarMatrix();
        
            for (int i=0; i<n; i++){
                determinante*= this.matrix[i][i];

            }
        }
        else{
            determinante =0;
        }
        return determinante;
    }
    
    // Apresenta a matriz linha por linha
    void printarMatrix(){
        for (int i=0; i<n; i++){
            System.out.println();
            for (int j=0; j<n; j++){
                System.out.print(this.matrix[i][j] +"   ");
           }
        }
        System.out.println("\n");
    }

    //Troca as linhas caso o denominador coeficiente do método moldaMatrix seja zero.
    void troca(int k){        
        double [] aux = new double [this.n];
        
        for (int j=0; j<this.n; j++){
            matrix[k][j] *= -1;
        }
        
        for (int i= k+1; i< this.n; i++) {
            if(this.matrix[i][k] != 0){

                aux= this.matrix[i];
                
                this.matrix[i] = this.matrix[k];
                this.matrix[k] = aux;
                break;
            }
        }
    }
    //gets para chamar na função principal
    public double getDeterminante() {
        return determinante;
    }
    public boolean getPossibilidade(){
        return possibilidade;
    }

    public static void main(String[] args){

        Scanner teclado=new Scanner(System.in);

        int n;

        //Recebe a matriz
        System.out.println("Digite uma dimensão para a matriz:");
        n= teclado.nextInt();
        double [][] matriz = new double [n][n];

        System.out.println("Digite as linhas da matriz:");
        
        for (int a=0; a<n; a++){
            System.out.println("Linha " +(a+1) +":");
            
            for (int b=0; b<n; b++){
                matriz[a][b]=teclado.nextDouble();
            }
        }

        //Cria um objeto da classe Matrix
        Matrix matrix= new Matrix(n,matriz);
        
        //Verifica se é possível realizar por Gauss 
        if (matrix.getPossibilidade()==true){
            System.out.println("O determinante da matriz é : " +matrix.getDeterminante());            
        }
        //Caso contrário, determinante =0
        else{
            System.out.println("O determinante da matriz é 0."); 
        }

    }
}