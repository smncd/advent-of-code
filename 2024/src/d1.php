<?php declare(strict_types=1);

$lists = [
    'left' => [],
    'right' => [],
];

$file = file(filename: '/data/d1.txt');

foreach ($file as $key => $line) {
    $ids = explode(separator: '   ', string: $line);

    $lists['left'][] = $ids[0];
    $lists['right'][] = $ids[1];
}

foreach ($lists as $key => $value) {
    sort(array: $lists[$key]);
}

assert(assertion: count(value: $lists['left']) === count(value: $lists['right']));

$totalDistance = 0;
$similarityScore = 0;

foreach ($lists['left'] as $key => $value) {
    $distance = abs(num: $value - $lists['right'][$key]);

    $timesInRightList = count(value: array_filter(
        array: $lists['right'],
        callback: static fn (int $id): bool => $id == $value
    ));

    $totalDistance += $distance;
    $similarityScore += $value * $timesInRightList;
}

echo "Total distance: $totalDistance\n";
echo "Similarity score: $similarityScore\n";
