<?php declare(strict_types=1);

$file = file_get_contents(
    filename: __DIR__ . '/../data/d3.txt',
);

$matches = [];


preg_match_all(
    pattern: '/mul\((?<first>[0-9]+)\,(?<second>[0-9]+)\)/',
    subject: $file,
    matches: $matches,
);

$firsts = $matches['first'];
$seconds = $matches['second'];

assert(count($firsts) === count($seconds));

$total = 0;

foreach ($firsts as $index => $first) {
    $second = $seconds[$index];

    $total += ($first * $second);
}

echo "Total: $total";