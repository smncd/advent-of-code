<?php declare(strict_types=1);
/**
 * Totally ended up GPT-ing this one, but hey, I think I learned something!
 * 
 * My first attempt totally tried to overcomplicate it by creating a map,
 * going in each direction, and from the respective direction going to all 
 * of it's directions, and so on. Clearly overcomplicating it.
 * 
 */

function countWord(array $map, string $word): int
{
    $count = 0;

    $rows = count(value: $map);
    $cols = count(value: $map[0]);

    $directions = [
        [0, 1],   // Horizontal right
        [0, -1],  // Horizontal left
        [1, 0],   // Vertical down
        [-1, 0],  // Vertical up
        [1, 1],   // Diagonal down-right
        [-1, -1], // Diagonal up-left
        [1, -1],  // Diagonal down-left
        [-1, 1],  // Diagonal up-right
    ];

    foreach ($map as $y => $row) {
        foreach ($row as $x => $_) {
            foreach ($directions as [$dy, $dx]) {
                $result = '';

                foreach (str_split($word) as $index => $letter) {
                    $ny = $y + $dy * $index;
                    $nx = $x + $dx * $index;

                    if ($nx < 0 || $ny < 0 || $nx >= $cols || $ny >= $rows) {
                        continue;
                    }
        
                    $result .= $map[$ny][$nx];
                }

                if($result === $word) {
                    ++$count;
                }

            }
        }
    }

    return $count;
}

$map = [];

$file = file(filename: __DIR__ . '/../data/d4.txt');

foreach ($file as $index => $line) {
    $map[$index] = str_split(string: $line);
}

$count = countWord(
    map: $map,
    word: 'XMAS',
);

echo "Total occurances of 'XMAS': $count";
