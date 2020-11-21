void *tf_memset(void *b, int c, int len)
{
    // int i;

    // // todo: while
    // for (i = 0; i < len; i++) {
    //     ((int *)b)[i] = c;
    //     printf("i: %d | value: %c | char: %c \n", i, ((char *)b)[i], c);
    // }
    // return b;

    unsigned char *ptr = b;
    while (len-- > 0)
        *ptr++ = c;
    return b;
}
