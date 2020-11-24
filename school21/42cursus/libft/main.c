#include "libft.h"

#include <stdlib.h>

#include <stdio.h>
#include <string.h>
#include <ctype.h>

char *success = "SUCCESS";
char *failure = "@ FAILURE";

char *res(int b)
{
    if (b == 0)
        return failure;
    return success;
}

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

// ft_strlen
void test_ft_strlen() 
{
    int r1 = ft_strlen("");
    printf("r1: %d\n", r1);
}

// ft_strlcpy
void inner_ft_strlcpy(char *src, ft_size_t dst_len, ft_size_t len)
{    
    char *dst_m = malloc(sizeof(char) * dst_len);
    char *dst_l = malloc(sizeof(char) * dst_len);
    
    int mr = ft_strlcpy(dst_m, src, len);
    int lr = strlcpy(dst_l, src, len);

    printf("lib: %s (return : %d) | ", dst_l, lr);
    printf("mine: %s (return : %d)\n",dst_m, mr);
}

void test_ft_strlcpy()
{
    inner_ft_strlcpy("123", 1 ,2);
    inner_ft_strlcpy("123", 1, 0);
    inner_ft_strlcpy("123456789", 1, -2147483649);
    inner_ft_strlcpy("123456789", 1, 3);
    inner_ft_strlcpy("", 1, 3);
}

// ft_strlcat
void inner_ft_strlcat(const char *dst,const char *src, ft_size_t dst_len) {
    char *dst_l = malloc(sizeof(dst));
    char *dst_m = malloc(sizeof(dst));
    int i = 0;

    while (dst[i] != '\0')
    {
        dst_l[i] = dst[i];
        dst_m[i] = dst[i];
        i += 1;
    }
    
    int mr = ft_strlcat(dst_m, src, dst_len);
    int lr = strlcat(dst_l, src, dst_len);

    printf("mine: %d | %s | ", mr, dst_m);
    printf("lib: %d | %s\n", lr, dst_l);
}

void test_ft_strlcat() {
    inner_ft_strlcat("dest", "src", 10);
    inner_ft_strlcat("dest", "src", 5);
    inner_ft_strlcat("d", "src", 5);
    inner_ft_strlcat("", "src", -1);
    inner_ft_strlcat("", "", -1000000000000000);
    inner_ft_strlcat("abc", "123", 7);
}

// ft_strchr
void inner_ft_strchr(const char *s, int c)
{
    char *lr = strchr(s, c);
    char *mr = ft_strchr(s, c);

    printf("mine: %s | %p | ", mr, mr);
    printf("lib: %s | %p \n", lr, lr);
}

void test_ft_strchr()
{
    inner_ft_strchr("1234", '2');
    inner_ft_strchr("1234", '\0');
    inner_ft_strchr("", '\0');
    inner_ft_strchr("", 'a');
    inner_ft_strchr("---", 'a');
}

// ft_strrchr
void inner_ft_strrchr(const char *s, int c)
{
    char *lr = strrchr(s, c);
    char *mr = ft_strrchr(s, c);

    printf("mine: %s | %p | ", mr, mr);
    printf("lib: %s | %p \n", lr, lr);
}

void test_ft_strrchr()
{
    inner_ft_strrchr("123", '1');
    inner_ft_strrchr("123", '2');
    inner_ft_strrchr("12311", '1');
    inner_ft_strrchr("", '1');
}

// ft_strnstr
void inner_ft_strnstr(const char *haystack, const char *needle, ft_size_t len)
{
    char *lr = strnstr(haystack, needle, len);
    char *mr = ft_strnstr(haystack, needle, len);

    printf("mine: %s | %p | ", mr, mr);
    printf("lib: %s | %p \n", lr, lr);
}

void test_ft_strnstr()
{
    inner_ft_strnstr("112", "1", 2);
    inner_ft_strnstr("123", "2", 2);
    inner_ft_strnstr("12345", "23", 4);
    inner_ft_strnstr("1234-2345-aa", "2345", 4); // WTF?! 
}

// ft_strncmp
void inner_ft_strncmp(const char *s1, const char *s2, ft_size_t n)
{
    int mr = ft_strncmp(s1, s2, n);
    int lr = strncmp(s1, s2, n);

    printf("mine: %d | ", mr);
    printf("lib: %d\n", lr);
}

void test_ft_strncmp()
{
    inner_ft_strncmp("123", "122", 3);
    inner_ft_strncmp("123", "123", 3);
    inner_ft_strncmp("123ab", "12345", 3);
    inner_ft_strncmp("123ab", "12345", 100);
    inner_ft_strncmp("", "b", 100);
    inner_ft_strncmp("a", "", 100);
    inner_ft_strncmp("1", "1", 0);
    inner_ft_strncmp("1a", "1b", 1);
}

// ft_isalpha
void test_ft_isalpha()
{
    int r = ft_isalpha('a');
    printf("r: %d\n", r);
}

void test_ft_isdigit()
{
    int r = ft_isdigit('d');
    printf("r: %d\n", r);
}

void test_ft_isalnum()
{
    int r = ft_isalnum('?');
    printf("r: %d\n", r);
}

// ft_isascii
void inner_ft_isascii(char c)
{
    int lr = isascii(c);
    int mr = ft_isascii(c);

    printf("mine: %d | ", mr);
    printf("lib: %d\n", lr);
}

void test_ft_isascii()
{
    inner_ft_isascii('a');
    inner_ft_isascii('?');
    inner_ft_isascii(-1);
    inner_ft_isascii(0);
}

// ft_isprint
void inner_ft_isprint(int c)
{
    int lr = isprint(c);
    int mr = ft_isprint(c);

    printf("val: %c (%d) # ", c, c);
    printf("mine: %d | ", mr);
    printf("lib: %d\n", lr);
}

void test_ft_isprint()
{
    int i = 0;
    while (i < 128)
        inner_ft_isprint(i++);   
}

// ft_toupper
void inner_ft_toupper(int c)
{
    int lr = toupper(c);
    int mr = ft_toupper(c);

    printf("%s: ", res(lr == mr));
    printf("mine: %c (%d) | ", mr, mr);
    printf("lib: %c (%d) \n", lr, lr);
}

void test_ft_toupper()
{
    int i = 0;
    while (i < 128)
        inner_ft_toupper(i++);
}

// ft_tolower
void inner_ft_tolower(int c)
{
    int lr = tolower(c);
    int mr = ft_tolower(c);

    printf("%s: ", res(lr == mr));
    printf("mine: %c (%d) | ", mr, mr);
    printf("lib: %c (%d) \n", lr, lr);
}

void test_ft_tolower()
{    
    int i = 0;
    while (i < 128)
        inner_ft_tolower(i++);
}


void test_ft_calloc()
{
    int count = 2;
    int size = 5;
    char *s = ft_calloc(count, size);
    int i = 0;
    while (i < count * size)
    {
        printf("%d | ", s[i]);
        i++;
    }
    printf("\n");
}

// -----------------
int main() 
{
    // test_bzero();
    // test_ft_memcpy();
    // test_ft_memccpy();
    // test_ft_memmove();
    // test_ft_memchr();
    // test_ft_memcmp();
    
    // test_ft_strlen();
    // test_ft_strlcpy();
    // test_ft_strlcat();
    // test_ft_strchr();
    // test_ft_strrchr();
    // test_ft_strnstr();
    // test_ft_strncmp();

    // test_atoi();
    // test_ft_isalpha();
    // test_ft_isdigit();
    // test_ft_isalnum();
    // test_ft_isascii();
    // test_ft_isprint();
    // test_ft_tolower();

    test_ft_calloc();
    
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
