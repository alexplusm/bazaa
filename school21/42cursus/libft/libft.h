
// todo: include guards

void *tf_memset(void *b, int c, int len);

void ft_bzero(void *s, int n);

void *ft_memcpy(void *restrict dst, const void *restrict src, int len);

void *ft_memccpy(void *restrict dst, const void *restrict src, int c, int n);

void *ft_memmove(void *dst, const void *src, int len);

void *ft_memchr(const void *s, int c, int n);

int ft_memcmp(const void *s1, const void *s2, int n);

int ft_strlen(const char *s);

int ft_strlcpy(char * restrict dst, const char * restrict src, int dstsize);

// ---
int ft_atoi(const char *str);
