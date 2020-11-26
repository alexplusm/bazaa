#include "libft.h"

#include <stdio.h>

char **ft_split(char const *s, char c)
{
//     unsigned char *ptr;
//     size_t strs_cnt;
//     char **res;
//     char *buff;

//     strs_cnt = 0;
//     ptr = (unsigned char *)s;
//     while (*ptr++)
//     {
//         if (*ptr == c)
//             strs_cnt++;
//     }
//     ptr = (unsigned char *)s;
    
//     res = malloc(sizeof(char *) * strs_cnt);
//     if (res == NULL)
//         return NULL;
    
    
//     size_t j = 0;

//     buff = ft_strchr((const char *)ptr, c);
//     while (buff != ptr)
//     {
        
//         j++;
//     }
//     buff;
//     res[j] = ft_strdup(buff); // TODO: process error

//     while (j < strs_cnt && *ptr)
//     {
//         buff = ft_strrchr((const char *)ptr, c);
//         // buff++;
//         res[j] = ft_strdup(buff); // TODO: process error
//         ptr += ft_strlen(res[j]);
//         j++;
//     }
//     return res;
    printf("%s, %c", s, c);
    return NULL;
}
