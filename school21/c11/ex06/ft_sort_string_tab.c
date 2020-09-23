#include <stdio.h>
#include <stdlib.h>

int ft_strcmp(char *s1, char *s2)
{
    int i;

    i = 0;
    while (s1[i] == s2[i] && (s1[i] != '\0' || s2[i] != '\0'))
        i += 1;
    return s1[i] - s2[i];
}

void swap(int i, int j,char **arr)
{
    char *temp;

    temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}

int partition(char **arr, unsigned int start_i, unsigned int end_i)
{
    int pivot_i;
    int min_i;
    int j;

    pivot_i = end_i;
    min_i = (start_i - 1);
    j = start_i;

    while (ft_strcmp(arr[end_i], arr[j]) >= 0 && j < end_i)
    {
        if (ft_strcmp(arr[pivot_i], arr[j]) > 0)
        {
            min_i += 1;
            swap(min_i, j, arr);
        }

        j += 1;
    }
    swap(min_i + 1, end_i, arr);
    return (min_i + 1);
}

void quick_sort(char **arr, int start_i, int end_i)
{
    if (start_i >= end_i)
        return ;

    int pi = partition(arr, start_i, end_i);
    
    quick_sort(arr, start_i, pi - 1);
    quick_sort(arr, pi + 1, end_i);
}

void ft_sort_string_tab()
{
    char **arr;
    
    arr = malloc( sizeof( char *) * 4);

    arr[0] = "7";
    arr[1] = "2";
    arr[2] = "1";
    arr[3] = NULL;

    int i = 0;
    while (arr[i])
        i += 1;
    
    quick_sort(arr, 0, i - 1);

    int j = 0;
    while (j < 3)
    {
        printf("%s\n", arr[j]);
        j++;
    }
}