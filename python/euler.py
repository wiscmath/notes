import math
import time


def euler1():
    return math.exp(1)

def euler2(n: int):
    return (1 + 1 / n) ** n

def euler3(precision: int):
    e = 0
    n = 0
    while True:
        term = 1 / math.factorial(n)
        e += term
        if n > precision:
            break
        n += 1
    return round(e, precision)

if __name__ == '__main__':
    t1 = time.perf_counter_ns()
    r1 = euler1()
    t2 = time.perf_counter_ns()

    t3 = time.perf_counter_ns()
    r2 = euler2(1000000)
    t4 = time.perf_counter_ns()

    t5 = time.perf_counter_ns()
    r3 = euler3(1000)
    t6 = time.perf_counter_ns()

    print(f'euler1: {r1}, {t2 - t1}')
    print(f'euler2: {r2}, {t4 - t3}')
    print(f'euler3: {r3}, {t6 - t5}')
