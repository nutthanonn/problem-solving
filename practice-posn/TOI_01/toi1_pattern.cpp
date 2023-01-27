#include <iostream>
using namespace std;


int main()
{
    int n;
    cin>>n;
    int p=70;
    string col[n];
    for(int j=0;j<n;j++)
        for(int i=0;i<p;i++)
            col[j]+="o";
    int max_use=0;
    
    for(int i=0;i<n;i++)
    {
        int p,q,r;
        cin>>p>>q>>r;
        max_use=max(max_use,p);
        for(int j=0;j<r;j++)
        {
            col[p-1][q+j-1]='x';
        }
    }

    for(int i=0;i<max_use;i++)
        cout<<col[i]<<endl;
    return 0;
}