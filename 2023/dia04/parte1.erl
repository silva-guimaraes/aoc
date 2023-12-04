-module(parte1).
-export([start/0]).

parse_input(_, eof, List) -> List;

parse_input(File, {ok, Line}, List) ->
    % regex omegalul
    [_, Game] = string:split(string:trim(Line), ": "),
    [_Win, _Yours] = string:split(Game, " | "),
    Win = string:tokens(_Win, " "),
    Yours = string:tokens(_Yours, " "),

    parse_input(File, file:read_line(File), [{Win, Yours}|List]).

parse_input() ->
    {ok, File} = file:open("input.txt", [read]),
    parse_input(File, file:read_line(File), []).

start() ->
    Input = parse_input(),
    Cards = 
    [ 
     case length(Win) - length(Win -- Yours) of
         0 -> 0;
         X -> math:pow(2, X-1)
     end
     ||
     {Win, Yours}<- Input 
    ],
    erlang:display(erlang:round(lists:sum(Cards))).

