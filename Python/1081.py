def visita (grafo, tempo, no, vertices, visitas):
    for i in range (0, vertices):
        variavel = "  "*tempo
        if (grafo[no][i] == 2) and (visitas[i] == 1):
            print(str(variavel) + str(no)+"-"+str(i)+" pathR(G,"+str(i)+")")
            grafo[no][i] = 1

            visitas[no] =0
            visitas[i] = 0
            visita(grafo,tempo+1,i,vertices,visitas)

        elif (visitas[i] == 0)  and (grafo[no][i]!=0):
            print(str(variavel) + str(no) + "-" + str(i))


def main ():
    n = int(input())
    tempo = 1
    caso = 1

    while (n!= 0):
        vertices,arestas = input().split()

        vertices = int(vertices)
        arestas = int(arestas)

        grafo =[]
        visitas =[0]*vertices

        for i in range (vertices):
            linha = []
            for j in range(vertices):
                linha.append(0)
            grafo.append(linha)

        for j in range (arestas):
            origem, destino = input().split()
            origem = int(origem)
            destino = int(destino)

            grafo[origem][destino]=2

            visitas[destino] =1
            visitas[origem] = 1

        print("Caso " + str(caso) + ":")
        visita(grafo, tempo, 0, vertices, visitas)
        print()
        for i in range (0,len(visitas)):
            if (visitas[i]== 1):
                visita(grafo,tempo,i,vertices,visitas)
                print()
        caso = caso +1
        n = n-1


main()