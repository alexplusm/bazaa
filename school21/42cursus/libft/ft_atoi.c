/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_atoi.c                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:08:22 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:08:27 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

// TODO: INTO lib and test
int ft_isspace(char c)
{
    if ((c >= 9 && c <= 13) || c == 32)
        return c;
    return 0;
}

// TODO: INTO lib and test
int ft_pow(int num, int power) 
{
    if (power < 0)
        return 0;
    if (power == 0)
        return 1;
    return num * ft_pow(num, power - 1);
}

// int ft_atoi(const char *str)
// {
//     int unsigned result = 0;
//     int power = 0;
//     int digit = 0;
//     int sign = 1;
    
//     while (ft_isspace(*str))
//         str++;
//     if (*str == '-' || *str == '+') 
//     {
//         if (*str == '-')
//             sign = -1;
//         str++;
//     }
//     while (ft_isdigit(*str))
//         str++;
//     str--;
//     while (ft_isdigit(*str))
//     {
//         digit = *str - '0';
//         result += digit * ft_pow(10, power);
//         str--;
//         power++;
//     }
//     return result * sign;
// }




# define FT_ULONG_MAX	((unsigned long)(~0L))
# define FT_LONG_MAX	((long)(FT_ULONG_MAX >> 1))

int		ft_atoi(const char *str)
{
	unsigned long	result;
	unsigned long	border;
	size_t			i;
	int				sign;

	result = 0;
	border = (unsigned long)(FT_LONG_MAX / 10);
	i = 0;
	while (ft_isspace(str[i]))
		i++;
	sign = (str[i] == '-') ? -1 : 1;
	if (str[i] == '-' || str[i] == '+')
		i++;
	while (ft_isdigit(str[i]))
	{
		if ((result > border || (result == border && (str[i] - '0') > 7))
															&& sign == 1)
			return (-1);
		else if ((result > border || (result == border && (str[i] - '0') > 8))
																&& sign == -1)
			return (0);
		result = result * 10 + (str[i++] - '0');
	}
	return ((int)(result * sign));
}
