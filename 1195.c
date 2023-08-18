#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct sNoA {
  int chave;
  struct sNoA *esq;
  struct sNoA *dir;
} TNoA;

TNoA *criaNo(int chave) {
  TNoA *no;
  no = (TNoA *)malloc(sizeof(TNoA));
  no->chave = chave;
  no->esq = NULL;
  no->dir = NULL;
  return no;
}
void insere(TNoA **no, int chave) {
  if (*no == NULL) {
    *no = criaNo(chave);
  } else {
    if (chave < (*no)->chave) {
      insere(&((*no)->esq), chave);
    } else {
      insere(&((*no)->dir), chave);
    }
  }
}

void prefixo(TNoA *no) {
  if (no != NULL) {

    printf(" %d", no->chave);
    prefixo(no->esq);
    prefixo(no->dir);
  }
}
void infixo(TNoA *no) {
  if (no != NULL) {
    infixo(no->esq);
    printf(" %d", no->chave);
    infixo(no->dir);
  }
}

void posfixo(TNoA *no) {
  if (no != NULL) {
    posfixo(no->esq);
    posfixo(no->dir);
    printf(" %d", no->chave);
  }
}
void print(TNoA *raiz) {
  printf("\nPre.:");
  prefixo(raiz);

  printf("\nIn..:");
  infixo(raiz);

  printf("\nPost:");
  posfixo(raiz);

  printf("\n\n");
}

int main(void) {
  int nos, valor, casos, i;
  i = 0;
  scanf("%d", &casos);

  while (i < casos) {
    printf("Case %d:", i + 1);
    scanf("%d", &nos);
    TNoA *raiz = NULL;

    for (int j = 0; j < nos; j++) {
      scanf("%d", &valor);
      insere(&raiz, valor);
    }
    print(raiz);
    i++;
  }

  return 0;
}