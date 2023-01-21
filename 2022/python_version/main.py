import sys

def main():
    heap = []
    with open(sys.argv[1]) as file:
        sum = 0
        for line in file.readlines():
            line = line.strip()
            if line:
                sum += int(line)
            else:
                heap.append(sum)
                sum = 0
    result = 0
    i = 0
    heap.sort()
    while heap and i < 3:
        value = heap.pop()
        result += value
        i += 1
    print(result)


if __name__ == "__main__":
    main()
