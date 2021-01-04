#include <stdio.h>

#include "ft_printf.h"

int main()
{
    int r = ft_printf("kek");
    printf("res: %d\n", r);
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
