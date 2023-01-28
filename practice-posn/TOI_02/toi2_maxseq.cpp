#include <iostream>
#include <vector>
using namespace std;

// kadane's algorithm
int main()
{
    int n;
    cin>>n;
    int seq[n];
    for(int i=0;i<n;i++) cin>>seq[i];

    double max_so_far=0,max_ending_here=0;
    int max_start_index=0,startIndex=0,max_end_index=-1;

    for(int i=0;i<n;i++)
    {
        if(0 > max_ending_here+seq[i]) 
        {
            startIndex = i+1;
            max_ending_here = 0;
        }
        else 
        {
            max_ending_here += seq[i];
        }

        if(max_ending_here>max_so_far) 
        {
            max_so_far=max_ending_here;
            max_start_index=startIndex;
            max_end_index=i;
        }
    }

    vector <int> max_sub;
    int col=0;

    if(max_start_index <= max_end_index)
    {
        for(int i=max_start_index;i<=max_end_index;i++)
        {
            col+=seq[i];
            max_sub.push_back(seq[i]);
        }
    }

    if(col<=0)
    {
        cout<<"Empty sequence";
        return 0;
    }

    for(auto x : max_sub)
    {
        cout<< x << " ";
    }
    cout<<endl<<col;


    return 0;
}