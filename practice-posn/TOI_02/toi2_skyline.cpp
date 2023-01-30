#include <iostream>
using namespace std;
#define MAX 256

/*

2
1 11 5
2 6 7

*/
int tower[MAX];

int main()
{
    int n;
    cin>>n;
    for(int i=0;i<n;i++)
    {
        int start,hight,end;
        cin>>start>>hight>>end;
        for(int j=start;j<end;j++)
            tower[j]=max(tower[j], hight);
    }
    
    int temp=0;
    for(int i=0;i<MAX;i++)
    {
        if(temp==0&&tower[i]!=0){
            temp=tower[i];
            cout<<i<<" "<<tower[i]<<" ";
        }

        if(tower[i]!=temp){
            temp=tower[i];
            cout<<i<<" "<<tower[i]<<" ";
        }
    }
    
    return 0;
}