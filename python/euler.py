import math
import time
import timeit


def euler1():
    return math.exp(1)

def euler2(n:int):
    return (1 + 1 / n) ** n

def euler3(precision:int):
    e = 0
    n = 0
    while True:
        term = 1 / math.factorial(n)
        e += term
        if abs(term) < precision:
            break
        n += 1
    return round(e, precision)

def euler4(precision:int):
    e = 1
    fact = 1
    for i in range(1, precision + 1):
        fact *= 1
        e += 1 / fact
    return e

if __name__ == '__main__':
    t1 = time.perf_counter()
    euler1()
    t2 = time.perf_counter()

    t3 = time.perf_counter()
    euler2(1000000)
    t4 = time.perf_counter()
    print(f'euler1: {t2 - t1}')
    print(f'euler2: {t4 - t3}')

    t5 = timeit.timeit(stmt=euler1)
    print(f'euler1: {t5}')
