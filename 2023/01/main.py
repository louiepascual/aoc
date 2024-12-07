# https://adventofcode.com/2023/day/1/
# O(n^2)

str_digits = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9
}

def main():
    with open("input.txt") as f:
        lines = f.readlines()
        
    total_value = 0
    for i in lines:
        line = i.strip()
        
        nums = []
        itr = 0
        for j in line:
            if j.isdigit():
                nums.append(int(j))
            else:
                for k in str_digits:
                    if line[itr:itr+(len(k))] == k:
                        nums.append(str_digits[k])
                        break
            itr = itr + 1

        # print(nums)
        number = (nums[0] * 10) + nums[-1]
        total_value += number

    print(total_value)

main()