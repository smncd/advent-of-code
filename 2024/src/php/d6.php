<?php declare(strict_types=1);

class Guard
{
    public const DIRECTIONS = [
        'up' => '^',
        'right' => '>',
        'down' => 'v',
        'left' => '<',
    ];

    public ?int $y = null;
    public ?int $x = null;

    public bool $onMap = true;

    public ?string $direction = null;

    public function __construct(array $map) 
    {
        foreach ($map as $y => $row) {
            foreach ($row as $x => $content) {
                if (
                    in_array(needle: $content, haystack: array_values(array: $this::DIRECTIONS)) &&
                    !$this->y &&
                    !$this->x 
                ) {
                    $this->y = $y;
                    $this->x = $x;
                    $this->direction = array_search(
                        needle: $content,
                        haystack: $this::DIRECTIONS, 
                    );
                }
            }
        }
    }

    public function turn(): void 
    {
        $this->direction = match ($this->direction) {
            'up' => 'right',
            'right' => 'down',
            'down' => 'left',
            'left' => 'up',
        };
    }
}

$map = [];

$file = file(filename: '/data/d6.txt');

foreach ($file as $index => $line) {
    $map[$index] = str_split(string: $line);
}

$guard = new Guard(map: $map);

$positions = [
    ['y' => $guard->y, 'x' => $guard->x]
];

while ($guard->onMap) {
    $nextPosition = (object) match ($guard->direction) {
        'up' => ['y' => $guard->y - 1, 'x' => $guard->x],
        'right' => ['y' => $guard->y, 'x' => $guard->x + 1],
        'down' => ['y' => $guard->y + 1, 'x' => $guard->x],
        'left' => ['y' => $guard->y, 'x' => $guard->x - 1],
    };

    $inBounds = (
        ($guard->direction === 'up' && $nextPosition->y >= 0) ||
        ($guard->direction === 'down' && $nextPosition->y <= count($map) - 1) ||
        ($guard->direction === 'left' && $nextPosition->x >= 0) ||
        ($guard->direction === 'right' && $nextPosition->x <= count($map[0]) - 1)
    );

    $nextPositionContent = fn (): mixed => $map[$nextPosition->y][$nextPosition->x];

    if (!$inBounds || !is_string(value: $nextPositionContent())) {
        $guard->onMap = false;
        break;
    }

    if ($nextPositionContent() === '#') {
        $guard->turn();
        continue;
    }

    if (!in_array(needle: $nextPosition, haystack: $positions)) {
        $positions[] = $nextPosition;
    }

    $guard->y = $nextPosition->y;
    $guard->x = $nextPosition->x;
}

$result = count(value: $positions);

echo "Guard will visit {$result} positions";