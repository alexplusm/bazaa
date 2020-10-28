#include <stdio.h>

int		is_printable(char c)
{
	return (c >= 32 && c <= 126);
}

int		is_space(char c)
{
	return ((c >= 9 && c <= 13) || c == ' ');
}

int		validate_base3(char *base)
{
	char	*ptr;
	int		i;
	int		vals[95];

	i = 0;
	while (i < 95)
    {
        vals[i] = 0;
        i += 1;
    }
	i = 0;
	while (base[i])
	{
		if (!is_printable(base[i]))
			return (0);
		if (base[i] == '+' || base[i] == '-')
			return (0);
		if (vals[base[i] - 32])
			return (0);
		vals[base[i] - 32] = 1;
		i++;
	}
	if (i < 2)
		return (0);
	return (i);
}

int		parse_positive(char *str, char *base, int base_len)
{
	char	*ptr;
	int		i;
	int		v;

	v = 0;
	while (*str)
	{
		i = 0;
		ptr = base;
		while (*ptr)
		{
            printf("*ptr: %d\n", *ptr);

			if (*str == *ptr)
				break ;
			i++;
			ptr++;
		}
		if (i >= base_len)
			break ;
		v *= base_len;
		v -= i;
		str++;
	}
	return (v);
}

int		ft_atoi_base(char *str, char *base)
{
	int		base_len;
	int		v;
	int		is_positive;

	base_len = validate_base3(base);

	if (!base_len)
		return (0);
	is_positive = 0;
	while (is_space(*str))
		str++;
	if (*str == '+' || *str == '-')
	{
		is_positive = -1;
		if (*str == '+')
			is_positive = 1;
		str++;
	}
	v = parse_positive(str, base, base_len);
	v *= -(is_positive + !is_positive);
	return (v);
}




//--------



#include <stdio.h>

int main()
{
    char *str = "- c ";
    char *base = " bc";

    int res = ft_atoi_base(str, base);
    
    printf("res: %d\n", res);
}
