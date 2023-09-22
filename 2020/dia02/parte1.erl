-module(parte1).
-export([
        start/0
        ]).

parse(Line) ->
    [Range, [Letter|_], Password] = string:tokens(Line, " "),
    [MinString, MaxString] = string:tokens(Range, "-"),
    {Min, _} = string:to_integer(MinString),
    {Max, _} = string:to_integer(MaxString),
    {Min, Max, Letter, Password}.

valid({Min, Max, Letter, Password}) ->
    Len = string:len([X || X <- Password, X == Letter]),
    (Len =< Max) and (Len >= Min).

start() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = binary:bin_to_list(Binary),
    Lines = string:tokens(File, "\n"),
    Passwords = [parse(X) || X <- Lines],
    Valid = string:len([X || X <- Passwords, valid(X)]),
    io:format("~p~n", [Valid]).

