
// TODO: if (ft_isascii)
int ft_isalpha(int c)
{
    char ch;
    ch = 'A';

    while(ch <= 'Z')
    {
        if (c == ch)
            return c;
        ch += 1;
    }
    ch = 'a';
    while(ch <= 'z')
    {
        if (c == ch)
            return c;
        ch += 1;
    }

    return 0;
}
