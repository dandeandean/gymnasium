#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char * mergeAlternately(char * word1, char * word2) {
	int len1, len2, i,j;
	len1 = strlen(word1); len2 = strlen(word2);
	char * out = malloc(
		(len1 + len2 + 1)* sizeof(char)
		);
	i = 0, j = 0;
	while (i < len1 || i < len2)
	{
		if (i < len1) out[j++] = word1[i];
		if (i < len2) out[j++] = word2[i];
		i ++;
	}
	out[j] = '\0';
	return out;
}

int main(int argc, char * argv[]) {
    printf("MAIN.... \n");
    char* m = mergeAlternately(argv[1],argv[2]);
    printf("%s\n",m);
    int i = 0;
    char * c = "helloworld";
    while (i < 10) {
	    printf("%c",c[i++]);
    }
    return 0;
}

