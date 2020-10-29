#include "ft_memset.c"

// for testing
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main()
{
	int data_length = 4;
	void *data = malloc(sizeof(char) * data_length);
	void *data2 = malloc(sizeof(char) * data_length);
	unsigned char c = '4';

	void *res = memset(data, c, data_length);
	
	int i;
	for (i = 0; i < data_length; i ++) {
		printf("res: %c\n", ((char *)res)[i]);
	}
	for (i = 0; i < data_length; i ++) {
		printf("data: %c\n", ((char *)data)[i]);
	}

	printf("----\n");

	void *res2 = tf_memset(data2, c, data_length);
	for (i = 0; i < data_length; i++) {
		printf("res2: %c\n", ((char *)res2)[i]);
	}
	for (i = 0; i < data_length; i++) {
		printf("data2: %c\n", ((char *)data2)[i]);
	}
}
