A = [
    [1.0, 2.0, 4.0],
    [3.0, 2.0, 2.0],
    [3.0, 1.0, 2.0]
]
B = [2.0, 4.0, 6.0]


def swap_rows(A, B, row1, row2):
    A[row1], A[row2] = A[row2], A[row1]
    B[row1], B[row2] = B[row2], B[row1]


def divide_row(A, B, row, divider):
    A[row] = [a / divider for a in A[row]]
    B[row] /= divider


def combine_rows(A, B, row, source_row, weight):
    A[row] = [(a + k * weight) for a, k in zip(A[row], A[source_row])]
    B[row] += B[source_row] * weight


def solve_SOLE_by_gaussian_elimination(A, B):
    column = 0
    while (column < len(B)):
        current_row = None

        for r in range(column, len(A)):
            if (current_row is None
                    or abs(A[r][column]) > abs(A[current_row][column])):
                current_row = r
        if current_row is None:
            print('no solution')
            return None

        if current_row != column:
            swap_rows(A, B, current_row, column)

        # normalize row
        divide_row(A, B, column, A[column][column])

        for r in range(column + 1, len(A)):
            combine_rows(A, B, r, column, -A[r][column])

        column += 1

    X = [0 for b in B]

    for i in range(len(B) - 1, -1, -1):
        X[i] = B[i] - sum(x * a for x, a in zip(X[(i + 1):], A[i][(i + 1):]))

    print("Result:")
    print("\n".join("X{}={:10.2f}".format(i + 1, x) for i, x in enumerate(X)))

    return X


solve_SOLE_by_gaussian_elimination(A, B)
