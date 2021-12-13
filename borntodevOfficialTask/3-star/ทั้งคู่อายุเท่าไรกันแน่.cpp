#include <bits/stdc++.h>
using namespace std;

int main()  {
    int N, M, Y;
    cin >> N >> M >> Y;
    int b = (Y*(M-1) + N) / (M-1);
    int a = N+b;
    cout << a << " " << b << endl;
    return 0;
}