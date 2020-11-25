
// todo: include guards
#include <stdlib.h>
#include <errno.h>

typedef unsigned long ft_size_t;

void *tf_memset(void *b, int c, int len);

void ft_bzero(void *s, int n);

void *ft_memcpy(void *dst, const void *src, int len);

void *ft_memccpy(void *dst, const void *src, int c, int n);

void *ft_memmove(void *dst, const void *src, int len);

void *ft_memchr(const void *s, int c, int n);

int ft_memcmp(const void *s1, const void *s2, int n);

// ---

size_t ft_strlen(const char *s);

int ft_strlcpy(char *dst, const char *src, ft_size_t dstsize);

size_t ft_strlcat(char *dst, const char *src, size_t dstsize);

char *ft_strchr(const char *s, int c);

char *ft_strrchr(const char *s, int c);

char *ft_strnstr(const char *haystack, const char *needle, ft_size_t len);

int ft_strncmp(const char *s1, const char *s2, ft_size_t n);

// ---

int ft_atoi(const char *str);

int ft_isalpha(int c);

int ft_isdigit(int c);

int ft_isalnum(int c);

int ft_isascii(int c);

int ft_isprint(int c);

int ft_toupper(int c);

int ft_tolower(int c);

void *ft_calloc(ft_size_t count, ft_size_t size);

char *ft_strdup(const char *s1);

// ---

char *ft_substr(char const *s, unsigned int start,size_t len);

char *ft_strjoin(char const *s1, char const *s2);
