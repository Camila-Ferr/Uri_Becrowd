N = int(input())
test=1
while N!=0:
    participantes = list(map(int, input().split()))
    Numero = 0
    print('Teste',test)
    for a in range (0,N):
        Numero += 1
        if Numero == participantes[a]:
            print(Numero)
            break
    print()
    N=int(input())
    test=test+1