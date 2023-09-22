-module(parte1).
-export([
        start/0
        ]).


access(Forest, X, Y) -> 
    Width = length(lists:nth(1, Forest)),
    lists:nth(X rem Width + 1, lists:nth(Y+1, Forest)).


path(Forest, _, Y) when Y+1 > length(Forest) -> [];
path(Forest, X, Y) ->
    [access(Forest, X, Y) | path(Forest, X+3, Y+1)].

start() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = binary:bin_to_list(Binary),
    Forest = string:tokens(File, "\n"),
    Path = path(Forest, 0, 0),
    io:format("~p~n", [length([X || X <- Path, X == $#])]).

