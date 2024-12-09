lists = %{
  "left" => [],
  "right" => []
}

file_lines = File.stream!("/data/d1.txt")
  |> Enum.to_list()

lists =
  for line <- file_lines, reduce: lists do
    acc ->
      ids = String.split(line, "   ")
      acc
      |> Map.update!("left", &([hd(ids) | &1]))
      |> Map.update!("right", &([Enum.at(ids, 1) | &1]))
  end

lists = Map.update!(lists, "left", &Enum.reverse/1)
lists = Map.update!(lists, "right", &Enum.reverse/1)


IO.inspect(lists)
