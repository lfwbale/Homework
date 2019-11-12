#include<stdio.h>
const int maxn=10000;
int main(){
    int equ[maxn];
    int num=0;
    while(scanf("%d",&equ[num])!=EOF){  //把多项式的系数和指数全部输入数组
        num++;
    }
    //如果是0多项式
    if(num==2&&equ[1]==0){
        printf("0 0");
    }
    else{
    int count=num/2;//一个系数对应一个指数
        int a[count],b[count];//a是系数 b是指数；
        for(int i=0;i<count;i++){
            a[i]=equ[i*2];
            b[i]=equ[i*2+1];
        }
        for(int i=0;i<num;i++){
            if(i%2==0){//是系数
                equ[i]=a[i/2]*b[i/2];
            }
            else{
                equ[i]=b[i/2]-1;//
            }
        
        }
        //系数为0则不输出
                printf("%d %d",equ[0],equ[1]);//第一个要输出的数
       
        for(int j=2;j<num;j++){
            if(j%2==0&&equ[j+1]!=-1){
                printf(" %d %d",equ[j],equ[j+1]);
            }
        }
}
    return 0;
}

