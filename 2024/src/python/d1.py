
lists = {
    'left': [],
    'right': [],
}

with open('/data/d1.txt') as data:
    for line in data:
        ids = line.split("   ")

        lists['left'].append(int(ids[0]))
        lists['right'].append(int(ids[1]))

for i, l in lists.items():
    lists[i] = sorted(l)

total_distance = 0
similarity_score = 0

for index, value in enumerate(lists['left']):
    distance = abs(value - lists['right'][index])

    times_in_right_list = len([id for id in lists['right'] if id == value])

    total_distance += distance
    similarity_score += value * times_in_right_list

print(f"Total distance: {total_distance}")
print(f"Similarity score: {similarity_score}")