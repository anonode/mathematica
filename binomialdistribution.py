from math import comb

success = float(input("What is the success rate? "))
failure = 1-success   # F = 1 - S  according to the definitions provided in the Bernoulli Trials


def exactly() -> float:
    need = int(input("Please enter how many success values you need (enter an integer): "))
    outcomes = int(input("Number of outcomes: "))
    prob = comb(outcomes, need) * pow(success, need) * pow(failure, outcomes - need) # nCk * (p(s))^k * (p(f))^(n-k)
    return prob
    

def upto() -> float:
    need = int(input("Please enter how many success values you need (enter an integer): "))
    outcomes = int(input("Number of outcomes: "))
    prob = 0
    for i in range(0,need + 1):
        prob += comb(outcomes,i) * pow(success, i) * pow(failure, outcomes-i)
    return prob


options = "This program has two options:\n\"exactly\": where you might need \
to calculate the probably of getting exactly n positive outcomes\n"\
      "\"up to\": probability of up to n positive outcomes"

print(options)

choice = input("Enter an option: ")
while(choice.casefold() != "exactly" and choice.casefold() != "up to"):
    print("please enter valid input")
    print(options)
    choice = input("Enter an option: ")
    
    
try:
    if(choice.casefold() == "exactly"):
        print(f"\nresult of exactly(): {exactly()}\n")
    elif(choice.casefold() == "up to"):
        print(f"\nresult of upto(): {upto()}\n")
except:
    print("an error occurred somehow, employ superadvanced debugging tactics")
    
    

