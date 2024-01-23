n = int(input())
for i in range (0,n):
    soma = 0
    quadrados = int(input())
    for j in range(1,quadrados+1):
        soma += 2**(j-1)

    print("%d kg" %(int((soma/12)/1000)))