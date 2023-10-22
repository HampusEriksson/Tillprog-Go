vowels = ["a", "e", "i", "o", "u", "y", "å", "ä", "ö"]
consonants = [
    "b",
    "c",
    "d",
    "f",
    "g",
    "h",
    "j",
    "k",
    "l",
    "m",
    "n",
    "p",
    "q",
    "r",
    "s",
    "t",
    "v",
    "w",
    "x",
    "z",
]

name = input("Enter a name: ")

num = len(consonants)
num2 = len(vowels)

cons = False
vow = False

mysen = ""

for char in name:
    for i in range(num):
        if char == consonants[i]:
            mysen += f"{char}o{char}"
            cons = True
        else:
            cons = False

    if cons == False:
        for i in range(num2):
            if char == vowels[i]:
                mysen += char

print(mysen)
