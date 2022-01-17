//พีระมิด(ระดับ 2)
#include <stdio.h>

int main() {
    int number;
    scanf("%d", &number);
    int round = number - 1; 
    for(int t = 0; t < number; t++) {

        //space
        for(int i = 0; i < round; i++) {
            printf(" ");
        }
        

        //star
        for(int i = 0; i < t*2+1; i++) {
            printf("*");
        }
        round--;
        printf("\n");
    }
    
    return 0;
}