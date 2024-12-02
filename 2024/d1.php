<?php declare(strict_types=1);

$lists = [
    'left' => [],
    'right' => [],
];

$file = file(filename: __DIR__ . '/d1.txt');

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

foreach ($lists['left'] as $key => $value) {
    $sum = $value - $lists['right'][$key];

    if ($sum < 0) $sum = -$sum;

    $totalDistance += $sum;
}

echo "Total distance: $totalDistance\n";

/**
 * Part two
 */

$similarityScore = 0;

foreach ($lists['left'] as $value) {
    $times = count(value: array_filter(
        array: $lists['right'],
        callback: static fn (int $id): bool => $id == $value
    ));

    $similarityScore += $value * $times;
}

echo "Similarity score: $similarityScore\n";

