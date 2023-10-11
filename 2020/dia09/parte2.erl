-module(parte2).
-export([
        start/0
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    Lines = string:split(File, "\n", all),
    [begin {Y, _} = string:to_integer(X), Y end || X <- Lines].


find_valid(List, Preamble) when is_number(Preamble) ->
    Tail = lists:sublist(List, Preamble),
    Head = lists:sublist(List, Preamble+1, length(List)-Preamble+1),
    find_valid(Tail, Head);

find_valid(_, []) -> undefined;
find_valid([_|Tail] = Preamble, [Current|Head]) ->
    case [X + Y || X <- Preamble, Y <- Preamble, X + Y == Current] of
        [] -> Current;
        _ ->
            find_valid(Tail ++ [Current], Head)
    end.

find_sublist(List, Start, Len, Target) ->
    Sublist = lists:sublist(List, Start, Len),
    case lists:sum(Sublist) of
        X when X > Target -> 
            find_sublist(List, Start+1, 1, Target);
        X when X == Target -> 
            lists:min(Sublist) + lists:max(Sublist);
        X when X < Target ->
            find_sublist(List, Start, Len+1, Target);
        _ -> 
            undefined
    end.

find_sublist(List, Target) -> find_sublist(List, 1, 1, Target).

start() ->
    Input = parse_input(),
    Target = find_valid(Input, 25),
    erlang:display(find_sublist(Input, Target)).



