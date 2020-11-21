#include "libft.h"

#include <stdio.h>
#include <stdlib.h>

// ft_atoi
void test_atoi_inner(char *s)
{
    int rm;
    int rl;

    rm = ft_atoi(s);
    rl = atoi(s); 
    printf("equals: %d | input: %s | lib: %d | mine: %d\n", rl == rm, s, rl, rm);
}

void test_atoi() 
{
    char *s1 = "-123";
    char *s2 = "-2147483648";
    char *s3 = "2147483647";
    char *s4 = "0";
    char *s5 = "--0";
    char *s6 = "+023";
    char *s7 = "-34b56";
    char *s8 = "a)";
    char *s9 = "99988-a";

    test_atoi_inner(s1);
    test_atoi_inner(s2);
    test_atoi_inner(s3);
    test_atoi_inner(s4);
    test_atoi_inner(s5);
    test_atoi_inner(s6);
    test_atoi_inner(s7);
    test_atoi_inner(s8);
    test_atoi_inner(s9);
}

// ft_bzero
void test_bzero() {
    int len = 7;
    char *dest = malloc(sizeof(char) * len);

    ft_bzero(dest, len);

    while (len > 0)
    {
        printf("Is 0: %d \n", dest[len] == '\0');
        len--;
    }
}

// ft_memcpy
void test_ft_memcpy()
{
    int len = 5;
    char *dest = malloc(sizeof(char) * 5);
    char *str = "12345";

    char *new_d = ft_memcpy(dest, str, len);

    while (len > 0)
    {
        printf("v: %c | %c\n", *dest, *new_d);
        dest++;
        new_d++;
        len--;
    }
}

// ft_memccpy
void test_ft_memccpy() 
{
    int len = 6;
    char *dest = malloc(sizeof(char) * len);
    char *src = "123456789";

    char *res = ft_memccpy(dest,src, 'a', len);

    while(len > 0)
    {
        printf("val: %c\n", *res++);
        len -= 1;
    }
}

// ft_memmove
void test_ft_memmove()
{
    int len = 5;
    char *dst = malloc(sizeof(len));
    char *src = "1234567";
    
    char *res = ft_memmove(dst, src, len);

    while (len > 0) {
        printf("%c |", *res);
        res++;
        len--;
    }
}

// ft_memchr
void test_ft_memchr() 
{
    int len = 5;
    char *src = "123456";
    
    char *res = ft_memchr(src, 'a', len);

    printf("r: %p\n", res);
}

// ft_memcmp
void test_ft_memcmp() 
{
    int r = ft_memcmp("11119", "1111", 4);
    printf("r: %d\n", r);
}

int main() 
{
    // test_bzero();
    // test_ft_memcpy();
    // test_ft_memccpy();
    // test_ft_memmove();
    // test_ft_memchr();
    test_ft_memcmp();
    
    // test_atoi();
    return 0;
}


// TODO: test memset
// #include <string.h>
// #include <stdlib.h>

// int main()
// {
// 	int data_length = 4;
// 	void *data = malloc(sizeof(char) * data_length);
// 	void *data2 = malloc(sizeof(char) * data_length);
// 	unsigned char c = 121;
// 	int c2 = 666;

// 	void *res = memset(data, c, data_length);
	
// 	int i;
// 	for (i = 0; i < data_length; i ++) {
// 		printf("res: %c\n", ((char *)res)[i]);
// 	}
// 	for (i = 0; i < data_length; i ++) {
// 		printf("data: %c\n", ((char *)data)[i]);
// 	}

// 	printf("----\n");

// 	void *res2 = tf_memset(data2, c, data_length);
// 	for (i = 0; i < data_length; i++) {
// 		printf("res2: %c\n", ((char *)res2)[i]);
// 	}
// 	for (i = 0; i < data_length; i++) {
// 		printf("data2: %c\n", ((char *)data2)[i]);
// 	}
// }
