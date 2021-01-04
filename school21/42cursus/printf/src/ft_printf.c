// TODO: test
#include <stdio.h>

#include "ft_printf.h"

/* %[flags][width][.precision][length]specifier */
char *flags = "-0";
char *specifiers = "cspdiuxX%";

int ft_includes(char c, char *str) {
    int i;

    i = 0;
    while (str[i] != '\0')
    {
        if (str[i] == c)
            return (1);
        i++;
    }
    return (0);
}

int is_flag(char c)
{
    return ft_includes(c, flags);
}

int is_specifier(char c)
{
    return ft_includes(c, specifiers);
}

typedef struct  s_fmt_specifier {
    char        *flags;
    int         width;      // '*' is -1
    int         precision;  // '*' is -1
    char        specifier;
}               t_fmt_specifier;

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

t_fmt_specifier *create_arg(char *flags, int width, int precision, char specifier)
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

void debug_t_fmt_specifier(t_fmt_specifier *value)
{
    printf("t_fmt_specifier: FLAGS: %s | WIDTH: %d | PREC: %d | SPECIFIER: %c\n",
           value->flags, value->width, value->precision, value->specifier);
}

size_t ft_get_fmt_specifiers_count(char *str)
{
    size_t i;
    size_t count;

    i = 0;
    count = 0;
//     TODO: выделить в функцию !
    while (str[i] != '\0')
    {
        if (str[i] == '%') {
            i++;
            while (is_flag(str[i]))
                i++;
            if (str[i] == '*')
                i++;
            else
                while (ft_isdigit(str[i]))
                    i++;
            if (str[i] == '.')
            {
                str++;
                if (str[i] == '*')
                    str++;
                else
                    while(ft_isdigit(str[i]))
                        i++;
            }
            if (str[i] != '%')
                count++;
        }
        i++;
    }
    return (count);
}

int parse_width_or_precision(char *str, size_t *cursor)
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

t_fmt_specifier *ft_parse_one(char *str, size_t *size)
{
    char* flags;
    int width;
    int precision;
    char specifier;
    size_t cursor;

    precision = 0;
    cursor = 0;
    flags = NULL;
    specifier = '0';
    while (is_flag(str[cursor]))
        cursor++;
    if (cursor > 0) {
        flags = ft_substr(str,0, cursor);
    }
    width = parse_width_or_precision(str, &cursor); // width
    if (str[cursor] == '.') // precision
    {
        cursor++;
        precision = parse_width_or_precision(str, &cursor);
    }
    if (is_specifier(str[cursor]))
        specifier = str[cursor++];
    *size = cursor;     // TODO: if error -> free(flags); ? ? ?
    return create_arg(flags, width, precision, specifier);
}

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

//void kek(va_list *valist, size_t size, ...)
//{
//    va_start(*valist, size);
//}

int ft_printf(const char *f_str, ...)
{
    va_list valist;
    t_fmt_specifier *fmt_specifier;
    char    *str;
    size_t write_bytes;

    write_bytes = 0;

    size_t fmt_str_len;
    str = (char *)f_str;

//    size_t fmt_specifiers_count = ft_get_fmt_specifiers_count(str);
//    va_start(valist, fmt_specifiers_count);
    va_start(valist, f_str); // TODO: wtf?
//    kek(&valist, fmt_specifiers_count);

//    printf("### COUNT: %zu\n", fmt_specifiers_count);

    while (*str != '\0')
    {
        if (*str == '%')
        {
            str++; // info: to skip '%' char
            fmt_str_len = 0;
            fmt_specifier = ft_parse_one(str, &fmt_str_len);
            update_fmt_specifier(&valist, fmt_specifier);

//            debug_t_fmt_specifier(fmt_specifier);

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


//int main() {
//    int res = ft_printf("hello %d %d %d", 100, 99, -1);
//    printf("\n\nres: %d %% %% %% %\n", res);

//    t_fmt_specifier *a;
//    a = create_arg("123", 2, 0, 'a');
//    debug_t_fmt_specifier(a);

//    ft_printf("abcd kekus %-010c  looool  %-09999999.11d   %123%", 1, 2);
//    ft_printf("1lol %-10d", 123);
//    printf("|\n");
//    printf("1lol %-10d", 123);
//    printf("|\n");
////    ---
//    ft_printf("2lol %10d", 123);
//    printf("|\n");
//    printf("2lol %10d", 123);
//    printf("|\n");
////    ---
//    ft_printf("3lol %-010d", 123);
//    printf("|\n");
//    printf("3lol %-010d", 123);
//    printf("|\n");
//    ---
//    ft_printf("4: %-*d", 10, 123);
//    printf("|\n");
//    printf("4: %-*d",10, 123);
//    printf("|\n");
//    ---
//    printf("%.*s", 3, "abcdef");
//    printf("\n");
//    printf("%.3s", "abcdef");

//    printf("%9%, %d\n", 40404040400404);
//}
