import random
import subprocess

proc = subprocess.Popen("go run .",
                        encoding='utf-8',
                        shell=True,
                        stdin=subprocess.PIPE,
                        stdout=subprocess.PIPE,
                        stderr=subprocess.PIPE,
                        )


def get_i():
    for i in range(5):
        yield random.randint(i, i + 100)


inputs = "\n".join(
    [f"insert {i} clz_{i} email_{i}" for i in get_i()] + [".exit"]
)

out, error = proc.communicate(input=inputs)

print(out)
# print(error)
