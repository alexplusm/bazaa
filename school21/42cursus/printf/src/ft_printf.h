/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_atoi.c                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <cdeon@student.21-school.ru>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/01/02 09:08:22 by cdeon             #+#    #+#             */
/*   Updated: 2021/01/02 09:08:27 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#ifndef FT_PRINTF
#define FT_PRINTF

// TODO: test !!!
# include <stdio.h>

# include <stdarg.h>
# include <stdlib.h>
# include <unistd.h>

# include "libft/libft.h"

typedef struct  s_fmt_specifier {
    char        *flags;
    int         width;      // '*' is -1
    int         precision;  // '*' is -1
    char        specifier;
}               t_fmt_specifier;

static const char *flags = "-0";
static const char *specifiers = "cspdiuxX%";

int ft_printf(const char *f_str, ...);

// TODO: remove?
t_fmt_specifier *ft_parse_fmt_specifier(char *str, size_t *size);

int ft_includes(char c, const char *str);

#endif
