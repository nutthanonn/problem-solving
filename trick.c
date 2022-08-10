#include <stdio.h>

int main()
{
    char name[1024];
    scanf("%[^\n]s", name);
    printf("Hello! %s", name);
    return 0;
}
