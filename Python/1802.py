matrix_livros = []
soma=[]

for i in range (0,5):
    livros = list(map(int,input().split()))
    del (livros[0])
    livros.sort(reverse=True)
    matrix_livros.append(livros)
k=int(input())

for a in matrix_livros[0]:
    for b in matrix_livros[1]:
        for c in matrix_livros[2]:
            for d in matrix_livros[3]:
                for e in matrix_livros[4]:
                    soma.append(a+b+c+d+e)
soma.sort(reverse=True)
print(sum(soma[0:k]))