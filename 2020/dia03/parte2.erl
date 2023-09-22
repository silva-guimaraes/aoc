-module(parte2).
-export([
        start/0
        ]).


access(Forest, X, Y) -> 
    Width = length(lists:nth(1, Forest)),
    lists:nth(X rem Width + 1, lists:nth(Y+1, Forest)).


path(Forest, _, Y, _, _) when Y+1 > length(Forest) -> [];
path(Forest, X, Y, Xdegree, Ydegree) ->
    [access(Forest, X, Y) | 
     path(Forest, X+Xdegree, Y+Ydegree, Xdegree, Ydegree)].

start() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = binary:bin_to_list(Binary),
    Forest = string:tokens(File, "\n"),
    Ret =
    length([X || X <- path(Forest, 0, 0, 1, 1), X == $#]) *
    length([X || X <- path(Forest, 0, 0, 3, 1), X == $#]) *
    length([X || X <- path(Forest, 0, 0, 5, 1), X == $#]) *
    length([X || X <- path(Forest, 0, 0, 7, 1), X == $#]) *
    length([X || X <- path(Forest, 0, 0, 1, 2), X == $#]),
    io:format("~p~n", [Ret]).

