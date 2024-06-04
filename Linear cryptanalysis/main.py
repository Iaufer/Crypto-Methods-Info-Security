# sbox = [[0, 3],
#         [1, 2],
#         [2, 7],
#         [3, 6],
#         [4, 1],
#         [5, 4],
#         [6, 0],
#         [7, 5]
# ]

sbox = [[0, 6],
        [1, 7],
        [2, 4],
        [3, 3],
        [4, 2],
        [5, 5],
        [6, 1],
        [7, 0]
]

def count_weight ( num ):
    weight = 0
    while num :
        weight += 1
        num &= num -1
    return weight


def analysis_sub_block (sbox) -> dict :
    table = { i : { j: 0 for j in range (1 , 8) }
          for i in range (1 , 8) }
    for i in range (1 , 8) :
        for j in range (1 , 8) :
            for x in range (0 , 8) :
                if count_weight(( i & x) ^ (j & sbox[x][1]) ) % 2 == 0:
                    table [i ][ j ] += 1
    return table

def print_table(table):
    print("Table:")
    print("    " + " ".join(f"{j:2}" for j in range(1, 8)))
    for i in range(1, 8):
        row = " ".join(f"{table[i][j]:2}" for j in range(1, 8))
        print(f"{i:2}: {row}")

table = analysis_sub_block(sbox)
print_table(table)