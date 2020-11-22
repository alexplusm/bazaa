
char *ft_strrchr(const char *s, int c)
{
    char *ptr;

    ptr = (char *)s;
    while (*ptr != '\0')
        ptr++;
    while (ptr != s)
    {
        if (*ptr == c)
            return ptr;
        ptr--;
    }
    return (*ptr == c) ? ptr : 0;
}
