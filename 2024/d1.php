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

$total = 0;

foreach ($lists['left'] as $key => $value) {
    $sum = $value - $lists['right'][$key];

    if ($sum < 0) $sum = -$sum;

    $total += $sum;
}

echo $total;