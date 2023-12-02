


-module(parte1).
-export([start/0]).

parse_input(_, eof) -> [];
parse_input(File, {ok, Line}) ->
    [
     string:trim(Line) | parse_input(File, file:read_line(File))
    ].

parse_input() ->
    {ok, File} = file:open("bigboy.txt", [read]),
    parse_input(File, file:read_line(File)).


remove_char([])                 -> [];
remove_char([H|T]) when H < $a  -> [H | remove_char(T)];
remove_char([_|T])              -> remove_char(T).

first_last(Numbers) ->
    String = [lists:nth(1, Numbers), lists:last((Numbers))],
    {Number, _} = string:to_integer(String),
    Number.


start() ->
    Start = erlang:timestamp(),
    Input = parse_input(),
    io:format("time: ~f~n", [timer:now_diff(erlang:timestamp(), Start) / 1000000]),
    Numbers = lists:map(fun remove_char/1, Input),
    erlang:display(lists:sum(lists:map(fun first_last/1, Numbers))).
