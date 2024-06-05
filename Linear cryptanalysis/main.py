import random


def encoding(key,plaintext,sbox):
	

	#round 1
	ciphertext=[]
	for i in range(9): #сложение с ключом
		ciphertext.append(plaintext[i]^key[i])
	for i in range(3): #замена
		index=int((str(ciphertext[i*3])+str(ciphertext[i*3+1])+str(ciphertext[i*3+2])),base=2)
		new_value=sbox[index] #замена для i-го блока
		# print(index, new_value)
		for j in range(3):
			ciphertext[i*3+j]=new_value[j]
			# print(ciphertext)

	new_ciphertext=[0 for i in range(9)] #перестановка
	for i in range(3):
		for j in range(3):
			new_ciphertext[j*3+i]=ciphertext[i*3+j]
	ciphertext=new_ciphertext
	
	#round 2
	for i in range(9): #сложение с ключом
		ciphertext[i]=ciphertext[i]^key[i]
	for i in range(3): #замена
		index=int((str(ciphertext[i*3])+str(ciphertext[i*3+1])+str(ciphertext[i*3+2])),base=2)
		new_value=sbox[index] #замена для i-го блока
		for j in range(3):
			ciphertext[i*3+j]=new_value[j]

		 #перестановка
	new_ciphertext=[0 for i in range(9)]
	for i in range(3):
		for j in range(3):
			new_ciphertext[j*3+i]=ciphertext[i*3+j]
	ciphertext=new_ciphertext
	
	#round 3
	for i in range(9): #сложение с ключом
		ciphertext[i]=ciphertext[i]^key[i]
	for i in range(3): #замена
		index=int((str(ciphertext[i*3])+str(ciphertext[i*3+1])+str(ciphertext[i*3+2])),base=2)
		new_value=sbox[index] #замена для i-го блока
		for j in range(3):
			ciphertext[i*3+j]=new_value[j]
	ciphertext=new_ciphertext
	
	#финальное сложение с ключом
	for i in range(9): #сложение с ключом
		ciphertext[i]=ciphertext[i]^key[i]

	return ciphertext

# print("Шифрование: ")
sbox=[[1,1,0],[1,1,1],[1,0,0],[0,1,1],[0,1,0],[1,0,1],[0,0,1],[0,0,1]] 
# key=[0,1,0,1,0,1,1,1,1]
# plaintext=[0,0,0,1,1,1,1,1,1]

# ciphertext=encoding(key,plaintext,sbox)
# print("Открытый текст: "+str(plaintext))
# print("Шифртекст: "+str(ciphertext))
	
str_to_list = lambda x: [int(i) for i in x]

def analysis(sbox):
    random.seed()
    x_list=[] # открытый текст
    y_list=[]
    for i in range(10000):
        x_int=random.randrange(512) #генерируем int, который влезает в 9 бит
        x = f"{x_int:09b}"
        x = str_to_list(x)
        x_list.append(x)
        y_list.append(encoding([0,0,1,0,1,0,0,1, 1], x, sbox))
        # print(x_list)
        # print(y_list)
    sum = calc(x_list, y_list)
    print(sum)
    return sum






def calc(x, y):
    sum = 0
    for i in  range(len(x)):
        xy = y[i][0] ^ y[i][1] ^  x[i][0] ^ x[i][1]
        if xy == 0:
            sum = sum + 1
    return sum


analysis(sbox)