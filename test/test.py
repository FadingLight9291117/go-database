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
        yield i


inputs = "\n".join(
    [f"insert {i} clz email" for i in get_i()] + [".exit"]
)

out, error = proc.communicate(input=inputs)

print(out)
# print(error)
