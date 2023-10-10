#include <string.h>
#include <stdio.h>
#include <stdlib.h>

char * remove_prefix(const char * str, int i) {
	// remove the first i of a str
	int l = strlen(str); 
	char * s = (char *) malloc( (l  - i + 1) * sizeof(char));
	int j=0;
	while (i < l) {
		s[j++] = str[i++];
	}
	s[l] = '\0';
	return s;
}
char * gcdOfStrings(char * str1, char * str2) {
	int l1 = strlen(str1);
	int l2 = strlen(str2);
	int cum_len = l1+l2;
	char * ab = (char * ) malloc((l1+l2+1) * sizeof(char)); ab[cum_len] = '\0';
	char * ba = (char * ) malloc((l1+l2+1) * sizeof(char)); ba[cum_len] = '\0';
	ab = strcat(ab, str1); ab = strcat(ab, str2); // this is dumb
	ba = strcat(ba, str2); ba = strcat(ba, str1);
	if (strcmp(ab,ba)) 	
		return "";
	if (l1 == l2) 
		return str1;
	if (l1 > l2)  
		return gcdOfStrings(remove_prefix(str1,l2),str2);
	return gcdOfStrings(str1, remove_prefix(str2,l1));
}

int main(int argc, char * argv[]) {
	if (argc == 3) {
		char * y = gcdOfStrings(argv[1],argv[2]);
		printf("%s\n",y);
		return 0;
	}
	char * a = "abc";
	printf("%s\n",remove_prefix(a,0));
	a = a+ sizeof(char);
	printf("%s\n",remove_prefix(a,0));
	printf("%s\n",remove_prefix(a,2));
	printf("%s\n",remove_prefix(a,1));
	return 0;
}
