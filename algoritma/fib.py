import sys

table = [0]*1000

def fib(n):
    if n<=1:
        return n
    else:
        return(fib(n-1) + fib(n-2))

def main():
    num = int(input("Enter a number : "))
    for i in range(num):
        print(fib(i), end=' - ', flush=True)

if __name__ == '__main__':
    main()