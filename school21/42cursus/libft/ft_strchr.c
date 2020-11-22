char *ft_strchr(const char *s, int c)
{
    char *ptr;

    ptr = (char *)s;
    while(*ptr != '\0')
    {
        if (*ptr == c)
            return ptr;
        ptr++;
    }
    return (c == '\0') ? ptr : 0;
}