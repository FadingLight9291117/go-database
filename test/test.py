import subprocess

proc = subprocess.Popen("go run .",
                        encoding='utf-8',
                        shell=True,
                        stdin=subprocess.PIPE,
                        stdout=subprocess.PIPE,
                        stderr=subprocess.PIPE,
                        )

inputs = "\n".join(
    [
        "insert 0 clz email",
        "insert 0 clz email",
        "insert 0 clz email",
        "select",
        ".exit"
    ]
)

out, error = proc.communicate(input=inputs)

print(out)
# print(error)
