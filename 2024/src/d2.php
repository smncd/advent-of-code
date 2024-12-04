<?php declare(strict_types=1);

function isSafeRange(int|float $number): bool 
{
    return $number <= 3 && $number >= 1;
}

function processLevels(
    array $levels,
): bool
{
    $isSafe = false;

    $isIncreasing = true;
    $isDecreasing = true;

    foreach ($levels as $index => $level) {
        if (!isset($levels[$index + 1])) break;
    
        $next = $levels[$index + 1];
        $difference = $next - $level;
    
        if ($difference > 0) {
            $isIncreasing = $isIncreasing && isSafeRange(number: $difference);
            $isDecreasing = false;
        } elseif ($difference < 0) {
            $isDecreasing = $isDecreasing && isSafeRange(number: abs(num: $difference));
            $isIncreasing = false;
        } else {
            $isIncreasing = false;
            $isDecreasing = false;
        }   
    
        $isSafe = $isIncreasing || $isDecreasing;
    }

    return $isSafe;
}

$file = file(filename: __DIR__ . '/../data/d2.txt');

$safeReports = 0;
$safeReportsWithDampener = 0;

foreach ($file as $index => $line) {
    $levels = array_map(
        callback: 'intval', 
        array: explode(
            separator: ' ',
            string: $line,
        )
    );

    $isSafe = processLevels(
        levels: $levels,
    );

    if ($isSafe) {
        $safeReports += 1;
    } else {
        foreach ($levels as $index => $level) {
            $levelsToCheck = $levels;

            array_splice(
                array: $levelsToCheck, 
                offset: $index, 
                length: 1
            );

            if(processLevels(levels: $levelsToCheck)) {
                $safeReportsWithDampener += 1;
                break;
            }
        } 
    }

}

$total = $safeReportsWithDampener + $safeReports;

echo "Number of safe reports: $safeReports\n";
echo "Number of safe reports (with Problem Dampener): $safeReportsWithDampener\n";
echo "Total: {$total}\n";