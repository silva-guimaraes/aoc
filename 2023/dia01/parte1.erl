


-module(parte1).
-export([start/0]).

parse_input(_, eof, Lines) -> Lines;
parse_input(File, {ok, Line}, Lines) ->
    parse_input(File, file:read_line(File), [string:trim(Line) | Lines]).

parse_input() ->
    {ok, File} = file:open("input.txt", [read]),
    parse_input(File, file:read_line(File), []).


remove_char([])                 -> [];
remove_char([H|T]) when H < $a  -> [H | remove_char(T)];
remove_char([_|T])              -> remove_char(T).

first_last(Numbers) ->
    String = [lists:nth(1, Numbers), lists:last((Numbers))],
    {Number, _} = string:to_integer(String),
    Number.


start() ->
    Input = parse_input(),
    Numbers = lists:map(fun remove_char/1, Input),
    erlang:display(lists:sum(lists:map(fun first_last/1, Numbers))).
