void *ft_memchr(const void *s, int c, int n)
{
    unsigned char *ptr;

    ptr = (unsigned char *)s;
    while (*ptr != c && n-- > 0)
        ptr++;
    return (n == -1) ? 0 : ptr;
}
