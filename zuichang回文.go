package main

import "fmt"

func main() {
	var s string ="ababcdadcba"
	a:=longestPalindrome(s)
	fmt.Printf("%d",a)
}
func longestPalindrome(s string) int {
	length:=len(s)
	sum:=1
	max:=1
	n:=make(map[uint8]int)
	for i:=0;i<length;i++{
		for j:=0;j<length;j++{

			for  j-i-1>0&&(j-i-1)%2==1{
				sum=j-i+1
				_,ok:=n[s[i]]
				if ok{
					i++
					j--
				}
				if j==i&&sum>max{
					max=sum
				}

			}
			n[s[i]]=j+1
		}
	}
	return max
}