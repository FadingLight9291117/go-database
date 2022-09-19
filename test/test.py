import random
import subprocess

proc = subprocess.Popen("go run .",
                        encoding='utf-8',
                        shell=True,
                        stdin=subprocess.PIPE,
                        stdout=subprocess.PIPE,
                        stderr=subprocess.PIPE,
                        )


def get_number():
    n = 0
    while True:
        n += 1
        yield random.randint(n, n + 100)


ns = []
size = 20
for i, n in enumerate(get_number()):
    ns.append(f"insert {n} clz_{n} email_{n}")
    if i == size:
        break

ns.append(".exit")

inputs = "\n".join(ns)

out, error = proc.communicate(input=inputs)

print(out)
print(error)
