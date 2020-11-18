
int is_digit(char c) {
    return c >= '0' && c <= '9';
}

int ft_pow(int num, int power) 
{
    if (power < 0)
        return 0;
    if (power == 0)
        return 1;
    return num * ft_pow(num, power - 1);
}

int ft_atoi(const char *str)
{
    int unsigned result = 0;
    int power = 0;
    int digit = 0;
    int sign = 1;

    if (*str == '-' || *str == '+') 
    {
        if (*str == '-')
            sign = -1;
        str++;
    }

    while (is_digit(*str))
        str++;
    str--;

    while (is_digit(*str))
    {
        digit = *str - '0';
        result += digit * ft_pow(10, power);
        str--;
        power++;
    }

    return result * sign;
}
