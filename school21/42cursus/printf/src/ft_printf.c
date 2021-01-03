//
// Created by Alexander Mogilevskiy on 27.12.2020.
//

#include <stdio.h>

// TODO: into header
#include <stdarg.h>
#include <stdlib.h>
#include <unistd.h>


// FROM LIBFT
// ----------

//void	*ft_memcpy(void *dst, const void *src, size_t len)
//{
//    unsigned char		*ptr_d;
//    unsigned char		*ptr_s;
//
//    if (len <= 0 || (dst == NULL && src == NULL))
//        return (dst);
//    ptr_d = (unsigned char *)dst;
//    ptr_s = (unsigned char *)src;
//    while (len > 0)
//    {
//        *ptr_d++ = *ptr_s++;
//        len -= 1;
//    }
//    return (dst);
//}

size_t	ft_strlen(const char *s)
{
    size_t	len;

    len = 0;
    while (s[len] != '\0')
        len += 1;
    return (len);
}

char	*ft_substr(char const *s, unsigned int start, size_t len)
{
    size_t	i;
    char	*str;

    i = 0;
    if (!s)
        return (NULL);
    if (start >= ft_strlen(s))
        start = ft_strlen(s);
    if (len > ft_strlen(s) - start)
        len = ft_strlen(s) - start;
    if (!(str = (char *)malloc(sizeof(char) * len + 1)))
        return (NULL);
    while (s[start + i] != '\0' && i < len)
    {
        str[i] = s[start + i];
        i++;
    }
    str[i] = '\0';
    return (str);
}

int	ft_isdigit(int c)
{
    return (c >= '0' && c <= '9');
}

int	ft_isspace_bonus(char c)
{
    return ((c >= 9 && c <= 13) || c == 32);
}

int		ft_atoi(const char *str)
{
    unsigned long	result;
    size_t			i;
    int				sign;

    result = 0;
    i = 0;
    while (ft_isspace_bonus(str[i]))
        i++;
    sign = (str[i] == '-') ? -1 : 1;
    if (str[i] == '-' || str[i] == '+')
        i++;
    while (ft_isdigit(str[i]))
    {
        result = result * 10 + (str[i++] - '0');
        if (result >= __LONG_MAX__ && sign == 1)
            return (-1);
        if ((result >= (unsigned long)__LONG_MAX__ + 1) && sign == -1)
            return (0);
    }
    return ((int)(result * sign));
}

// FROM LIBFT
// ----------


char *flags = "-0.*";
//char *format = "diouxXfFeEgGaAcsb"; // FROM MAN
char *specifiers = "cspdiuxX%";

// %[flags][width][.precision][length]specifier

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

typedef struct  f_arg {
    char        *flags;
    int         width;      // '*' is -1
    int         precision;  // '*' is -1
    char        specifier;
}               t_arg;

t_arg *create_arg(char *flags, int width, int precision, char specifier)
{
    t_arg *item;

    item = malloc(sizeof(t_arg));
    if (item == NULL)
        return (item);
    item->flags = flags;
    item->width = width;
    item->precision = precision;
    item->specifier = specifier;
    return item;
}

void debug_t_arg(t_arg *value)
{
    printf("t_arg: FLAGS: %s | WIDTH: %d | PREC: %d | SPECIFIER: %c\n",
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
        return (-1); // '*' symbol
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

t_arg *ft_parse_one(char *str, size_t *size)
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

int ft_printf(const char *f_str, ...)
{
    va_list valist;
    char    *str;

    size_t format_str_len;
    str = (char *)f_str;

    size_t fmt_specifiers_count = ft_get_fmt_specifiers_count(str);
    va_start(valist, fmt_specifiers_count);

    printf("### COUNT: %zu\n", fmt_specifiers_count);

    while (*str != '\0')
    {
        if (*str == '%')
        {
            str++; // info: to skip '%' char
            format_str_len = 0;
            t_arg *val = ft_parse_one(str, &format_str_len);


            if (val->specifier != '%')
                printf("AARRGG: %d\n", va_arg(valist, int));  // todo: info  if specifier != '%';

            debug_t_arg(val);

            printf("count: %zu\n", format_str_len);

            str += format_str_len;
        } else {
            write(0, str, 1);
            str++;
        }
    }

    va_end(valist);
    return (0);
}


int main() {
//    int res = ft_printf("hello %d %d %d", 100, 99, -1);
//    printf("\n\nres: %d %% %% %% %\n", res);

//    t_arg *a;
//    a = create_arg("123", 2, 0, 'a');
//    debug_t_arg(a);

    ft_printf("abcd kekus %-010c  looool  %-09999999.11d   %123%", 1, 2);

//    printf("%123%\n", 10000000);
}
