def check(a):
   for i in range(2, a):
      if a % i == 0:
         return False
   return True

def find(n):
    
   # check first I guess
   if check(n):
      return n

   low = n - 1
   high = n + 1

   #finding the nearest prime number
   while True:
      if check(low):
         return low
      elif check(high):
         return high
      else:
         low -= 1
         high += 1

def main():
    n = int(input("input : "))
    print(find(n))    

if __name__ == '__main__':
    main()

