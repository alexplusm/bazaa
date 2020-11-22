
// TODO: if (ft_isascii)
int ft_isdigit(int c)
{
    char digit = '0';

    while (digit <= '9')
    {
        if (c == digit)
            return c;
        digit++;
    }
    return 0;
}

