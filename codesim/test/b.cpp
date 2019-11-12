#include<stdio.h>
int main(){
	int num=0;
	int a[1000],b[1000];
	while(scanf("%d %d",&a[num],&b[num])!=EOF){
		num++;
	}
	if(num==1&&b[0]==0){
		printf("0 0");
	}
	else {
    for(int i=0;i<num;i++){
		a[i]=a[i]*b[i];
		b[i]=b[i]-1;
		if(i==0){
		printf("%d %d",a[i],b[i]);
	   }	   
	   else{
	   	if(b[i]!=-1) printf(" %d %d",a[i],b[i]); 
	   }
	}	
  }
  return 0;
}
