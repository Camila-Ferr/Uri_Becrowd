dicionario = {}
n = int(input())

for _ in range (0,n):
  lingua = input()
  chave = input()
  dicionario[lingua] = chave

lista = int(input())
for _ in range (0,lista):
  print(input())
  chave = input()
  print(dicionario[chave]+"\n")