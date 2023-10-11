-module(parte1).
-export([
        start/0,
        find_seat/1,
        seat_id/1
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    string:split(File, "\n", all).

find_seat(Plane, LowerRow, UpperRow, LowerColumn, UpperColumn) ->
    RowCut = erlang:round((UpperRow - LowerRow)/2),
    ColumnCut = erlang:round((UpperColumn - LowerColumn)/2),
    case Plane of
        [] -> {UpperRow, UpperColumn};
        [$F|T] -> find_seat(T, LowerRow, UpperRow - RowCut, LowerColumn, UpperColumn);
        [$B|T] -> find_seat(T, LowerRow + RowCut, UpperRow, LowerColumn, UpperColumn);
        [$R|T] -> find_seat(T, LowerRow, UpperRow, LowerColumn + ColumnCut, UpperColumn);
        [$L|T] -> find_seat(T, LowerRow, UpperRow, LowerColumn, UpperColumn - ColumnCut);
        _ -> undefined
    end.

find_seat(Plane) ->
    find_seat(Plane, 0, 127, 0, 7).

seat_id({Row, Column}) -> Row * 8 + Column.

start() ->
    Input = parse_input(), 
    Seats = [find_seat(Seat) || Seat <- Input],
    Ids = [seat_id(Id) || Id <- Seats],
    erlang:display(lists:max(Ids)).


