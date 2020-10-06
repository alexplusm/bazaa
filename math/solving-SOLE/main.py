# # from copy import deepcopy

# # def transform_matrix_to_reduced_row_echelon_form(A_input):
# #     A = deepcopy(A_input)

# #     row_count = len(A)
# #     column_count = len(A[0])
    
# #     lead = 0

# #     for row_index in range(row_count):
# #         if lead >= column_count:
# #             return
# #         i = row_index

# #         while A[i][lead] == 0:
# #             i += 1
# #             if i == row_count:
# #                 i = row_index
# #                 lead += 1

# #                 if column_count == lead:
# #                     return
        
# #         A[i], A[row_index] = A[row_index], A[i]

# #         current_item = A[row_index][lead]
# #         A[row_index] = [item / float(current_item) for item in A[row_index]]

# #         for i in range(row_count):
# #             if i != row_index:
# #                 current_item = A[i][lead]
# #                 A[i] = [iv - current_item * rv for rv, iv in zip(A[row_index], A[i])]

# #         lead += 1
# #     return A

# # def solve_SOLE_by_gaussian_elimination(A, B):
# #     """
# #     todo: write description
# #     return vector
# #     """
# #     A_rref = transform_matrix_to_reduced_row_echelon_form(A)

# #     print(A_rref)

# #     result = []
    
# #     cursor = len(A_rref[0])

# #     # for B_index in reversed(range(len(B))):
# #     #     factor_x = A_rref[B_index][cursor]
# #     #     x = B[B_index] / float(factor_x)

# # mtx = [
# #    [ 1, 2, 4],
# #    [3, 2, 2],
# #    [3, 1, 2,]
# # ]
# # b = [2, 4]
 
# # result = solve_SOLE_by_gaussian_elimination(mtx, b)

# def ToReducedRowEchelonForm( M):
#     if not M: return
#     lead = 0
#     rowCount = len(M)
#     columnCount = len(M[0])
#     for r in range(rowCount):
#         if lead >= columnCount:
#             return
#         i = r
#         while M[i][lead] == 0:
#             i += 1
#             if i == rowCount:
#                 i = r
#                 lead += 1
#                 if columnCount == lead:
#                     return
#         M[i],M[r] = M[r],M[i]
#         lv = M[r][lead]
#         M[r] = [ mrx / float(lv) for mrx in M[r]]
#         for i in range(rowCount):
#             if i != r:
#                 lv = M[i][lead]
#                 M[i] = [ iv - lv*rv for rv,iv in zip(M[r],M[i])]
#         lead += 1

# mtx1 = [
#    [ 1, 2, 8],
#    [3, 2, 2],
#    [3, 1, 2,]
# ]
# print('------')
# ToReducedRowEchelonForm( mtx1 )
# for rw in mtx1:
#   print('; '.join( (str(rv) for rv in rw) ))

# todo: new version

# --- исходные данные
# myA = [
#     [1.0, -2.0, 3.0, -4.0],
#     [3.0, 3.0, -5.0, -1.0],
#     [3.0, 0.0, 3.0, -10.0],
#     [-2.0, 1.0, 2.0, -3.0]
# ]
# myB = [2.0, -3.0, 8.0, 5.0]

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
