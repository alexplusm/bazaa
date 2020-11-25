void *ft_memchr(const void *s, int c, int n)
{
    char *ptr;

    ptr = (char *)s;
    while (n-- > 0 && *ptr != c)
        ptr++;
    return (n == -1) ? 0 : ptr;
}
