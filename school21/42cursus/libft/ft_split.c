#include "libft.h"

#include <stdio.h>

// TODO: compare with "ft_strdel" and "ft_free_res"
// void static ft_clean(char **array, size_t size)
// {
//     while(size > 0)
//     {
//         free(array[size]);
//         size--;
//     }
// }

static void	ft_strdel(char **as)
{
	if (as && *as)
	{
		free(*as);
		*as = NULL;
	}
}

static void		ft_free_res(char **words, size_t size)
{
	while (size--)
		ft_strdel(&(words[size]));
	free(*words);
}

size_t static ft_str_cnt_char(char const *s, char c)
{
    size_t strs_cnt;
    unsigned char *ptr;

    ptr = (unsigned char *)s;
    strs_cnt = 0;
    while (*ptr)
    {
        if (*ptr == c)
        {
            while (*ptr == c)
                ptr++;        
            strs_cnt++;
        }
        ptr++;
    }
    return strs_cnt;
}

char **ft_split(char const *s, char c)
{
    size_t strs_cnt;
    size_t start;
    size_t cursor;
    char **res;
    char *ptr;
    char *new_str;

    if (s == NULL)
        return NULL;
    
    char *set;
    set = malloc(sizeof(char) * 2);
    if (set == NULL)
        return NULL;
    set[0] = c;
    set[1] = '\0';

    size_t iii = 0;

    ptr = ft_strtrim(s, set);
    strs_cnt = ft_str_cnt_char(ptr, c) + 1;
    res = malloc(sizeof(char *) * (strs_cnt + 1));
    if (res == NULL)
        return (NULL);
    
    start = 0;
    cursor = 0;
    while (ptr[cursor] != '\0')
    {
        if (ptr[cursor] == c)
        {
            while (ptr[cursor] == c)
                cursor++;
            new_str = ft_substr(ptr, start, cursor - start);
            if (new_str == NULL)
            {
                ft_free_res(res, iii-1);
                return (NULL);
            }
                
            res[iii] = ft_strtrim(new_str, set);
            if (res[iii] == NULL)
            {
                ft_free_res(res, iii);
                return (NULL);
            }
            
            iii++;
            start = cursor;
        }
        cursor++;
    }

    new_str = ft_substr(ptr, start, cursor - start);
    res[iii] = new_str; // TODO: protect
    iii++;
    res[iii] = NULL;
    free(set);
    return (res);
}







// -------




// static size_t	ft_count_words(char const *s, char c)
// {
// 	size_t words;

// 	words = 0;
// 	while (*s)
// 	{
// 		while (*s == c)
// 			s++;
// 		if (*s)
// 		{
// 			words++;
// 			while (*s && *s != c)
// 				s++;
// 		}
// 	}
// 	return (words);
// }

// static char		*ft_get_word(char *word, char c)
// {
// 	char *start;

// 	start = word;
// 	while (*word && *word != c)
// 		word++;
// 	*word = '\0';
// 	return (ft_strdup(start));
// }

// static char		**ft_get_words(char *s, char c, size_t words_count)
// {
// 	char	**words;
// 	char	*word;
// 	size_t	i;

// 	i = 0;
// 	if ((words = (char **)malloc(sizeof(char *) * (words_count + 2))))
// 	{
// 		while (i < words_count)
// 		{
// 			while (*s == c)
// 				s++;
// 			if (*s)
// 			{
// 				if (!(word = ft_get_word(s, c)))
// 				{
// 					ft_free_res(words, i);
// 					return (NULL);
// 				}
// 				words[i++] = word;
// 				s += (ft_strlen(word) + 1);
// 			}
// 		}
// 		words[i] = NULL;
// 	}
// 	return (words);
// }

// char			**ft_split(char const *s, char c)
// {
// 	char	**words;
// 	char	*line;

// 	if (!s || !(line = ft_strdup((char *)s)))
// 		return (NULL);
// 	words = ft_get_words(line, c, ft_count_words(line, c));
// 	free(line);
// 	return (words);
// }
