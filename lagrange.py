from sympy import *

class polinomioLagrange ():

    def __init__(self, pontos):
        self.pontos = pontos
        self.indice = []
        self.monta_polinomio()

    def monta_polinomio(self):
        init_printing()

        #define uma variável
        x = var('x')
        X = []
        Y=[]

        #separa as cpprdemadas
        for i in range (0,len(self.pontos),2):
            X.append(self.pontos[i])
            Y.append(self.pontos[i+1])


        for i in range (0,len(X)):
            numerador = 1
            denominador = 1

            #Coloca a fórmula em prática

            for j in range (0, len(X)):
                #impede o denominador de ser 0
                if (j != i):
                    numerador *= (x - X[j])
                    denominador *= X[i] - X[j]
            L = (numerador/denominador)
            self.indice.append(L)

    def print (self): #printa as equações que formam o polinômio
        print()
        print('Polinômio = ', end=" ")

        for a in range (0,len(self.indice)):
            if (a == len(self.indice)-1):
                print(self.indice[a])
            else:
                print(self.indice[a], '+', end="")
        print("\n")


pol = polinomioLagrange([-1,4,0,1,2,-1])
pol.print()