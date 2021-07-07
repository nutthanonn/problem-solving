#include <iostream>
using namespace std;

int main() {
    int t[] = {};
    int a;
    for(int b=0;b<5;b++){
        scanf("%d", &a);
        t[b] = a;
    }
    for(int i=0;i<5;i++){
        printf("%d\n", t[i]);
    }
    
    return 0;
}