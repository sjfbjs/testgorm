import time

sum = 0

start = time.time()
for i in range(10000001):
    sum += i
end = time.time()

print(sum)
print(end - start)