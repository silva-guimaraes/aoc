-module(parte1).
-export([
        start/0
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    Lines = string:split(File, "\n", all),
    [begin {Y, _} = string:to_integer(X), Y end || X <- Lines].

get_diff([_]) -> [3];
get_diff([A, B|T]) ->
    [B - A| get_diff([B|T])].


count_diff([], X) -> X;
count_diff([H|T], X) ->
    #{H := N} = X,
    count_diff(T, X#{H := N+1}).

count_diff(List) -> 
    count_diff(List, #{1 => 0, 2 => 0, 3 => 0}).

start() ->
    Input = parse_input(),
    Adapters = [0] ++ lists:sort(Input),
    Diff = get_diff(Adapters),
    #{1 := OneJolt, 3 := ThreeJolt} = count_diff(Diff),
    erlang:display(Diff),
    erlang:display(OneJolt * ThreeJolt).
