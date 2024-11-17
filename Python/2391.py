entradas = int(input())
sequencia = list(map(int,input().split()))
numero_progressao = 1

if entradas == 3:
    razao = sequencia[1] - sequencia[0]
    razao_prox = sequencia[2] - sequencia[1]
    if razao_prox!=razao:
        numero_progressao +=  1

elif entradas !=1 :
    razao= sequencia[1] - sequencia[0]
    for a in range (0,len(sequencia)-2):
        razao_prox= sequencia[a + 1] - sequencia[a]

        if razao_prox!=razao:
            numero_progressao +=  1
            razao= sequencia[a + 2] - sequencia[a + 1]

    if razao!=sequencia[entradas - 1]-sequencia[entradas - 2]:
        numero_progressao +=  1

print(numero_progressao)