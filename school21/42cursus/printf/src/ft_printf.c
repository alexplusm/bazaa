//
// Created by Alexander Mogilevskiy on 27.12.2020.
//

#include <stdio.h>

// TODO: into header
#include <stdarg.h>
#include <stdlib.h>

// ---
typedef struct  f_arg {
    char*       flags;
    int         width;
    char        precision;
    char        format;
}               t_arg;

char *flags = "-0.*";
//char *format = "diouxXfFeEgGaAcsb"; // FROM MAN
char *format = "cspdiuxX%";

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

int is_format(char c)
{
    return ft_includes(c, format);
}

t_arg create_arg(char format, char flag) {
    t_arg item;

    item = malloc(sizeof(t_arg));
    if (item == NULL)
        return (item);
    item.format = format;
    item.flag = flag;
//    item.width = width;
    return item;
}

// ---

// %[flags][width][.precision][length]specifier

int ft_parse_one(char *str)
{
    char* flags;
    char format;
    size_t cursor;

    cursor = 0;
    while (is_flag(str[cursor]))
        cursor++;
    if (cursor > 0) {
        flags = malloc(sizeof(char) * (cursor + 1));
//        flags = ft_strdup(str, ) // TODO
    }
//     TODO: width | precision | ...
    if (is_format(str[cursor]))
    {
        format = str[cursor];
        cursor++;
    }

//     IF ERROR -> free(flags);

    return cursor;
}


int ft_parse(const char *f_str)
{
    int i;
    int count;
    char *copy_str;

    copy_str = (char *) f_str;

    i = 0;
    count = 0;
    while (*copy_str != '\0')
    {
        if (*copy_str == '%')
        {
            ft_parse_one(copy_str)
            count++;
        }
        copy_str++;
    }
    return count;
}

int ft_printf(const char *f_str, ...)
{
    int args_count;
    va_list valist;


    args_count = ft_parse(f_str);
    va_start(valist, args_count);
    int i = 0;
    while (i < args_count) {
        printf("%d : %d\n", i, va_arg(valist, int) );
        i++;
    }
    va_end(valist);
    return 0;
}


int main() {
    int res = ft_printf("hello %d %d %d", 100, 99, -1);

    printf("\n\nres: %d %% %% %% %\n", res);
}
