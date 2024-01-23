#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct sNoA {
    int chave;
    struct sNoA *esq;
    struct sNoA *dir;
} TNoA;

TNoA *criaNo(int chave){
  TNoA*no;
  no = (TNoA *) malloc(sizeof(TNoA));
  no->chave = chave;
  no->esq = NULL;
  no->dir = NULL;
  return no;
}
void insere(TNoA*no, int* nivel){
  if (*nivel == 1){
    return;
  }
  else{
    no->esq = criaNo(0);
    no->dir = criaNo(0);
    
    *nivel = *nivel-1;
    insere(no->esq, nivel);
    insere(no->dir, nivel);
    *nivel = *nivel+1;
  }
}

void competicao(TNoA*no, int *chave, int jog1, int jog2){
  
  if(no->esq == NULL){
    *chave = *chave +1;
    no->chave = *chave;
    return;
  }
  else{
  
    competicao(no->esq,chave,jog1,jog2);
    competicao(no->dir,chave,jog1,jog2);
    
    if ((((no->esq)->chave) == jog1) || (((no->dir)->chave) == jog1)){
      no->chave = jog1;
      
    }
      
    else if ((((no->esq)->chave) == jog2) ||(((no->dir)->chave) == jog2)){
      no->chave = jog2;
    }
    else{
      no->chave = no->esq->chave;
    }
  }
}

void percorre(TNoA*no, int* aparicao1, int* aparicao2, int jog1, int jog2 ){
  
  if (((no->chave) == jog1) || ((no->chave) == jog1)){
      *aparicao1 +=1;    
  }
  else if (((no->chave) == jog2) ||((no->chave) == jog2)){
      *aparicao2 +=1;
  }
  
  if (no->esq== NULL){
    return;
  }
  else{
    percorre(no->esq, aparicao1, aparicao2, jog1, jog2 );
    percorre(no->dir, aparicao1, aparicao2, jog1, jog2 );
  }

    
}


void imprime(TNoA *nodo, int tab) {
    for (int i = 0; i < tab; i++) {
        printf("-");
    }
    if (nodo != NULL) {
        printf("%d\n", nodo->chave);
        imprime(nodo->esq, tab + 2);
        printf("\n");
        imprime(nodo->dir, tab + 2);
    } else printf("vazio");
}

int criaArvore(int jog1, int jog2){
  
  int chave = 0;
  int aparicao1 = 0;
  int aparicao2 = 0;
  int nivel = 5;
  
  TNoA*raiz = criaNo(0);
  
  insere(raiz,&nivel);
  competicao(raiz, &chave,jog1,jog2);
  percorre(raiz, &aparicao1, &aparicao2, jog1, jog2 );

  if (aparicao1> aparicao2){
    return (aparicao1 - aparicao2);
  }
  
  return (aparicao2 - aparicao1);
}

int main (void){
  int jog1,jog2,jogo;
  
  scanf("%d",&jog1);
  scanf("%d",&jog2);
  jogo = criaArvore(jog1,jog2);

  if (jogo == 4){
    printf("oitavas\n");
  }
  else if (jogo == 3){
    printf("quartas\n");
  }
   else if (jogo == 2){
    printf("semifinal\n");
  }
  else{
    printf("final\n");
  }
}
