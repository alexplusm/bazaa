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

#include "ft_printf.h"

int is_flag(char c)
{
    return ft_includes(c, flags);
}

int is_specifier(char c)
{
    return ft_includes(c, specifiers);
}

static t_fmt_specifier *ft_create_fmt_specifier(char *flags, int width, int precision, char specifier)
{
    t_fmt_specifier *item;

    item = malloc(sizeof(t_fmt_specifier));
    if (item == NULL)
        return (item);
    item->flags = flags;
    item->width = width;
    item->precision = precision;
    item->specifier = specifier;
    return item;
}

static int ft_parse_width_or_precision(char *str, size_t *cursor)
{
    int value;
    char *num_str;
    size_t inner_cursor;

    inner_cursor = *cursor;
    value = 0;
    if (str[inner_cursor] == '*')
    {
        *cursor = inner_cursor + 1;
        return (-1); // '*' symbol
    }

    size_t start = inner_cursor;
    while (ft_isdigit(str[inner_cursor]))
        inner_cursor++;
    if (start < inner_cursor)
    {
        num_str = ft_substr(str, start, inner_cursor);
        value = ft_atoi(num_str);
        free(num_str);
    }
    *cursor = inner_cursor;
    return (value);
}

t_fmt_specifier *ft_parse_fmt_specifier(char *str, size_t *size)
{
    char* flags;
    int width;
    int precision;
    char specifier;
    size_t cursor;

    precision = -2; // undefined value
    cursor = 0;
    flags = NULL;
    specifier = '0';
    while (is_flag(str[cursor]))
        cursor++;
    if (cursor > 0) {
        flags = ft_substr(str,0, cursor);
    }
    width = ft_parse_width_or_precision(str, &cursor); // width
    if (str[cursor] == '.') // precision
    {
        cursor++;
        precision = ft_parse_width_or_precision(str, &cursor);
    }
    if (is_specifier(str[cursor]))
        specifier = str[cursor++];
    *size = cursor;     // TODO: if error -> free(flags); ? ? ?
    return ft_create_fmt_specifier(flags, width, precision, specifier);
}
