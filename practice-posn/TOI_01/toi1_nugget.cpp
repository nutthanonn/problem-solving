#include <iostream>
using namespace std;

int nugget[100];

int main()
{
    int n;
    cin>>n;
    nugget[6]=1;
    nugget[9]=1;
    nugget[20]=1;

    for(int i=12;i<=n;i++)
        if(nugget[i-6]||nugget[i-9]||nugget[i-20])
            nugget[i]=1;

    if(n<6)
    {
        cout<<"no";
        return 0;
    }

    for(int i=6;i<=n;i++)
        if(nugget[i])
            cout<<i<<endl;

    return 0;
}   