#include <stdio.h>

#include "ft_printf.h"

void test1()
{
//    int lr, mr;
//
//    mr = ft_printf("%5.0d", 0);
//    printf("|");
//    printf("\n");
//    lr = printf("%5.0d", 0);
//    printf("|");
//    printf("\n mine: %d | lib: %d\n", mr, lr);


    // ---

    ft_printf("this %d number", 17);
    printf("\n");
    printf("this %d number", 17);


//    printf("%5.0d|\n", 0);
//    ft_printf("%5.0d\n", -1);
//
//    printf("%5.d|\n", 0);
}

int main()
{
    test1();
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
