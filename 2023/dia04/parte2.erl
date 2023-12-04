


-module(parte2).
-export([start/0]).

parse_input(_, eof, List) -> List;

parse_input(File, {ok, Line}, List) ->
    [_, Game] = string:split(string:trim(Line), ": "),
    [_Win, _Yours] = string:split(Game, " | "),
    Win = string:tokens(_Win, " "),
    Yours = string:tokens(_Yours, " "),

    parse_input(File, file:read_line(File), [{Win, Yours}|List]).

parse_input() ->
    {ok, File} = file:open("input.txt", [read]),
    parse_input(File, file:read_line(File), []).

sumlist(B, [])  -> B;
sumlist([], A)  -> A;
sumlist([A| AT], [B | BT]) ->
    [A + B | sumlist(AT, BT)].

loop([], _, Total) -> Total;
loop([{Win, Yours} | T], [N | Copies], Total) ->
    case length(Win) - length(Win -- Yours) of
        0 -> loop(T, Copies, Total + N);
        X -> loop(T, sumlist(lists:duplicate(X, N), Copies), Total + N)
    end.

loop(Cards) ->
    loop(Cards, lists:duplicate(length(Cards), 1), 0).

start() ->
    Input = parse_input(),
    loop(lists:reverse(Input)).




