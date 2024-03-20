#include<stdio.h>
void xorswap(int*, int*);  //function prototype

void xorswap(int *p1, int *p2){
    *p1 = *p1^*p2;
    *p2 = *p2^*p1;
    *p1 = *p1^*p2;
}

int main(){
    int first, second;
    printf("First val: ");
    scanf("%d", &first);
    printf("Second value: ");
    scanf("%d", &second);
    printf("\n...swapping with xorswap()\n");
    xorswap(&first, &second);
    printf("\nfirst: %i\nsecond: %i\n", first, second);
}