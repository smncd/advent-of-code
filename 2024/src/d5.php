<?php declare(strict_types=1);

function formatData(array $file, string $pattern, string $separator): array
{
    return array_values(array: array_map(
        callback: static fn (string $line): array => array_map(
            callback: static fn (string $number): int => intval(value: $number), 
            array: explode(
                separator: $separator,
                string: $line,
            )
        ),
        array: preg_grep(
            pattern: $pattern,
            array: $file,
        ),
    ));
}


$file = file(filename: '/data/d5.txt');

$pageOrderingRules = formatData(
    file: $file,
    pattern: '/([0-9]{2})\|([0-9]{2})/',
    separator: '|'
);


$pagesToProducte = formatData(
    file: $file,
    pattern: '/(^[0-9,]*$)/',
    separator: ','
);

$isInRightOrder = [];

foreach ($pagesToProducte as $produceIndex => $pages) {
    $correctOrder = true;

    foreach ($pages as $pageIndex => $page) {
        foreach ($pageOrderingRules as $rule) {
            if (
                in_array(needle: $rule[0], haystack: $pages) &&
                in_array(needle: $rule[1], haystack: $pages) 
            ) {
                $ruleZeroIndex =  array_search(needle: $rule[0], haystack: $pages);
                $ruleOneIndex =  array_search(needle: $rule[1], haystack: $pages);

                if (
                    !is_int($ruleZeroIndex) ||
                    !is_int($ruleOneIndex) ||
                    $ruleOneIndex <= $ruleZeroIndex
                ) {
                    $correctOrder = false;
                    break;
                }
            }

            $correctOrder = true;
        }

    }

    if ($correctOrder) $isInRightOrder[] = $pages;
}

$result = 0;

foreach ($isInRightOrder as $pages) {
    $result += $pages[(count($pages) - 1) / 2];
}


echo "$result";