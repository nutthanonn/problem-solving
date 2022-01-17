#include <stdio.h>

int main() {
    float THAI, MATH, ENG, SCI, SPORT;
    scanf("%f%f%f%f%f", &THAI, &MATH, &ENG, &SCI, &SPORT);
    printf("THAI = %.1f\n", THAI);
    printf("MATH = %.1f\n", MATH);
    printf("ENGLISH = %.1f\n", ENG);
    printf("SCIENCE = %.1f\n", SCI);
    printf("SPORT = %.1f\n", SPORT);
    printf("---\n");
    printf("GPA = %.1f", (THAI+MATH+ENG+SCI+SPORT)/5);
    return 0;
}