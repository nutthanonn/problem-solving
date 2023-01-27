#include <iostream>
using namespace std;


int main()
{
    int l,n;
    cin>>l>>n;
    string word[n];
    for(int i=0;i<n;i++)
        cin>>word[i];

    for(int i=0;i<n;i++)
    {
        int chain = 0;
        for(int j=0;j<l;j++)
        {
            if (word[i][j] != word[i+1][j])
            {
                chain++;
            }
        }   

        if (chain > 2)
        {
            cout << word[i];
            return 0;
        }
    }

    return 0;
}