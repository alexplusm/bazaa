### INFO

exclude /pft

#### Makefile snippet for testing
test: $(NAME)
    $(CC) $(NAME) main.c -o main.out
    ./main.out

"%[flags][width][.precision][length]specifier"


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
