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

char *ft_prepare_flags(char *flags)
{
    char *new_flags;
    int i;
    int j;

    if (flags == NULL)
        return (flags);
    if (ft_includes('-', flags) && ft_includes('0', flags))
    {
        new_flags = malloc(sizeof(char) * ft_strlen(flags));
        if (new_flags == NULL)
            return (new_flags);
        i = 0;
        j = 0;
        while (flags[i] != '\n')
        {
            if (flags[i] == '0')
            {
                i++;
                continue;
            }
            new_flags[j] = flags[i];
            i++;
            j++;
        }
        free(flags);
        return (new_flags);
    }
    return (flags);
}

void debug_t_fmt_specifier(t_fmt_specifier *value)
{
    printf("t_fmt_specifier: FLAGS: %s | WIDTH: %d | PREC: %d | SPECIFIER: %c\n",
           value->flags, value->width, value->precision, value->specifier);
}

// TODO: why this func not need ?
//size_t ft_get_fmt_specifiers_count(char *str)
//{
//    size_t i;
//    size_t count;
//
//    i = 0;
//    count = 0;
////     TODO: выделить в функцию !
//    while (str[i] != '\0')
//    {
//        if (str[i] == '%') {
//            i++;
//            while (is_flag(str[i]))
//                i++;
//            if (str[i] == '*')
//                i++;
//            else
//                while (ft_isdigit(str[i]))
//                    i++;
//            if (str[i] == '.')
//            {
//                str++;
//                if (str[i] == '*')
//                    str++;
//                else
//                    while(ft_isdigit(str[i]))
//                        i++;
//            }
//            if (str[i] != '%')
//                count++;
//        }
//        i++;
//    }
//    return (count);
//}
// --------

void ft_put_char_by(char c, size_t count)
{
    while(count-- > 0)
        write(0, &c, 1);
}

size_t print_fmt_specifier(va_list *valist, t_fmt_specifier *fmt_specifier)
{
    size_t write_bites;
    int value;
    char *value_str;
    size_t fmt_value_size;

    write_bites = 0;

    if (fmt_specifier->specifier == 'd')
    {
        value = va_arg(*valist, int);
        value_str = ft_itoa(value);
        fmt_value_size = ft_strlen(value_str);

        if (fmt_specifier->width > (int)fmt_value_size)
        {
            if (fmt_specifier->flags != NULL)
            {
                if (ft_includes('-', fmt_specifier->flags))
                {
                    write_bites = fmt_specifier->width; // TODO : ?
                    write(0, value_str, fmt_value_size);
                    ft_put_char_by(' ', fmt_specifier->width - fmt_value_size);
                }
                else if (ft_includes('0', fmt_specifier->flags))
                {
                    write_bites = fmt_specifier->width; // TODO : ?
                    ft_put_char_by('0', fmt_specifier->width - fmt_value_size);
                    write(0, value_str, fmt_value_size);
                }
            }
            else
            {
                write_bites = fmt_specifier->width; // TODO : ?
                ft_put_char_by(' ', fmt_specifier->width - fmt_value_size);
                write(0, value_str, fmt_value_size);
            }
        }
        else
        {
            write_bites = fmt_value_size; // TODO : ?
            write(0, value_str, fmt_value_size);
        }


        free(value_str);

    }

    return (write_bites);

    //            if (val->specifier != '%')
    //                printf("AARRGG: %d\n", va_arg(valist, int)); // TODO: print arg
}

void update_fmt_specifier(va_list *valist, t_fmt_specifier *fmt_specifier)
{
    int value;

    if (fmt_specifier->width == -1)
    {
        value = va_arg(*valist, int);
        fmt_specifier->width = value;
    }
    if (fmt_specifier->precision == -1)
    {
        value = va_arg(*valist, int);
        fmt_specifier->precision = value;
    }
}

int ft_printf(const char *f_str, ...)
{
    va_list valist;
    t_fmt_specifier *fmt_specifier;
    char    *str;
    size_t fmt_str_len;
    size_t write_bytes;

    // size_t fmt_specifiers_count = ft_get_fmt_specifiers_count(str);
    // va_start(valist, fmt_specifiers_count);
    va_start(valist, f_str); // TODO: wtf?

    write_bytes = 0;
    str = (char *)f_str;

    while (*str != '\0')
    {
        if (*str == '%')
        {
            str++; // info: to skip '%' char
            fmt_str_len = 0;
            fmt_specifier = ft_parse_fmt_specifier(str, &fmt_str_len);
            update_fmt_specifier(&valist, fmt_specifier);

            write_bytes += print_fmt_specifier(&valist, fmt_specifier);
            str += fmt_str_len;
        } else {
            write(0, str, 1);
            str++;
            write_bytes++;
        }
    }
    va_end(valist);
    return (write_bytes);
}
