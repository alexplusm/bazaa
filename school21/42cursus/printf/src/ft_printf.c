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

void ft_put_char_by(char c, size_t count)
{
    while(count-- > 0)
        ft_putchar_fd(c, 1); // write(0, &c, 1);
}

char *build_int_fmt_str(int i_value, t_fmt_specifier *fmt_specifier)
{
    int negative;
    int value; // TODO: long?
    int spaces_size;
    int zeros_size;
    int right_align;

    char* result;

    negative = i_value > 0;
    value = negative ? -i_value : i_value;
    spaces_size = negative ? fmt_specifier->width - 1 : fmt_specifier->width;
    right_align = ft_includes('-', fmt_specifier->flags);

    if (right_align)
    {

    }
    else
    {

    }
}

size_t print_fmt_specifier(va_list *valist, t_fmt_specifier *fmt_specifier)
{
    size_t write_bites;
    int value;
    char *value_str;
    size_t fmt_value_size;

//    size_t real_width;
    write_bites = 0;

//    real_width = fmt_specifier->width > fmt_specifier->precision ? fmt_specifier->width : fmt_specifier->precision;

    if (fmt_specifier->specifier == 'd')
    {
        value = va_arg(*valist, int); // todo: long?
        value_str = ft_itoa(value);


        if ( fmt_specifier->precision != -2 && (fmt_specifier->width > fmt_specifier->precision + (int)ft_strlen(value_str)))
        {
            write_bites = fmt_specifier->width;

            if (!ft_includes('-', fmt_specifier->flags))
            {
                if (fmt_specifier->precision < (int)ft_strlen(value_str))
                {
                    ft_put_char_by(' ', fmt_specifier->width - (int)ft_strlen(value_str));
                    write(1, value_str, ft_strlen(value_str));
                    return (write_bites);
                }


                if (value < 0)
                {
                    free(value_str);
                    value_str = ft_itoa(-value);
                    ft_put_char_by(' ', fmt_specifier->width - fmt_specifier->precision - 1);
                    ft_put_char_by('-', 1);
                    if (fmt_specifier->precision > (int)ft_strlen(value_str))
                        ft_put_char_by('0', fmt_specifier->precision - (int)ft_strlen(value_str));
                    write(1, value_str, ft_strlen(value_str));
                    return (write_bites);
                }

                ft_put_char_by(' ', fmt_specifier->width - fmt_specifier->precision);
                if (fmt_specifier->precision > (int)ft_strlen(value_str))
                    ft_put_char_by('0', fmt_specifier->precision - (int)ft_strlen(value_str));
                write(1, value_str, ft_strlen(value_str));
                return (write_bites);
            }
            else
            {
                if (value < 0)
                {
                    free(value_str);
                    value_str = ft_itoa(-value);
                    ft_put_char_by('-', 1);
                    if (fmt_specifier->precision > (int)ft_strlen(value_str))
                        ft_put_char_by('0', fmt_specifier->precision - (int)ft_strlen(value_str));
                    write(1, value_str, ft_strlen(value_str));
                    if (fmt_specifier->precision > (int)ft_strlen(value_str))
                        ft_put_char_by(' ', fmt_specifier->width - fmt_specifier->precision - 1);
                    else
                        ft_put_char_by(' ', fmt_specifier->width - (int)ft_strlen(value_str) - 1);
                    return (write_bites);
                }
                // TODO: process negative
                if (fmt_specifier->precision > (int)ft_strlen(value_str))
                    ft_put_char_by('0', fmt_specifier->precision - (int)ft_strlen(value_str));
                write(1, value_str, ft_strlen(value_str));
                if (fmt_specifier->precision > (int)ft_strlen(value_str))
                    ft_put_char_by(' ', fmt_specifier->width - fmt_specifier->precision);
                else
                    ft_put_char_by(' ', fmt_specifier->width - (int)ft_strlen(value_str));
                return (write_bites);
            }
        }

        if (value == 0 && fmt_specifier->precision == 0)
        {
            free(value_str);
            value_str = malloc(sizeof(char) * 1);
            if (value_str == NULL)
                return (0);
            value_str[0] = '\0';
        }

        fmt_value_size = ft_strlen(value_str);

        if (fmt_specifier->width > (int)fmt_value_size) // TODO
        {
            if (fmt_specifier->flags != NULL)
            {
                if (ft_includes('-', fmt_specifier->flags))
                {
                    write_bites = fmt_specifier->width; // TODO : ?
                    write(1, value_str, fmt_value_size);
                    ft_put_char_by(32, fmt_specifier->width - fmt_value_size); // TODO: define SPACE_CONST
                }
                else if (ft_includes('0', fmt_specifier->flags))
                {
                    write_bites = fmt_specifier->width; // TODO : ?

                    if (value < 0)
                    {
                        ft_put_char_by('-', 1);
                        free(value_str);
                        value_str = ft_itoa(-value);
                        fmt_value_size = ft_strlen(value_str);
                        fmt_specifier->width--;
                    }

                    ft_put_char_by('0', fmt_specifier->width - fmt_value_size);
                    write(1, value_str, fmt_value_size);
                }
            }
            else
            {
                write_bites = fmt_specifier->width; // TODO : ?
                ft_put_char_by(' ', fmt_specifier->width - fmt_value_size);
                write(1, value_str, fmt_value_size);
            }
        }
        else if (fmt_specifier->precision > (int)fmt_value_size)
        {
            write_bites = fmt_specifier->precision;
            if (value < 0)
            {
                ft_put_char_by('-', 1);
                free(value_str);
                value_str = ft_itoa(-value);
                fmt_value_size = ft_strlen(value_str);
            }
            ft_put_char_by('0', fmt_specifier->precision - fmt_value_size);
            write(1, value_str, fmt_value_size);
        }
        else
        {
            write_bites = fmt_value_size; // TODO : ?
            write(1, value_str, fmt_value_size);
        }
        free(value_str);
    }
    return (write_bites);
}

// process asterisks
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
            write(1, str, 1);
            str++;
            write_bytes++;
        }
    }
    va_end(valist);
    return (write_bytes);
}

//int ft_printf(const char *f_str, ...)
//{
//    (void)f_str;
//    write(1, "this 17 number", 14);
//    return 14;
////    return printf("this 17 number");
//}
