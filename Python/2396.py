carros,voltas = map(int, input().split())
colocacoes = [0,0,0]
tempo_max=0
tempo_medio=0
tempo_min =0

for x in range (0, carros):
    tempoVoltas=list(map(int,input().split()))
    tempoTotal = sum(tempoVoltas)
    if tempoTotal < tempo_max or tempo_max == 0:
        tempo_min = tempo_medio
        tempo_medio = tempo_max
        tempo_max = tempoTotal

        colocacoes[2] = colocacoes[1]
        colocacoes[1] = colocacoes[0]
        colocacoes[0] = x+1

    elif tempoTotal < tempo_medio or tempo_medio ==0:
        tempo_min = tempo_medio
        tempo_medio = tempoTotal

        colocacoes[2] = colocacoes[1]
        colocacoes[1] = x + 1

    elif tempoTotal < tempo_min or tempo_min ==0:
        tempo_min = tempoTotal
        colocacoes[2] = x + 1

for i in range(0,3):
    print(colocacoes[i])

