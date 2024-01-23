N=int(input())
Soma1=0
Soma2=0
while N!=0:
    for a in range(0,N):
        pontos=list(map(int,input().split()))
        if pontos[0]>pontos[1]:
            Soma1=Soma1+1
        elif pontos[1]>pontos[0]:
            Soma2=Soma2+1
    print(Soma1, Soma2)
    N=int(input())
    Soma1=0
    Soma2=0